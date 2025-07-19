package esi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"eve-profit2/internal/models"

	"github.com/stretchr/testify/assert"
)

// TestESIClient_GetMarketHistory tests retrieving market history from ESI
func TestESIClient_GetMarketHistory(t *testing.T) {
	t.Run("should retrieve market history for valid region and type", func(t *testing.T) {
		// Given: Mock ESI server responding with market history
		mockResponse := []models.MarketHistory{
			{
				Date:       time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
				Average:    5.45,
				Highest:    5.80,
				Lowest:     5.10,
				OrderCount: 1234,
				Volume:     50000000,
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify correct ESI endpoint
			expectedPath := "/v1/markets/10000002/history/"
			assert.Equal(t, expectedPath, r.URL.Path)
			assert.Equal(t, "type_id=34", r.URL.RawQuery)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(mockResponse)
		}))
		defer server.Close()

		// When: ESI Client fetches market history
		client := NewESIClient(WithBaseURL(server.URL))
		ctx := context.Background()

		history, err := client.GetMarketHistory(ctx, 10000002, 34) // The Forge, Tritanium

		// Then: Should return market history
		assert.NoError(t, err)
		assert.Len(t, history, 1)
		assert.Equal(t, 5.45, history[0].Average)
		assert.Equal(t, int64(1234), history[0].OrderCount)
		assert.Equal(t, int64(50000000), history[0].Volume)
	})

	t.Run("should handle empty market history gracefully", func(t *testing.T) {
		// Given: ESI server returning empty history
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[]`))
		}))
		defer server.Close()

		// When: ESI Client tries to fetch data
		client := NewESIClient(WithBaseURL(server.URL))
		ctx := context.Background()

		history, err := client.GetMarketHistory(ctx, 10000002, 34)

		// Then: Should return empty slice without error
		assert.NoError(t, err)
		assert.Len(t, history, 0)
	})
}

// TestESIClient_GetTypeInfo tests retrieving type information from ESI
func TestESIClient_GetTypeInfo(t *testing.T) {
	t.Run("should retrieve type information for valid type ID", func(t *testing.T) {
		// Given: Mock ESI server responding with type info
		mockResponse := models.TypeInfo{
			TypeID:      34,
			Name:        "Tritanium",
			Description: "The most common ore type in the known universe.",
			GroupID:     18,
			Volume:      0.01,
			Mass:        1.0,
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify correct ESI endpoint
			expectedPath := "/v3/universe/types/34/"
			assert.Equal(t, expectedPath, r.URL.Path)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(mockResponse)
		}))
		defer server.Close()

		// When: ESI Client fetches type info
		client := NewESIClient(WithBaseURL(server.URL))
		ctx := context.Background()

		typeInfo, err := client.GetTypeInfo(ctx, 34) // Tritanium

		// Then: Should return type information
		assert.NoError(t, err)
		assert.Equal(t, int32(34), typeInfo.TypeID)
		assert.Equal(t, "Tritanium", typeInfo.Name)
		assert.Equal(t, 0.01, typeInfo.Volume)
	})

	t.Run("should handle invalid type ID", func(t *testing.T) {
		// Given: ESI server returning 404 for invalid type
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "Type not found"}`))
		}))
		defer server.Close()

		// When: ESI Client tries to fetch invalid type
		client := NewESIClient(WithBaseURL(server.URL))
		ctx := context.Background()

		typeInfo, err := client.GetTypeInfo(ctx, 999999)

		// Then: Should return error
		assert.Error(t, err)
		assert.Nil(t, typeInfo)
		assert.Contains(t, err.Error(), "ESI client error")
	})
}

// TestESIClient_BatchOperations tests batch processing capabilities
func TestESIClient_BatchOperations(t *testing.T) {
	t.Run("should handle batch market data requests efficiently", func(t *testing.T) {
		// Given: ESI server that tracks request patterns
		requestLog := make([]string, 0)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestLog = append(requestLog, r.URL.Path+"?"+r.URL.RawQuery)
			w.Header().Set("Content-Type", "application/json")

			if r.URL.Path == "/v1/markets/10000002/orders/" {
				w.Write([]byte(`[]`))
			} else if r.URL.Path == "/v1/markets/10000002/history/" {
				w.Write([]byte(`[]`))
			}
		}))
		defer server.Close()

		// When: Requesting data for multiple types
		client := NewESIClient(WithBaseURL(server.URL))
		ctx := context.Background()

		typeIDs := []int32{34, 35, 36} // Tritanium, Pyerite, Mexallon
		regionID := int32(10000002)    // The Forge

		// Concurrent requests for orders and history
		results := make(chan error, len(typeIDs)*2)

		for _, typeID := range typeIDs {
			go func(tid int32) {
				_, err := client.GetMarketOrders(ctx, regionID, tid)
				results <- err
			}(typeID)

			go func(tid int32) {
				_, err := client.GetMarketHistory(ctx, regionID, tid)
				results <- err
			}(typeID)
		}

		// Collect results
		var errorCount int
		for i := 0; i < len(typeIDs)*2; i++ {
			if err := <-results; err != nil {
				errorCount++
			}
		}

		// Then: All requests should succeed
		assert.Equal(t, 0, errorCount)
		assert.Equal(t, len(typeIDs)*2, len(requestLog))

		// Verify correct endpoints were called
		orderRequests := 0
		historyRequests := 0
		for _, req := range requestLog {
			if req == "/v1/markets/10000002/orders/?type_id=34" ||
				req == "/v1/markets/10000002/orders/?type_id=35" ||
				req == "/v1/markets/10000002/orders/?type_id=36" {
				orderRequests++
			} else if req == "/v1/markets/10000002/history/?type_id=34" ||
				req == "/v1/markets/10000002/history/?type_id=35" ||
				req == "/v1/markets/10000002/history/?type_id=36" {
				historyRequests++
			}
		}

		assert.Equal(t, 3, orderRequests)
		assert.Equal(t, 3, historyRequests)
	})
}
