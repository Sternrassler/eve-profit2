package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"eve-profit2/internal/models"
	"eve-profit2/internal/repository"
)

// MarketService handles market data operations with ESI integration
type MarketService struct {
	esiClient ESIClient
	cache     map[string]interface{}
	cacheMux  sync.RWMutex
	cacheTTL  time.Duration
}

func NewMarketService(esiClient ESIClient) *MarketService {
	return &MarketService{
		esiClient: esiClient,
		cache:     make(map[string]interface{}),
		cacheTTL:  5 * time.Minute, // 5-minute cache TTL
	}
}

// getCacheKey generates a cache key for market data
func (s *MarketService) getCacheKey(regionID int32, typeIDs []int32) string {
	return fmt.Sprintf("market_%d_%v", regionID, typeIDs)
}

// getCachedData retrieves data from cache if available and not expired
func (s *MarketService) getCachedData(key string) (*MarketDataResponse, bool) {
	s.cacheMux.RLock()
	defer s.cacheMux.RUnlock()

	if data, exists := s.cache[key]; exists {
		if response, ok := data.(*MarketDataResponse); ok {
			if time.Since(response.UpdatedAt) < s.cacheTTL {
				return response, true
			}
		}
	}
	return nil, false
}

// setCachedData stores data in cache
func (s *MarketService) setCachedData(key string, data *MarketDataResponse) {
	s.cacheMux.Lock()
	defer s.cacheMux.Unlock()
	s.cache[key] = data
}

// MarketDataRequest represents a request for market data
type MarketDataRequest struct {
	RegionID int32   `json:"region_id"`
	TypeIDs  []int32 `json:"type_ids"`
}

// MarketDataResponse represents aggregated market data
type MarketDataResponse struct {
	RegionID  int32                            `json:"region_id"`
	Data      map[int32]*models.ItemPrice      `json:"data"`
	Orders    map[int32][]models.MarketOrder   `json:"orders,omitempty"`
	History   map[int32][]models.MarketHistory `json:"history,omitempty"`
	UpdatedAt time.Time                        `json:"updated_at"`
}

// GetMarketData retrieves comprehensive market data for specified types in a region
func (s *MarketService) GetMarketData(ctx context.Context, req MarketDataRequest) (*MarketDataResponse, error) {
	if err := s.validateMarketDataRequest(req); err != nil {
		return nil, err
	}

	// Check cache first
	cacheKey := s.getCacheKey(req.RegionID, req.TypeIDs)
	if cachedData, found := s.getCachedData(cacheKey); found {
		return cachedData, nil
	}

	results, err := s.fetchMarketDataConcurrently(ctx, req)
	if err != nil {
		return nil, err
	}

	response := s.aggregateMarketDataResponse(req.RegionID, results)
	s.setCachedData(cacheKey, response)

	return response, nil
}

// validateMarketDataRequest validates the market data request
func (s *MarketService) validateMarketDataRequest(req MarketDataRequest) error {
	if len(req.TypeIDs) == 0 {
		return fmt.Errorf("no type IDs provided")
	}
	if req.RegionID <= 0 {
		return fmt.Errorf("invalid region ID: %d", req.RegionID)
	}
	return nil
}

// marketDataResult represents the result of fetching data for a single type
type marketDataResult struct {
	typeID  int32
	orders  []models.MarketOrder
	history []models.MarketHistory
	err     error
}

// fetchMarketDataConcurrently fetches market data for all type IDs concurrently
func (s *MarketService) fetchMarketDataConcurrently(ctx context.Context, req MarketDataRequest) ([]marketDataResult, error) {
	results := make(chan marketDataResult, len(req.TypeIDs))
	var wg sync.WaitGroup

	for _, typeID := range req.TypeIDs {
		wg.Add(1)
		go func(tid int32) {
			defer wg.Done()
			s.fetchSingleTypeMarketData(ctx, req.RegionID, tid, results)
		}(typeID)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return s.collectMarketDataResults(results)
}

// fetchSingleTypeMarketData fetches market data for a single type ID
func (s *MarketService) fetchSingleTypeMarketData(ctx context.Context, regionID, typeID int32, results chan<- marketDataResult) {
	orders, err := s.esiClient.GetMarketOrders(ctx, regionID, typeID)
	if err != nil {
		results <- marketDataResult{
			typeID: typeID,
			err:    fmt.Errorf("failed to get orders for type %d: %w", typeID, err),
		}
		return
	}

	history, err := s.esiClient.GetMarketHistory(ctx, regionID, typeID)
	if err != nil {
		results <- marketDataResult{
			typeID: typeID,
			err:    fmt.Errorf("failed to get history for type %d: %w", typeID, err),
		}
		return
	}

	results <- marketDataResult{
		typeID:  typeID,
		orders:  orders,
		history: history,
	}
}

// collectMarketDataResults collects and validates all market data results
func (s *MarketService) collectMarketDataResults(results <-chan marketDataResult) ([]marketDataResult, error) {
	var collectedResults []marketDataResult

	for res := range results {
		if res.err != nil {
			return nil, res.err
		}
		collectedResults = append(collectedResults, res)
	}

	return collectedResults, nil
}

// aggregateMarketDataResponse aggregates results into the final response
func (s *MarketService) aggregateMarketDataResponse(regionID int32, results []marketDataResult) *MarketDataResponse {
	response := &MarketDataResponse{
		RegionID:  regionID,
		Data:      make(map[int32]*models.ItemPrice),
		Orders:    make(map[int32][]models.MarketOrder),
		History:   make(map[int32][]models.MarketHistory),
		UpdatedAt: time.Now(),
	}

	for _, res := range results {
		response.Orders[res.typeID] = res.orders
		response.History[res.typeID] = res.history

		// Calculate current market prices from orders
		itemPrice := s.calculateItemPrice(res.orders, res.history)
		itemPrice.TypeID = res.typeID
		itemPrice.LastUpdated = response.UpdatedAt
		response.Data[res.typeID] = itemPrice
	}

	return response
}

// calculateItemPrice aggregates market orders into current price information
func (s *MarketService) calculateItemPrice(orders []models.MarketOrder, _ []models.MarketHistory) *models.ItemPrice {
	price := &models.ItemPrice{
		BuyMax:     0,
		SellMin:    0,
		BuyVolume:  0,
		SellVolume: 0,
	}

	if len(orders) == 0 {
		return price
	}

	// Separate buy and sell orders
	var buyOrders, sellOrders []models.MarketOrder
	for _, order := range orders {
		if order.IsBuyOrder {
			buyOrders = append(buyOrders, order)
		} else {
			sellOrders = append(sellOrders, order)
		}
	}

	// Calculate best buy price (highest buy order)
	for _, order := range buyOrders {
		if order.Price > price.BuyMax {
			price.BuyMax = order.Price
		}
		price.BuyVolume += int64(order.VolumeRemain)
	}

	// Calculate best sell price (lowest sell order)
	for i, order := range sellOrders {
		if i == 0 || order.Price < price.SellMin {
			price.SellMin = order.Price
		}
		price.SellVolume += int64(order.VolumeRemain)
	}

	return price
}

// ItemService handles SDE item operations
type ItemService struct {
	sdeRepo      *repository.SDERepository
	cacheManager interface{}
}

func NewItemService(sdeRepo interface{}, cacheManager interface{}) *ItemService {
	repo, ok := sdeRepo.(*repository.SDERepository)
	if !ok {
		return &ItemService{}
	}
	return &ItemService{
		sdeRepo:      repo,
		cacheManager: cacheManager,
	}
}

// GetItemByID retrieves an item by its type ID
func (s *ItemService) GetItemByID(typeID int32) (*models.Item, error) {
	if s.sdeRepo == nil {
		return nil, fmt.Errorf("SDE repository not available")
	}

	sdeItem, err := s.sdeRepo.GetItemByID(typeID)
	if err != nil {
		// Check if item was not found
		if err.Error() == fmt.Sprintf("item not found: typeID %d", typeID) {
			return nil, ErrItemNotFound
		}
		return nil, err
	}

	// Convert SDE item to models.Item
	item := &models.Item{
		TypeID:   sdeItem.TypeID,
		TypeName: sdeItem.TypeName,
		GroupID:  sdeItem.GroupID,
		Volume:   sdeItem.Volume,
		Mass:     0, // Mass not available in SDE
	}

	return item, nil
}

// SearchItems searches for items by name pattern
func (s *ItemService) SearchItems(pattern string) ([]*models.Item, error) {
	if s.sdeRepo == nil {
		return nil, fmt.Errorf("SDE repository not available")
	}

	sdeItems, err := s.sdeRepo.SearchItems(pattern, 50) // Default limit of 50
	if err != nil {
		return nil, err
	}

	// Convert SDE items to models.Item pointers
	var items []*models.Item
	for _, sdeItem := range sdeItems {
		item := &models.Item{
			TypeID:   sdeItem.TypeID,
			TypeName: sdeItem.TypeName,
			GroupID:  sdeItem.GroupID,
			Volume:   sdeItem.Volume,
			Mass:     0, // Mass not available in SDE
		}
		items = append(items, item)
	}

	return items, nil
}
