package repository_test

import (
	"testing"

	"eve-profit2/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDatabasePath = "../../../data/sqlite-latest.sqlite"

func TestSDERepositoryGetItemByID(t *testing.T) {
	// Arrange
	repo, err := repository.NewSDERepository("../../../data/sqlite-latest.sqlite")
	require.NoError(t, err)
	defer repo.Close()

	// Act
	item, err := repo.GetItemByID(34) // Tritanium

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, int32(34), item.TypeID)
	assert.Equal(t, "Tritanium", item.TypeName)
	assert.Equal(t, int32(18), item.GroupID) // Mineral group
	assert.True(t, item.Published)
}

func TestSDERepositoryGetItemByIDNotFound(t *testing.T) {
	// Arrange
	repo, err := repository.NewSDERepository(testDatabasePath)
	require.NoError(t, err)
	defer repo.Close()

	// Act
	item, err := repo.GetItemByID(999999) // Non-existent item

	// Assert
	assert.Error(t, err)
	assert.Nil(t, item)
	assert.Contains(t, err.Error(), "item not found")
}

func TestSDERepositorySearchItems(t *testing.T) {
	// Arrange
	repo, err := repository.NewSDERepository(testDatabasePath)
	require.NoError(t, err)
	defer repo.Close()

	// Act
	items, err := repo.SearchItems("Tritanium", 5)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, items)
	assert.LessOrEqual(t, len(items), 5)

	// Check that results contain search term
	for _, item := range items {
		assert.Contains(t, item.TypeName, "Tritanium")
		assert.True(t, item.Published)
	}
}

func TestSDERepositoryDatabaseConnection(t *testing.T) {
	// Arrange & Act
	repo, err := repository.NewSDERepository(testDatabasePath)
	require.NoError(t, err)
	defer repo.Close()

	// Assert
	assert.NotNil(t, repo)

	// Test that database is accessible
	err = repo.Ping()
	assert.NoError(t, err)
}
