package service_test

import (
	"testing"

	"eve-profit2/internal/repository"
	"eve-profit2/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDatabasePath = "../../../data/sqlite-latest.sqlite"

// TestItemServiceGetItemByID tests the item service with real SDE data
func TestItemServiceGetItemByID(t *testing.T) {
	// Given: Real SDE repository and item service
	sdeRepo, err := repository.NewSDERepository(testDatabasePath)
	require.NoError(t, err)
	defer sdeRepo.Close()

	itemService := service.NewItemService(sdeRepo, nil)

	t.Run("should return item details for valid ID", func(t *testing.T) {
		// When: Get Tritanium (TypeID: 34)
		item, err := itemService.GetItemByID(34)

		// Then: Should return item details
		assert.NoError(t, err)
		assert.NotNil(t, item)
		assert.Equal(t, int32(34), item.TypeID)
		assert.Equal(t, "Tritanium", item.TypeName)
		assert.Greater(t, item.Volume, 0.0)
	})

	t.Run("should return error for invalid ID", func(t *testing.T) {
		// When: Get non-existent item
		item, err := itemService.GetItemByID(999999)

		// Then: Should return error
		assert.Error(t, err)
		assert.Nil(t, item)
	})
}

// TestItemServiceSearchItems tests item search functionality
func TestItemServiceSearchItems(t *testing.T) {
	// Given: Item service with SDE
	sdeRepo, err := repository.NewSDERepository(testDatabasePath)
	require.NoError(t, err)
	defer sdeRepo.Close()

	itemService := service.NewItemService(sdeRepo, nil)

	t.Run("should find items matching search query", func(t *testing.T) {
		// When: Search for "Tritanium" (we know this exists)
		items, err := itemService.SearchItems("Tritanium", 5)

		// Then: Should return results
		assert.NoError(t, err)
		assert.NotEmpty(t, items)
		assert.LessOrEqual(t, len(items), 5)

		// Should contain Tritanium
		found := false
		for _, item := range items {
			if item.TypeID == 34 {
				found = true
				assert.Equal(t, "Tritanium", item.TypeName)
				break
			}
		}
		assert.True(t, found, "Should find Tritanium in search results")
	})
}
