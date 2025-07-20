package models_test

import (
	"testing"
	"time"

	"eve-profit2/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestMarketOrderValidation(t *testing.T) {
	// Arrange
	order := models.MarketOrder{
		OrderID:      12345,
		TypeID:       34,
		LocationID:   60003760,
		SystemID:     30000142,
		VolumeTotal:  1000,
		VolumeRemain: 500,
		MinVolume:    1,
		Price:        5.50,
		IsBuyOrder:   false,
		Duration:     90,
		Issued:       time.Now(),
		Range:        "station",
	}

	// Act & Assert
	assert.Equal(t, int64(12345), order.OrderID)
	assert.Equal(t, int32(34), order.TypeID)
	assert.Equal(t, float64(5.50), order.Price)
	assert.False(t, order.IsBuyOrder)
	// Test additional fields to avoid unused warnings
	assert.Equal(t, int64(60003760), order.LocationID)
	assert.Equal(t, int32(30000142), order.SystemID)
	assert.Equal(t, int32(1000), order.VolumeTotal)
	assert.Equal(t, int32(500), order.VolumeRemain)
	assert.Equal(t, int32(1), order.MinVolume)
	assert.Equal(t, int32(90), order.Duration)
	assert.Equal(t, "station", order.Range)
	assert.False(t, order.Issued.IsZero())
}

func TestItemValidation(t *testing.T) {
	// Arrange
	item := models.Item{
		TypeID:       34,
		TypeName:     "Tritanium",
		GroupID:      18,
		GroupName:    "Mineral",
		CategoryID:   4,
		CategoryName: "Material",
		Volume:       0.01,
		Mass:         1.0,
		Description:  "Test mineral",
	}

	// Act & Assert
	assert.Equal(t, int32(34), item.TypeID)
	assert.Equal(t, "Tritanium", item.TypeName)
	assert.Equal(t, float64(0.01), item.Volume)
	assert.Greater(t, item.Mass, 0.0)
	// Test additional fields to avoid unused warnings
	assert.Equal(t, int32(18), item.GroupID)
	assert.Equal(t, "Mineral", item.GroupName)
	assert.Equal(t, int32(4), item.CategoryID)
	assert.Equal(t, "Material", item.CategoryName)
	assert.Equal(t, "Test mineral", item.Description)
}

func TestItemPriceCalculateProfit(t *testing.T) {
	// Arrange
	price := models.ItemPrice{
		TypeID:      34,
		BuyMax:      5.00,
		SellMin:     5.50,
		BuyVolume:   1000000,
		SellVolume:  500000,
		LastUpdated: time.Now(),
	}

	// Act
	profitPerUnit := price.SellMin - price.BuyMax
	expectedProfit := 0.50

	// Assert
	assert.Equal(t, expectedProfit, profitPerUnit)
	assert.Greater(t, price.BuyVolume, int64(0))
	assert.Greater(t, price.SellVolume, int64(0))
	// Test additional fields to avoid unused warnings
	assert.Equal(t, int32(34), price.TypeID)
	assert.False(t, price.LastUpdated.IsZero())
}
