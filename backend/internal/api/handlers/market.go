package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MarketHandler struct {
	// marketService *service.MarketService
}

func NewMarketHandler(marketService interface{}) *MarketHandler {
	return &MarketHandler{}
}

func (h *MarketHandler) GetItemPrices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Market prices endpoint - not implemented yet",
		"item_id": c.Param("item_id"),
	})
}

func (h *MarketHandler) GetItemOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Market orders endpoint - not implemented yet",
		"item_id": c.Param("item_id"),
	})
}

func (h *MarketHandler) GetPriceHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Price history endpoint - not implemented yet",
		"item_id": c.Param("item_id"),
	})
}
