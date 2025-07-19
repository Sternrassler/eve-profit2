package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
)

type CacheManager struct {
	marketCache    *bigcache.BigCache
	characterCache *bigcache.BigCache
	sdeCache       *bigcache.BigCache
}

type CacheConfig struct {
	Shards             int
	LifeWindow         time.Duration
	CleanWindow        time.Duration
	MaxEntriesInWindow int
	MaxEntrySize       int
	HardMaxCacheSize   int
}

// NewCacheManager creates a new cache manager with optimized configurations
func NewCacheManager() (*CacheManager, error) {
	// Market data cache - short TTL for live data
	marketConfig := bigcache.Config{
		Shards:             1024,
		LifeWindow:         5 * time.Minute,
		CleanWindow:        1 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		HardMaxCacheSize:   256, // 256MB
	}

	marketCache, err := bigcache.New(context.Background(), marketConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create market cache: %w", err)
	}

	// Character data cache - medium TTL
	characterConfig := bigcache.Config{
		Shards:             256,
		LifeWindow:         15 * time.Minute,
		CleanWindow:        2 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       2000,
		HardMaxCacheSize:   128, // 128MB
	}

	characterCache, err := bigcache.New(context.Background(), characterConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create character cache: %w", err)
	}

	// SDE data cache - long TTL for static data
	sdeConfig := bigcache.Config{
		Shards:             512,
		LifeWindow:         24 * time.Hour,
		CleanWindow:        10 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       1000,
		HardMaxCacheSize:   512, // 512MB
	}

	sdeCache, err := bigcache.New(context.Background(), sdeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create SDE cache: %w", err)
	}

	return &CacheManager{
		marketCache:    marketCache,
		characterCache: characterCache,
		sdeCache:       sdeCache,
	}, nil
}

// Market Cache Methods
func (c *CacheManager) SetMarketData(key string, data interface{}, ttl time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.marketCache.Set(key, jsonData)
}

func (c *CacheManager) GetMarketData(key string, dest interface{}) error {
	data, err := c.marketCache.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

// Character Cache Methods
func (c *CacheManager) SetCharacterData(key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.characterCache.Set(key, jsonData)
}

func (c *CacheManager) GetCharacterData(key string, dest interface{}) error {
	data, err := c.characterCache.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

// SDE Cache Methods
func (c *CacheManager) SetSDEData(key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.sdeCache.Set(key, jsonData)
}

func (c *CacheManager) GetSDEData(key string, dest interface{}) error {
	data, err := c.sdeCache.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

// Utility Methods
func (c *CacheManager) DeleteMarketData(key string) error {
	return c.marketCache.Delete(key)
}

func (c *CacheManager) DeleteCharacterData(key string) error {
	return c.characterCache.Delete(key)
}

func (c *CacheManager) DeleteSDEData(key string) error {
	return c.sdeCache.Delete(key)
}

// Cache Statistics
type CacheStats struct {
	MarketStats    bigcache.Stats `json:"market_stats"`
	CharacterStats bigcache.Stats `json:"character_stats"`
	SDEStats       bigcache.Stats `json:"sde_stats"`
}

func (c *CacheManager) GetStats() CacheStats {
	return CacheStats{
		MarketStats:    c.marketCache.Stats(),
		CharacterStats: c.characterCache.Stats(),
		SDEStats:       c.sdeCache.Stats(),
	}
}

// Reset all caches
func (c *CacheManager) Reset() error {
	if err := c.marketCache.Reset(); err != nil {
		return err
	}
	if err := c.characterCache.Reset(); err != nil {
		return err
	}
	if err := c.sdeCache.Reset(); err != nil {
		return err
	}
	return nil
}

// Close all caches
func (c *CacheManager) Close() error {
	if err := c.marketCache.Close(); err != nil {
		return err
	}
	if err := c.characterCache.Close(); err != nil {
		return err
	}
	if err := c.sdeCache.Close(); err != nil {
		return err
	}
	return nil
}
