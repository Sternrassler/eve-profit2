package cache_test

import (
	"testing"
	"time"

	"eve-profit2/internal/cache"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCacheManagerCreation(t *testing.T) {
	// Act
	cacheManager, err := cache.NewCacheManager()

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, cacheManager)
}

func TestCacheManagerSetAndGetMarketData(t *testing.T) {
	// Arrange
	cacheManager, err := cache.NewCacheManager()
	require.NoError(t, err)

	key := "market:orders:10000002"
	testData := map[string]interface{}{
		"type_id":   34,
		"price":     100.50,
		"volume":    1000,
		"region_id": 10000002,
	}

	// Act
	err = cacheManager.SetMarketData(key, testData, 300*time.Second) // 5 minutes
	require.NoError(t, err)

	// Assert
	var retrievedData map[string]interface{}
	err = cacheManager.GetMarketData(key, &retrievedData)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedData)
}

func TestCacheManagerSetAndGetCharacterData(t *testing.T) {
	// Arrange
	cacheManager, err := cache.NewCacheManager()
	require.NoError(t, err)

	key := "character:assets:123456"
	testData := map[string]interface{}{
		"character_id": 123456,
		"assets":       []string{"item1", "item2"},
	}

	// Act
	err = cacheManager.SetCharacterData(key, testData)
	require.NoError(t, err)

	// Assert
	var retrievedData map[string]interface{}
	err = cacheManager.GetCharacterData(key, &retrievedData)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedData)
}

func TestCacheManagerGetNonExistentData(t *testing.T) {
	// Arrange
	cacheManager, err := cache.NewCacheManager()
	require.NoError(t, err)

	// Act
	var data map[string]interface{}
	err = cacheManager.GetMarketData("non-existent-key", &data)

	// Assert
	assert.Error(t, err) // Should return error for non-existent key
}

func TestCacheManagerSDEData(t *testing.T) {
	// Arrange
	cacheManager, err := cache.NewCacheManager()
	require.NoError(t, err)

	key := "sde:types:34"
	testData := map[string]interface{}{
		"type_id":   34,
		"type_name": "Tritanium",
		"group_id":  18,
	}

	// Act
	err = cacheManager.SetSDEData(key, testData)
	require.NoError(t, err)

	// Assert
	var retrievedData map[string]interface{}
	err = cacheManager.GetSDEData(key, &retrievedData)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedData)
}
