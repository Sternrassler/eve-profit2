package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CharacterHandler struct {
	// characterService *service.CharacterService
}

func NewCharacterHandler(characterService interface{}) *CharacterHandler {
	return &CharacterHandler{}
}

func (h *CharacterHandler) GetCharacterInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Character info endpoint - not implemented yet",
	})
}

func (h *CharacterHandler) GetAssets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Character assets endpoint - not implemented yet",
	})
}

func (h *CharacterHandler) GetWallet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Character wallet endpoint - not implemented yet",
	})
}

func (h *CharacterHandler) GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Character orders endpoint - not implemented yet",
	})
}

func (h *CharacterHandler) GetSkills(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Character skills endpoint - not implemented yet",
	})
}

func (h *CharacterHandler) InitiateLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "EVE SSO login endpoint - not implemented yet",
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
