package esi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"eve-profit2/internal/models"
)

// HTTP Header constants to prevent string duplication
const (
	HeaderUserAgent   = "User-Agent"
	HeaderContentType = "Content-Type"
	HeaderAccept      = "Accept"

	// Content types
	ContentTypeJSON = "application/json"

	// User agent
	UserAgentValue = "EVE-Profit2/1.0"

	// Error message templates
	ErrCreateRequest  = "failed to create request: %w"
	ErrRequestFailed  = "request failed: %w"
	ErrDecodeResponse = "failed to decode response: %w"
)

// ESIClient handles communication with EVE ESI API
type ESIClient struct {
	baseURL     string
	httpClient  *http.Client
	rateLimit   int
	retryLimit  int
	rateLimiter chan struct{}
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

	for range ticker.C {
		select {
		case c.rateLimiter <- struct{}{}:
			// Token added
		default:
			// Channel full, skip
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
	if err := c.waitForRateLimit(ctx); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/markets/%d/orders/?type_id=%d", c.baseURL, regionID, typeID)

	var orders []models.MarketOrder
	err := c.executeWithRetry(ctx, url, &orders)
	return orders, err
}

// waitForRateLimit waits for a rate limiter token
func (c *ESIClient) waitForRateLimit(ctx context.Context) error {
	select {
	case <-c.rateLimiter:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// executeWithRetry performs HTTP request with retry logic
func (c *ESIClient) executeWithRetry(ctx context.Context, url string, result interface{}) error {
	var lastErr error

	for attempt := 0; attempt <= c.retryLimit; attempt++ {
		if err := c.performRequest(ctx, url, result); err != nil {
			lastErr = err
			if c.shouldRetry(err) && attempt < c.retryLimit {
				continue
			}
			return lastErr
		}
		return nil
	}
	return lastErr
}

// performRequest performs a single HTTP request
func (c *ESIClient) performRequest(ctx context.Context, url string, result interface{}) error {
	req, err := c.createRequest(ctx, url)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf(ErrRequestFailed, err)
	}
	defer resp.Body.Close()

	if err := c.handleHTTPError(resp); err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

// createRequest creates a properly configured HTTP request
func (c *ESIClient) createRequest(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf(ErrCreateRequest, err)
	}

	req.Header.Set(HeaderUserAgent, UserAgentValue)
	req.Header.Set(HeaderAccept, ContentTypeJSON)
	return req, nil
}

// handleHTTPError checks HTTP status and returns appropriate error
func (c *ESIClient) handleHTTPError(resp *http.Response) error {
	if resp.StatusCode >= 500 {
		return fmt.Errorf("ESI server error: status %d", resp.StatusCode)
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("ESI client error: status %d", resp.StatusCode)
	}
	return nil
}

// shouldRetry determines if an error should trigger a retry
func (c *ESIClient) shouldRetry(err error) bool {
	return strings.Contains(err.Error(), "ESI server error") ||
		strings.Contains(err.Error(), ErrRequestFailed)
}

// GetMarketHistory retrieves market history for a specific region and type
func (c *ESIClient) GetMarketHistory(ctx context.Context, regionID int32, typeID int32) ([]models.MarketHistory, error) {
	if err := c.waitForRateLimit(ctx); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/markets/%d/history/?type_id=%d", c.baseURL, regionID, typeID)

	var history []models.MarketHistory
	err := c.executeWithRetry(ctx, url, &history)
	return history, err
}

// GetTypeInfo retrieves type information from ESI
func (c *ESIClient) GetTypeInfo(ctx context.Context, typeID int32) (*models.TypeInfo, error) {
	if err := c.waitForRateLimit(ctx); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v3/universe/types/%d/", c.baseURL, typeID)

	var typeInfo models.TypeInfo
	err := c.executeWithRetry(ctx, url, &typeInfo)
	if err != nil {
		return nil, err
	}
	return &typeInfo, nil
}
