package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds all application configuration
type Config struct {
	// Server Configuration
	ServerPort string
	APIBaseURL string
	CORSOrigin string
	DebugMode  bool
	LogLevel   string

	// EVE ESI Configuration
	ESIClientID     string
	ESIClientSecret string
	ESICallbackURL  string
	ESIBaseURL      string
	ESIScopes       []string

	// EVE SSO Configuration
	EVESSOBaseURL      string
	EVESSOAuthorizeURL string
	EVESSOTokenURL     string
	EVESSOVerifyURL    string

	// Rate Limiting
	ESIRateLimit   int
	ESIBurstLimit  int
	ESITimeoutSecs int

	// Caching Configuration
	CacheTTLMarketOrders  time.Duration
	CacheTTLMarketHistory time.Duration
	CacheTTLTypeInfo      time.Duration
	CacheTTLCharacterInfo time.Duration

	// Database Configuration
	SDEDatabasePath string
}

// Load reads configuration from environment variables with sensible defaults
func Load() *Config {
	return &Config{
		// Server Configuration
		ServerPort: getEnv("SERVER_PORT", "9000"),
		APIBaseURL: getEnv("API_BASE_URL", "http://localhost:9000"),
		CORSOrigin: getEnv("CORS_ORIGIN", "http://localhost:3000"),
		DebugMode:  getEnvBool("DEBUG_MODE", true),
		LogLevel:   getEnv("LOG_LEVEL", "info"),

		// EVE ESI Configuration
		ESIClientID:     getEnv("ESI_CLIENT_ID", "0928b4bcd20242aeb9b8be10f5451094"),
		ESIClientSecret: getEnv("ESI_CLIENT_SECRET", "AQPjLZ3VYAewR59J5jStZs52dY7jISGVLwXv5NA"),
		ESICallbackURL:  getEnv("ESI_CALLBACK_URL", "http://localhost:9000/callback"),
		ESIBaseURL:      getEnv("ESI_BASE_URL", "https://esi.evetech.net"),
		ESIScopes: getEnvSlice("ESI_SCOPES", []string{
			"publicData",
			"esi-location.read_location.v1",
			"esi-location.read_ship_type.v1",
			"esi-skills.read_skills.v1",
			"esi-wallet.read_character_wallet.v1",
			"esi-universe.read_structures.v1",
			"esi-assets.read_assets.v1",
			"esi-fittings.read_fittings.v1",
			"esi-characters.read_standings.v1",
		}),

		// EVE SSO Configuration
		EVESSOBaseURL:      getEnv("EVE_SSO_BASE_URL", "https://login.eveonline.com"),
		EVESSOAuthorizeURL: getEnv("EVE_SSO_AUTHORIZE_URL", "https://login.eveonline.com/v2/oauth/authorize"),
		EVESSOTokenURL:     getEnv("EVE_SSO_TOKEN_URL", "https://login.eveonline.com/v2/oauth/token"),
		EVESSOVerifyURL:    getEnv("EVE_SSO_VERIFY_URL", "https://login.eveonline.com/oauth/verify"),

		// Rate Limiting
		ESIRateLimit:   getEnvInt("ESI_RATE_LIMIT", 150),
		ESIBurstLimit:  getEnvInt("ESI_BURST_LIMIT", 400),
		ESITimeoutSecs: getEnvInt("ESI_TIMEOUT_SECONDS", 30),

		// Caching Configuration
		CacheTTLMarketOrders:  time.Duration(getEnvInt("CACHE_TTL_MARKET_ORDERS", 300)) * time.Second,
		CacheTTLMarketHistory: time.Duration(getEnvInt("CACHE_TTL_MARKET_HISTORY", 3600)) * time.Second,
		CacheTTLTypeInfo:      time.Duration(getEnvInt("CACHE_TTL_TYPE_INFO", 86400)) * time.Second,
		CacheTTLCharacterInfo: time.Duration(getEnvInt("CACHE_TTL_CHARACTER_INFO", 1800)) * time.Second,

		// Database Configuration
		SDEDatabasePath: getEnv("SDE_DATABASE_PATH", "./data/sqlite-latest.sqlite"),
	}
}

// Helper functions for reading environment variables
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
		// Log warning for invalid values but continue with default
		fmt.Printf("Warning: Invalid integer value for %s: %s, using default: %d\n", key, value, defaultValue)
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
		// Log warning for invalid values but continue with default
		fmt.Printf("Warning: Invalid boolean value for %s: %s, using default: %t\n", key, value, defaultValue)
	}
	return defaultValue
}

func getEnvSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Fields(value) // Split by whitespace
	}
	return defaultValue
}

// GetServerAddress returns the full server address
func (c *Config) GetServerAddress() string {
	return ":" + c.ServerPort
}

// GetESIScopesString returns ESI scopes as space-separated string
func (c *Config) GetESIScopesString() string {
	return strings.Join(c.ESIScopes, " ")
}
