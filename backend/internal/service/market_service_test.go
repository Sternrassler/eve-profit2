package service

import (
	"context"
	"testing"
	"time"

	"eve-profit2/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockESIClient is a mock implementation of ESI client for testing
type MockESIClient struct {
	mock.Mock
}

func (m *MockESIClient) GetMarketOrders(ctx context.Context, regionID int32, typeID int32) ([]models.MarketOrder, error) {
	args := m.Called(ctx, regionID, typeID)
	return args.Get(0).([]models.MarketOrder), args.Error(1)
}

func (m *MockESIClient) GetMarketHistory(ctx context.Context, regionID int32, typeID int32) ([]models.MarketHistory, error) {
	args := m.Called(ctx, regionID, typeID)
	return args.Get(0).([]models.MarketHistory), args.Error(1)
}

func (m *MockESIClient) GetTypeInfo(ctx context.Context, typeID int32) (*models.TypeInfo, error) {
	args := m.Called(ctx, typeID)
	return args.Get(0).(*models.TypeInfo), args.Error(1)
}

// TestMarketService_GetMarketData tests comprehensive market data retrieval
func TestMarketService_GetMarketData(t *testing.T) {
	t.Run("should aggregate market data for multiple types", func(t *testing.T) {
		// Given: Mock ESI client with test data
		mockClient := new(MockESIClient)
		service := NewMarketService(mockClient)

		// Test data for Tritanium
		tritaniumOrders := []models.MarketOrder{
			{OrderID: 1, TypeID: 34, Price: 5.50, VolumeRemain: 1000, IsBuyOrder: false},
			{OrderID: 2, TypeID: 34, Price: 5.40, VolumeRemain: 500, IsBuyOrder: true},
		}
		tritaniumHistory := []models.MarketHistory{
			{Date: time.Now().AddDate(0, 0, -1), Average: 5.45, Volume: 100000},
		}

		// Test data for Pyerite
		pyeriteOrders := []models.MarketOrder{
			{OrderID: 3, TypeID: 35, Price: 1.20, VolumeRemain: 2000, IsBuyOrder: false},
			{OrderID: 4, TypeID: 35, Price: 1.15, VolumeRemain: 1500, IsBuyOrder: true},
		}
		pyeriteHistory := []models.MarketHistory{
			{Date: time.Now().AddDate(0, 0, -1), Average: 1.18, Volume: 50000},
		}

		// Setup mock expectations
		mockClient.On("GetMarketOrders", mock.Anything, int32(10000002), int32(34)).Return(tritaniumOrders, nil)
		mockClient.On("GetMarketHistory", mock.Anything, int32(10000002), int32(34)).Return(tritaniumHistory, nil)
		mockClient.On("GetMarketOrders", mock.Anything, int32(10000002), int32(35)).Return(pyeriteOrders, nil)
		mockClient.On("GetMarketHistory", mock.Anything, int32(10000002), int32(35)).Return(pyeriteHistory, nil)

		// When: Requesting market data
		req := MarketDataRequest{
			RegionID: 10000002,        // The Forge
			TypeIDs:  []int32{34, 35}, // Tritanium, Pyerite
		}

		ctx := context.Background()
		response, err := service.GetMarketData(ctx, req)

		// Then: Should return aggregated data
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, int32(10000002), response.RegionID)
		assert.Len(t, response.Data, 2)
		assert.Len(t, response.Orders, 2)
		assert.Len(t, response.History, 2)

		// Verify Tritanium data
		tritaniumData := response.Data[34]
		assert.NotNil(t, tritaniumData)
		assert.Equal(t, 5.40, tritaniumData.BuyMax)  // Highest buy order
		assert.Equal(t, 5.50, tritaniumData.SellMin) // Lowest sell order
		assert.Equal(t, int64(500), tritaniumData.BuyVolume)
		assert.Equal(t, int64(1000), tritaniumData.SellVolume)

		// Verify Pyerite data
		pyeriteData := response.Data[35]
		assert.NotNil(t, pyeriteData)
		assert.Equal(t, 1.15, pyeriteData.BuyMax)
		assert.Equal(t, 1.20, pyeriteData.SellMin)

		// Verify all mock expectations were met
		mockClient.AssertExpectations(t)
	})

	t.Run("should handle empty type IDs list", func(t *testing.T) {
		// Given: Market service
		mockClient := new(MockESIClient)
		service := NewMarketService(mockClient)

		// When: Requesting data with empty type IDs
		req := MarketDataRequest{
			RegionID: 10000002,
			TypeIDs:  []int32{},
		}

		ctx := context.Background()
		response, err := service.GetMarketData(ctx, req)

		// Then: Should return error
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Contains(t, err.Error(), "no type IDs provided")
	})

	t.Run("should handle ESI client errors gracefully", func(t *testing.T) {
		// Given: Mock ESI client that returns errors
		mockClient := new(MockESIClient)
		service := NewMarketService(mockClient)

		// Setup mock to return error
		mockClient.On("GetMarketOrders", mock.Anything, int32(10000002), int32(34)).Return([]models.MarketOrder{}, assert.AnError)

		// When: Requesting market data
		req := MarketDataRequest{
			RegionID: 10000002,
			TypeIDs:  []int32{34},
		}

		ctx := context.Background()
		response, err := service.GetMarketData(ctx, req)

		// Then: Should return error
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Contains(t, err.Error(), "failed to get orders")

		mockClient.AssertExpectations(t)
	})
}

// TestMarketService_CalculateItemPrice tests price calculation logic
func TestMarketService_CalculateItemPrice(t *testing.T) {
	t.Run("should calculate correct buy and sell prices", func(t *testing.T) {
		// Given: Market service and test orders
		mockClient := new(MockESIClient)
		service := NewMarketService(mockClient)

		orders := []models.MarketOrder{
			{OrderID: 1, Price: 5.50, VolumeRemain: 1000, IsBuyOrder: false}, // Sell order
			{OrderID: 2, Price: 5.60, VolumeRemain: 500, IsBuyOrder: false},  // Higher sell
			{OrderID: 3, Price: 5.40, VolumeRemain: 800, IsBuyOrder: true},   // Buy order
			{OrderID: 4, Price: 5.35, VolumeRemain: 600, IsBuyOrder: true},   // Lower buy
		}

		history := []models.MarketHistory{} // Not used in this calculation

		// When: Calculating item price
		itemPrice := service.calculateItemPrice(orders, history)

		// Then: Should return correct aggregated prices
		assert.Equal(t, 5.40, itemPrice.BuyMax)            // Highest buy order
		assert.Equal(t, 5.50, itemPrice.SellMin)           // Lowest sell order
		assert.Equal(t, int64(1400), itemPrice.BuyVolume)  // 800 + 600
		assert.Equal(t, int64(1500), itemPrice.SellVolume) // 1000 + 500
	})

	t.Run("should handle empty orders", func(t *testing.T) {
		// Given: Market service and no orders
		mockClient := new(MockESIClient)
		service := NewMarketService(mockClient)

		orders := []models.MarketOrder{}
		history := []models.MarketHistory{}

		// When: Calculating item price
		itemPrice := service.calculateItemPrice(orders, history)

		// Then: Should return zero values
		assert.Equal(t, 0.0, itemPrice.BuyMax)
		assert.Equal(t, 0.0, itemPrice.SellMin)
		assert.Equal(t, int64(0), itemPrice.BuyVolume)
		assert.Equal(t, int64(0), itemPrice.SellVolume)
	})

	t.Run("should handle only buy orders", func(t *testing.T) {
		// Given: Market service and only buy orders
		mockClient := new(MockESIClient)
		service := NewMarketService(mockClient)

		orders := []models.MarketOrder{
			{OrderID: 1, Price: 5.40, VolumeRemain: 800, IsBuyOrder: true},
			{OrderID: 2, Price: 5.35, VolumeRemain: 600, IsBuyOrder: true},
		}

		history := []models.MarketHistory{}

		// When: Calculating item price
		itemPrice := service.calculateItemPrice(orders, history)

		// Then: Should have buy data but no sell data
		assert.Equal(t, 5.40, itemPrice.BuyMax)
		assert.Equal(t, 0.0, itemPrice.SellMin)
		assert.Equal(t, int64(1400), itemPrice.BuyVolume)
		assert.Equal(t, int64(0), itemPrice.SellVolume)
	})
}
