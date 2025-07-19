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

// Test constants following Clean Code principles
const (
	testContentType     = "Content-Type"
	testApplicationJSON = "application/json"
)

// TestESIClientGetMarketOrders tests retrieving market orders from ESI
func TestESIClientGetMarketOrders(t *testing.T) {
	t.Run("should retrieve market orders for valid region and type", func(t *testing.T) {
		// Given: Mock ESI server responding with market orders
		mockResponse := []models.MarketOrder{
			{
				OrderID:      12345,
				TypeID:       34,       // Tritanium
				LocationID:   60003760, // Jita 4-4
				SystemID:     30000142, // Jita
				VolumeTotal:  1000,
				VolumeRemain: 500,
				MinVolume:    1,
				Price:        5.50,
				IsBuyOrder:   false,
				Duration:     90,
				Range:        "station",
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify correct ESI endpoint
			expectedPath := "/v1/markets/10000002/orders/"
			assert.Equal(t, expectedPath, r.URL.Path)
			assert.Equal(t, "type_id=34", r.URL.RawQuery)

			w.Header().Set(testContentType, testApplicationJSON)
			json.NewEncoder(w).Encode(mockResponse)
		}))
		defer server.Close()

		// When: ESI Client fetches market orders
		client := NewESIClient(WithBaseURL(server.URL))
		ctx := context.Background()

		orders, err := client.GetMarketOrders(ctx, 10000002, 34) // The Forge, Tritanium

		// Then: Should return market orders
		assert.NoError(t, err)
		assert.Len(t, orders, 1)
		assert.Equal(t, int64(12345), orders[0].OrderID)
		assert.Equal(t, int32(34), orders[0].TypeID)
		assert.Equal(t, 5.50, orders[0].Price)
		assert.False(t, orders[0].IsBuyOrder)
	})

	t.Run("should handle ESI server errors gracefully", func(t *testing.T) {
		// Given: ESI server returning 500 error
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Internal Server Error"}`))
		}))
		defer server.Close()

		// When: ESI Client tries to fetch data
		client := NewESIClient(WithBaseURL(server.URL))
		ctx := context.Background()

		orders, err := client.GetMarketOrders(ctx, 10000002, 34)

		// Then: Should return error and no data
		assert.Error(t, err)
		assert.Nil(t, orders)
		assert.Contains(t, err.Error(), "ESI server error")
	})
}

// TestESIClientRateLimiting tests ESI rate limiting compliance
func TestESIClientRateLimiting(t *testing.T) {
	t.Run("should respect 150 requests per second limit", func(t *testing.T) {
		// Given: ESI Client with rate limiting
		requestCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestCount++
			w.Header().Set(testContentType, testApplicationJSON)
			w.Write([]byte(`[]`))
		}))
		defer server.Close()

		client := NewESIClient(
			WithBaseURL(server.URL),
			WithRateLimit(150), // 150 requests per second
		)

		// When: Making multiple requests rapidly
		ctx := context.Background()
		startTime := time.Now()

		// Try to make 10 requests - should be throttled
		for i := 0; i < 10; i++ {
			_, _ = client.GetMarketOrders(ctx, 10000002, 34)
		}

		elapsed := time.Since(startTime)

		// Then: Should be rate limited (not instant)
		assert.Greater(t, elapsed, 50*time.Millisecond) // Should take some time due to rate limiting
		assert.Equal(t, 10, requestCount)
	})

	t.Run("should handle ESI rate limit headers", func(t *testing.T) {
		// Given: ESI server with rate limit headers
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-ESI-Error-Limit-Remain", "95")
			w.Header().Set("X-ESI-Error-Limit-Reset", "60")
			w.Header().Set(testContentType, testApplicationJSON)
			w.Write([]byte(`[]`))
		}))
		defer server.Close()

		// When: ESI Client makes request
		client := NewESIClient(WithBaseURL(server.URL))
		ctx := context.Background()

		_, err := client.GetMarketOrders(ctx, 10000002, 34)

		// Then: Should parse rate limit headers
		assert.NoError(t, err)
		// Rate limit info should be accessible (we'll implement this)
	})
}

// TestESIClientErrorHandling tests error handling and retry logic
func TestESIClientErrorHandling(t *testing.T) {
	t.Run("should retry on temporary network errors", func(t *testing.T) {
		// Given: Server that fails first request but succeeds on retry
		requestCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestCount++
			if requestCount == 1 {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			w.Header().Set(testContentType, testApplicationJSON)
			w.Write([]byte(`[]`))
		}))
		defer server.Close()

		// When: ESI Client makes request with retry enabled
		client := NewESIClient(
			WithBaseURL(server.URL),
			WithRetryAttempts(3),
		)
		ctx := context.Background()

		orders, err := client.GetMarketOrders(ctx, 10000002, 34)

		// Then: Should succeed after retry
		assert.NoError(t, err)
		assert.NotNil(t, orders)
		assert.Equal(t, 2, requestCount) // First failed, second succeeded
	})

	t.Run("should respect context timeout", func(t *testing.T) {
		// Given: Slow server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(200 * time.Millisecond) // Slow response
			w.Header().Set(testContentType, testApplicationJSON)
			w.Write([]byte(`[]`))
		}))
		defer server.Close()

		// When: ESI Client with short timeout
		client := NewESIClient(WithBaseURL(server.URL))
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()

		orders, err := client.GetMarketOrders(ctx, 10000002, 34)

		// Then: Should timeout
		assert.Error(t, err)
		assert.Nil(t, orders)
		assert.Contains(t, err.Error(), "context deadline exceeded")
	})
}

// TestESIClientParallelRequests tests concurrent request handling
func TestESIClientParallelRequests(t *testing.T) {
	t.Run("should handle multiple concurrent requests", func(t *testing.T) {
		// Given: ESI server that can handle multiple requests
		requestCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestCount++
			time.Sleep(10 * time.Millisecond) // Simulate processing time
			w.Header().Set(testContentType, testApplicationJSON)
			w.Write([]byte(`[]`))
		}))
		defer server.Close()

		// When: Making multiple concurrent requests
		client := NewESIClient(WithBaseURL(server.URL))
		ctx := context.Background()

		// Test concurrent requests to different regions
		regions := []int32{10000002, 10000043, 10000032} // The Forge, Domain, Sinq Laison

		type result struct {
			orders []models.MarketOrder
			err    error
		}

		results := make(chan result, len(regions))

		for _, regionID := range regions {
			go func(region int32) {
				orders, err := client.GetMarketOrders(ctx, region, 34)
				results <- result{orders: orders, err: err}
			}(regionID)
		}

		// Collect results
		var successCount int
		for i := 0; i < len(regions); i++ {
			res := <-results
			if res.err == nil {
				successCount++
			}
		}

		// Then: All requests should succeed
		assert.Equal(t, len(regions), successCount)
		assert.Equal(t, len(regions), requestCount)
	})
}
