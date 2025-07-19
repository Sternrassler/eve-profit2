package esi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"eve-profit2/internal/models"
)

// ESIClient handles communication with EVE ESI API
type ESIClient struct {
	baseURL     string
	httpClient  *http.Client
	rateLimit   int
	retryLimit  int
	rateLimiter chan struct{}
	rateMutex   sync.Mutex
}

// ClientOption configures the ESI client
type ClientOption func(*ESIClient)

// NewESIClient creates a new ESI client with default configuration
func NewESIClient(options ...ClientOption) *ESIClient {
	client := &ESIClient{
		baseURL:     "https://esi.evetech.net",
		httpClient:  &http.Client{Timeout: 30 * time.Second},
		rateLimit:   150, // ESI default: 150 requests per second
		retryLimit:  3,
		rateLimiter: make(chan struct{}, 150), // Buffer for rate limiting
	}

	for _, option := range options {
		option(client)
	}

	// Initialize rate limiter
	go client.initRateLimiter()

	return client
}

// initRateLimiter manages the rate limiting tokens
func (c *ESIClient) initRateLimiter() {
	ticker := time.NewTicker(time.Second / time.Duration(c.rateLimit))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case c.rateLimiter <- struct{}{}:
				// Token added
			default:
				// Channel full, skip
			}
		}
	}
}

// WithBaseURL sets a custom base URL for the ESI client
func WithBaseURL(url string) ClientOption {
	return func(c *ESIClient) {
		c.baseURL = url
	}
}

// WithRateLimit sets the rate limit for requests per second
func WithRateLimit(limit int) ClientOption {
	return func(c *ESIClient) {
		c.rateLimit = limit
		c.rateLimiter = make(chan struct{}, limit)
	}
}

// WithRetryAttempts sets the number of retry attempts
func WithRetryAttempts(attempts int) ClientOption {
	return func(c *ESIClient) {
		c.retryLimit = attempts
	}
}

// GetMarketOrders retrieves market orders for a specific region and type
func (c *ESIClient) GetMarketOrders(ctx context.Context, regionID int32, typeID int32) ([]models.MarketOrder, error) {
	// Wait for rate limit token
	select {
	case <-c.rateLimiter:
		// Got token, proceed
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	// Build URL with query parameters
	url := fmt.Sprintf("%s/v1/markets/%d/orders/?type_id=%d", c.baseURL, regionID, typeID)

	var lastErr error

	// Retry logic
	for attempt := 0; attempt <= c.retryLimit; attempt++ {
		// Create request
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		// Set headers
		req.Header.Set("User-Agent", "EVE-Profit2/1.0")
		req.Header.Set("Accept", "application/json")

		// Execute request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("request failed: %w", err)
			if attempt < c.retryLimit {
				continue // Retry on network errors
			}
			return nil, lastErr
		}
		defer resp.Body.Close()

		// Handle HTTP errors
		if resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("ESI server error: status %d", resp.StatusCode)
			if attempt < c.retryLimit {
				continue // Retry on server errors
			}
			return nil, lastErr
		}

		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("ESI client error: status %d", resp.StatusCode)
		}

		// Parse response
		var orders []models.MarketOrder
		if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}

		return orders, nil
	}

	return nil, lastErr
}

// GetMarketHistory retrieves market history for a specific region and type
func (c *ESIClient) GetMarketHistory(ctx context.Context, regionID int32, typeID int32) ([]models.MarketHistory, error) {
	// Wait for rate limit token
	select {
	case <-c.rateLimiter:
		// Got token, proceed
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	// Build URL with query parameters
	url := fmt.Sprintf("%s/v1/markets/%d/history/?type_id=%d", c.baseURL, regionID, typeID)

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("User-Agent", "EVE-Profit2/1.0")
	req.Header.Set("Accept", "application/json")

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Handle HTTP errors
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("ESI error: status %d", resp.StatusCode)
	}

	// Parse response
	var history []models.MarketHistory
	if err := json.NewDecoder(resp.Body).Decode(&history); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return history, nil
}

// GetTypeInfo retrieves type information from ESI
func (c *ESIClient) GetTypeInfo(ctx context.Context, typeID int32) (*models.TypeInfo, error) {
	// Wait for rate limit token
	select {
	case <-c.rateLimiter:
		// Got token, proceed
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	// Build URL
	url := fmt.Sprintf("%s/v3/universe/types/%d/", c.baseURL, typeID)

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("User-Agent", "EVE-Profit2/1.0")
	req.Header.Set("Accept", "application/json")

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Handle HTTP errors
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("ESI client error: status %d", resp.StatusCode)
	}

	// Parse response
	var typeInfo models.TypeInfo
	if err := json.NewDecoder(resp.Body).Decode(&typeInfo); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &typeInfo, nil
}
