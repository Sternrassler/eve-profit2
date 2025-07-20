package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"eve-profit2/internal/api/handlers"
	"eve-profit2/internal/repository"
	"eve-profit2/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestItemHandlerIntegration(t *testing.T) {
	// Set up real services for integration test
	gin.SetMode(gin.TestMode)

	// Create real SDE repository
	sdeRepo, err := repository.NewSDERepository("../../data/sqlite-latest.sqlite")
	if err != nil {
		t.Skipf("Skipping integration test: SDE database not available: %v", err)
		return
	}

	// Create real item service
	itemService := service.NewItemService(sdeRepo, nil)

	// Create handler with real service
	handler := handlers.NewItemHandler(itemService)

	t.Run("should get real item details for Tritanium", func(t *testing.T) {
		// Arrange
		router := gin.New()
		router.GET("/api/v1/items/:item_id", handler.GetItemDetails)

		req := httptest.NewRequest("GET", "/api/v1/items/34", nil) // Tritanium
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		// Print response for debugging
		t.Logf("Response: %s", w.Body.String())
	})

	t.Run("should search for real items", func(t *testing.T) {
		// Arrange
		router := gin.New()
		router.GET("/api/v1/items/search", handler.SearchItems)

		req := httptest.NewRequest("GET", "/api/v1/items/search?q=Tritanium", nil)
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		// Print response for debugging
		t.Logf("Response: %s", w.Body.String())
	})
}
