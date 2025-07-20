package service

import (
	"context"
	"errors"

	"eve-profit2/internal/models"
)

// Common service errors
var (
	ErrItemNotFound = errors.New("item not found")
	ErrInvalidID    = errors.New("invalid ID")
)

// ESIClient interface defines the contract for ESI API operations
type ESIClient interface {
	GetMarketOrders(ctx context.Context, regionID int32, typeID int32) ([]models.MarketOrder, error)
	GetMarketHistory(ctx context.Context, regionID int32, typeID int32) ([]models.MarketHistory, error)
	GetTypeInfo(ctx context.Context, typeID int32) (*models.TypeInfo, error)
}
