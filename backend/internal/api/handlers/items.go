package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"eve-profit2/internal/models"
	"eve-profit2/internal/service"

	"github.com/gin-gonic/gin"
)

// ItemServiceInterface defines the contract for item operations
type ItemServiceInterface interface {
	GetItemByID(typeID int32) (*models.Item, error)
	SearchItems(query string) ([]*models.Item, error)
}

type ItemHandler struct {
	itemService ItemServiceInterface
}

func NewItemHandler(itemService ItemServiceInterface) *ItemHandler {
	return &ItemHandler{
		itemService: itemService,
	}
}

type ItemResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func (h *ItemHandler) GetItemDetails(c *gin.Context) {
	itemIDStr := c.Param("item_id")

	// Validate item ID
	itemID, err := strconv.ParseInt(itemIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ItemResponse{
			Success: false,
			Error:   "Invalid item ID format",
		})
		return
	}

	// Get item from service
	item, err := h.itemService.GetItemByID(int32(itemID))
	if err != nil {
		if errors.Is(err, service.ErrItemNotFound) {
			c.JSON(http.StatusNotFound, ItemResponse{
				Success: false,
				Error:   "Item not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, ItemResponse{
			Success: false,
			Error:   "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, ItemResponse{
		Success: true,
		Data:    item,
	})
}

func (h *ItemHandler) SearchItems(c *gin.Context) {
	query := c.Query("q")

	// Validate query parameter
	if query == "" {
		c.JSON(http.StatusBadRequest, ItemResponse{
			Success: false,
			Error:   "Search query is required",
		})
		return
	}

	// Search items using service
	items, err := h.itemService.SearchItems(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ItemResponse{
			Success: false,
			Error:   "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, ItemResponse{
		Success: true,
		Data:    items,
	})
}

func (h *ItemHandler) GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Categories endpoint - not implemented yet",
	})
}
