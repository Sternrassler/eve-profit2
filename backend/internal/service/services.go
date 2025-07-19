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
	if len(req.TypeIDs) == 0 {
		return nil, fmt.Errorf("no type IDs provided")
	}

	response := &MarketDataResponse{
		RegionID:  req.RegionID,
		Data:      make(map[int32]*models.ItemPrice),
		Orders:    make(map[int32][]models.MarketOrder),
		History:   make(map[int32][]models.MarketHistory),
		UpdatedAt: time.Now(),
	}

	// Process each type ID concurrently
	type result struct {
		typeID  int32
		orders  []models.MarketOrder
		history []models.MarketHistory
		err     error
	}

	results := make(chan result, len(req.TypeIDs))
	var wg sync.WaitGroup

	for _, typeID := range req.TypeIDs {
		wg.Add(1)
		go func(tid int32) {
			defer wg.Done()

			// Fetch orders
			orders, err := s.esiClient.GetMarketOrders(ctx, req.RegionID, tid)
			if err != nil {
				results <- result{typeID: tid, err: fmt.Errorf("failed to get orders for type %d: %w", tid, err)}
				return
			}

			// Fetch history
			history, err := s.esiClient.GetMarketHistory(ctx, req.RegionID, tid)
			if err != nil {
				results <- result{typeID: tid, err: fmt.Errorf("failed to get history for type %d: %w", tid, err)}
				return
			}

			results <- result{
				typeID:  tid,
				orders:  orders,
				history: history,
			}
		}(typeID)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results and aggregate data
	for res := range results {
		if res.err != nil {
			return nil, res.err
		}

		response.Orders[res.typeID] = res.orders
		response.History[res.typeID] = res.history

		// Calculate current market prices from orders
		itemPrice := s.calculateItemPrice(res.orders, res.history)
		itemPrice.TypeID = res.typeID
		itemPrice.LastUpdated = response.UpdatedAt
		response.Data[res.typeID] = itemPrice
	}

	return response, nil
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
func (s *ItemService) SearchItems(pattern string, limit int) ([]models.Item, error) {
	if s.sdeRepo == nil {
		return nil, fmt.Errorf("SDE repository not available")
	}

	sdeItems, err := s.sdeRepo.SearchItems(pattern, limit)
	if err != nil {
		return nil, err
	}

	// Convert SDE items to models.Item
	var items []models.Item
	for _, sdeItem := range sdeItems {
		item := models.Item{
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
