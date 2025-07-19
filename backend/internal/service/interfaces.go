package service

import (
	"context"

	"eve-profit2/internal/models"
)

// ESIClient interface defines the contract for ESI API operations
type ESIClient interface {
	GetMarketOrders(ctx context.Context, regionID int32, typeID int32) ([]models.MarketOrder, error)
	GetMarketHistory(ctx context.Context, regionID int32, typeID int32) ([]models.MarketHistory, error)
	GetTypeInfo(ctx context.Context, typeID int32) (*models.TypeInfo, error)
}
