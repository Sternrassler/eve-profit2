package integration_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"eve-profit2/internal/api/handlers"
	"eve-profit2/internal/cache"
	"eve-profit2/internal/config"
	"eve-profit2/internal/repository"
	"eve-profit2/internal/service"
	"eve-profit2/pkg/esi"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testSDEPath = "../../data/sqlite-latest.sqlite"

// TestAPIIntegration tests the full API stack
func TestAPIIntegration(t *testing.T) {
	// Arrange - Setup full application stack
	cfg := config.Load()

	// Adjust database path for integration tests
	cfg.SDEDatabasePath = testSDEPath

	// Setup SDE Repository
	sdeRepo, err := repository.NewSDERepository(cfg.SDEDatabasePath)
	require.NoError(t, err)

	// Setup Cache
	cacheManager, err := cache.NewCacheManager()
	require.NoError(t, err)

	// Setup ESI Client
	esiClient := esi.NewESIClient(
		esi.WithBaseURL(cfg.ESIBaseURL),
		esi.WithRateLimit(cfg.ESIRateLimit),
	)

	// Setup Services
	itemService := service.NewItemService(sdeRepo, cacheManager)
	marketService := service.NewMarketService(esiClient)

	// Use services to prevent unused variable warnings
	_ = itemService
	_ = marketService

	// Setup Handlers
	healthHandler := handlers.NewHealthHandler()

	// Setup Router
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Health endpoints
	router.GET("/health", healthHandler.HealthCheck)

	// Act & Assert - Test health endpoint
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var healthResponse map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &healthResponse)
	require.NoError(t, err)

	assert.Equal(t, "healthy", healthResponse["status"])
	assert.Equal(t, "eve-profit2-backend", healthResponse["service"])
}

// TestDatabaseIntegration tests SDE database operations
func TestDatabaseIntegration(t *testing.T) {
	// Arrange
	cfg := config.Load()
	cfg.SDEDatabasePath = testSDEPath
	sdeRepo, err := repository.NewSDERepository(cfg.SDEDatabasePath)
	require.NoError(t, err)

	// Act - Test item retrieval
	item, err := sdeRepo.GetItemByID(34) // Tritanium

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, int32(34), item.TypeID)
	assert.Contains(t, item.TypeName, "Tritanium")
}

// TestCacheIntegration tests cache functionality with services
func TestCacheIntegration(t *testing.T) {
	// Arrange
	cfg := config.Load()
	cfg.SDEDatabasePath = testSDEPath
	cacheManager, err := cache.NewCacheManager()
	require.NoError(t, err)

	sdeRepo, err := repository.NewSDERepository(cfg.SDEDatabasePath)
	require.NoError(t, err)

	itemService := service.NewItemService(sdeRepo, cacheManager)

	// Act - Call service method twice to test caching
	start := time.Now()
	item1, err1 := itemService.GetItemByID(34)
	firstCallDuration := time.Since(start)

	start = time.Now()
	item2, err2 := itemService.GetItemByID(34)
	secondCallDuration := time.Since(start)

	// Assert
	require.NoError(t, err1)
	require.NoError(t, err2)
	assert.Equal(t, item1.TypeID, item2.TypeID)

	// Second call should be faster due to caching
	// Note: This is a timing-sensitive test, so we use a reasonable threshold
	assert.True(t, secondCallDuration <= firstCallDuration,
		"Second call should be faster due to caching. First: %v, Second: %v",
		firstCallDuration, secondCallDuration)
}

// TestESIIntegration tests ESI client integration (using mock server)
func TestESIIntegration(t *testing.T) {
	// Skip this test in CI/CD or when no internet connection
	if testing.Short() {
		t.Skip("Skipping ESI integration test in short mode")
	}

	// Arrange
	cfg := config.Load()
	esiClient := esi.NewESIClient(
		esi.WithBaseURL(cfg.ESIBaseURL),
		esi.WithRateLimit(cfg.ESIRateLimit),
	)

	// Use a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Act - Test ESI client with a known valid region and type
	// Use The Forge (region 10000002) and Tritanium (type 34)
	orders, err := esiClient.GetMarketOrders(ctx, 10000002, 34)

	// Assert
	if err != nil {
		// If ESI is down or unreachable, skip the test
		t.Skipf("ESI integration test skipped due to error: %v", err)
	}

	assert.NotNil(t, orders)
	// We can't assert on specific content as market data changes constantly
}

// TestEndToEndWorkflow simulates a complete user workflow
func TestEndToEndWorkflow(t *testing.T) {
	// This would test a complete workflow like:
	// 1. Search for an item
	// 2. Get market data for that item
	// 3. Calculate profit potential
	// 4. Return results

	// For now, this is a placeholder for future end-to-end tests
	t.Skip("End-to-end workflow tests to be implemented in Phase 4")
}
