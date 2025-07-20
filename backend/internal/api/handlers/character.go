package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Error message constants following DRY principle
const (
	ErrInvalidCharacterID = "Invalid character ID"
	ErrCharacterInfo      = "Failed to get character info"
	ErrCharacterAssets    = "Failed to get character assets"
	ErrCharacterWallet    = "Failed to get character wallet"
	ErrCharacterOrders    = "Failed to get character orders"
	ErrCharacterSkills    = "Failed to get character skills"
)

// CharacterService defines the interface for character operations
type CharacterService interface {
	GetCharacterInfo(characterID int32) (interface{}, error)
	GetCharacterAssets(characterID int32) (interface{}, error)
	GetCharacterWallet(characterID int32) (interface{}, error)
	GetCharacterOrders(characterID int32) (interface{}, error)
	GetCharacterSkills(characterID int32) (interface{}, error)
}

type CharacterHandler struct {
	characterService CharacterService
}

func NewCharacterHandler(characterService interface{}) *CharacterHandler {
	if service, ok := characterService.(CharacterService); ok {
		return &CharacterHandler{characterService: service}
	}
	// Return with nil service for now - Phase 4 implementation will add actual service
	return &CharacterHandler{}
}

// GetCharacterInfo retrieves character information by ID
func (h *CharacterHandler) GetCharacterInfo(c *gin.Context) {
	characterID, err := h.extractCharacterIDFromPath(c)
	if err != nil {
		h.respondWithError(c, http.StatusBadRequest, ErrInvalidCharacterID, err)
		return
	}

	if h.characterService == nil {
		h.respondWithNotImplemented(c, "Character info endpoint")
		return
	}

	data, err := h.characterService.GetCharacterInfo(characterID)
	if err != nil {
		h.respondWithError(c, http.StatusInternalServerError, ErrCharacterInfo, err)
		return
	}

	h.respondWithSuccess(c, data)
}

// GetAssets retrieves character assets
func (h *CharacterHandler) GetAssets(c *gin.Context) {
	characterID, err := h.extractCharacterIDFromPath(c)
	if err != nil {
		h.respondWithError(c, http.StatusBadRequest, ErrInvalidCharacterID, err)
		return
	}

	if h.characterService == nil {
		h.respondWithNotImplemented(c, "Character assets endpoint")
		return
	}

	data, err := h.characterService.GetCharacterAssets(characterID)
	if err != nil {
		h.respondWithError(c, http.StatusInternalServerError, ErrCharacterAssets, err)
		return
	}

	h.respondWithSuccess(c, data)
}

// GetWallet retrieves character wallet information
func (h *CharacterHandler) GetWallet(c *gin.Context) {
	characterID, err := h.extractCharacterIDFromPath(c)
	if err != nil {
		h.respondWithError(c, http.StatusBadRequest, ErrInvalidCharacterID, err)
		return
	}

	if h.characterService == nil {
		h.respondWithNotImplemented(c, "Character wallet endpoint")
		return
	}

	data, err := h.characterService.GetCharacterWallet(characterID)
	if err != nil {
		h.respondWithError(c, http.StatusInternalServerError, ErrCharacterWallet, err)
		return
	}

	h.respondWithSuccess(c, data)
}

// GetOrders retrieves character market orders
func (h *CharacterHandler) GetOrders(c *gin.Context) {
	characterID, err := h.extractCharacterIDFromPath(c)
	if err != nil {
		h.respondWithError(c, http.StatusBadRequest, ErrInvalidCharacterID, err)
		return
	}

	if h.characterService == nil {
		h.respondWithNotImplemented(c, "Character orders endpoint")
		return
	}

	data, err := h.characterService.GetCharacterOrders(characterID)
	if err != nil {
		h.respondWithError(c, http.StatusInternalServerError, ErrCharacterOrders, err)
		return
	}

	h.respondWithSuccess(c, data)
}

// GetSkills retrieves character skills
func (h *CharacterHandler) GetSkills(c *gin.Context) {
	characterID, err := h.extractCharacterIDFromPath(c)
	if err != nil {
		h.respondWithError(c, http.StatusBadRequest, ErrInvalidCharacterID, err)
		return
	}

	if h.characterService == nil {
		h.respondWithNotImplemented(c, "Character skills endpoint")
		return
	}

	data, err := h.characterService.GetCharacterSkills(characterID)
	if err != nil {
		h.respondWithError(c, http.StatusInternalServerError, ErrCharacterSkills, err)
		return
	}

	h.respondWithSuccess(c, data)
}

// InitiateLogin starts the EVE SSO login flow
func (h *CharacterHandler) InitiateLogin(c *gin.Context) {
	h.respondWithNotImplemented(c, "EVE SSO login endpoint")
}

// Helper methods following DRY principle

// extractCharacterIDFromPath extracts and validates character ID from URL path
func (h *CharacterHandler) extractCharacterIDFromPath(c *gin.Context) (int32, error) {
	characterIDStr := c.Param("characterID")
	if characterIDStr == "" {
		return 0, fmt.Errorf("character ID is required")
	}

	characterID, err := strconv.ParseInt(characterIDStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid character ID format: %w", err)
	}

	if characterID <= 0 {
		return 0, fmt.Errorf("character ID must be positive")
	}

	return int32(characterID), nil
}

// respondWithError sends a standardized error response
func (h *CharacterHandler) respondWithError(c *gin.Context, statusCode int, message string, err error) {
	c.JSON(statusCode, gin.H{
		"error":   message,
		"details": err.Error(),
	})
}

// respondWithSuccess sends a standardized success response
func (h *CharacterHandler) respondWithSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

// respondWithNotImplemented sends a standardized not implemented response
func (h *CharacterHandler) respondWithNotImplemented(c *gin.Context, endpoint string) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": endpoint + " - not implemented yet (Phase 4)",
		"phase":   "Phase 4 - API Handlers Implementation",
	})
}

func (h *CharacterHandler) HandleCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "EVE SSO callback endpoint - not implemented yet",
	})
}

func (h *CharacterHandler) RefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Token refresh endpoint - not implemented yet",
	})
}
