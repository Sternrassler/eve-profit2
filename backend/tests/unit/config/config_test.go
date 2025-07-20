package config_test

import (
	"os"
	"testing"

	"eve-profit2/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestConfigLoadWithDefaults(t *testing.T) {
	// Arrange
	os.Clearenv() // Clean environment

	// Act
	cfg := config.Load()

	// Assert
	assert.NotNil(t, cfg)
	assert.Equal(t, "9000", cfg.ServerPort)
	assert.Equal(t, "https://esi.evetech.net", cfg.ESIBaseURL)
	assert.Equal(t, 150, cfg.ESIRateLimit)
	assert.True(t, cfg.DebugMode) // Default is true in development
}

func TestConfigLoadWithEnvironmentVariables(t *testing.T) {
	// Arrange
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("ESI_BASE_URL", "https://test.evetech.net")
	os.Setenv("ESI_RATE_LIMIT", "200")
	os.Setenv("DEBUG_MODE", "false")
	defer os.Clearenv()

	// Act
	cfg := config.Load()

	// Assert
	assert.NotNil(t, cfg)
	assert.Equal(t, "8080", cfg.ServerPort)
	assert.Equal(t, "https://test.evetech.net", cfg.ESIBaseURL)
	assert.Equal(t, 200, cfg.ESIRateLimit)
	assert.False(t, cfg.DebugMode)
}

func TestConfigLoadWithESISettings(t *testing.T) {
	// Arrange
	os.Setenv("ESI_CLIENT_ID", "test-client-id")
	os.Setenv("ESI_CLIENT_SECRET", "test-secret")
	os.Setenv("ESI_CALLBACK_URL", "http://localhost:9000/callback")
	defer os.Clearenv()

	// Act
	cfg := config.Load()

	// Assert
	assert.Equal(t, "test-client-id", cfg.ESIClientID)
	assert.Equal(t, "test-secret", cfg.ESIClientSecret)
	assert.Equal(t, "http://localhost:9000/callback", cfg.ESICallbackURL)
}

func TestConfigLoadWithDebugMode(t *testing.T) {
	// Arrange
	os.Setenv("DEBUG_MODE", "true")
	defer os.Clearenv()

	// Act
	cfg := config.Load()

	// Assert
	assert.NotNil(t, cfg)
	assert.True(t, cfg.DebugMode)
}

func TestConfigLoadWithInvalidESIRateLimit(t *testing.T) {
	// Arrange
	os.Setenv("ESI_RATE_LIMIT", "invalid")
	defer os.Clearenv()

	// Act
	cfg := config.Load()

	// Assert
	assert.NotNil(t, cfg)
	assert.Equal(t, 150, cfg.ESIRateLimit) // Should use default
}
