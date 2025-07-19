package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	// itemService *service.ItemService
}

func NewItemHandler(itemService interface{}) *ItemHandler {
	return &ItemHandler{}
}

func (h *ItemHandler) SearchItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Item search endpoint - not implemented yet",
		"query":   c.Query("q"),
	})
}

func (h *ItemHandler) GetItemDetails(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Item details endpoint - not implemented yet",
		"item_id": c.Param("item_id"),
	})
}

func (h *ItemHandler) GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Categories endpoint - not implemented yet",
	})
}
