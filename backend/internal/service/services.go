package service

import (
	"fmt"

	"eve-profit2/internal/models"
	"eve-profit2/internal/repository"
)

// MarketService handles market data operations
type MarketService struct {
	// cacheManager *cache.CacheManager
}

func NewMarketService(cacheManager interface{}) *MarketService {
	return &MarketService{}
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
	}

	return item, nil
}

// SearchItems searches for items by name
func (s *ItemService) SearchItems(query string, limit int) ([]*models.Item, error) {
	if s.sdeRepo == nil {
		return nil, fmt.Errorf("SDE repository not available")
	}

	sdeItems, err := s.sdeRepo.SearchItems(query, limit)
	if err != nil {
		return nil, err
	}

	// Convert SDE items to models.Item
	items := make([]*models.Item, len(sdeItems))
	for i, sdeItem := range sdeItems {
		items[i] = &models.Item{
			TypeID:   sdeItem.TypeID,
			TypeName: sdeItem.TypeName,
			GroupID:  sdeItem.GroupID,
			Volume:   sdeItem.Volume,
		}
	}

	return items, nil
}

// ProfitService handles profit calculations
type ProfitService struct {
	// marketService *MarketService
	// itemService *ItemService
}

func NewProfitService(marketService *MarketService, itemService *ItemService) *ProfitService {
	return &ProfitService{}
}

// CharacterService handles character data operations
type CharacterService struct {
	// cacheManager *cache.CacheManager
}

func NewCharacterService(cacheManager interface{}) *CharacterService {
	return &CharacterService{}
}
