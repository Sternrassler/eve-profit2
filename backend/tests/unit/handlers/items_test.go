package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"eve-profit2/internal/api/handlers"
	"eve-profit2/internal/models"
	"eve-profit2/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockItemService for testing
type MockItemService struct {
	mock.Mock
}

func (m *MockItemService) GetItemByID(typeID int32) (*models.Item, error) {
	args := m.Called(typeID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Item), args.Error(1)
}

func (m *MockItemService) SearchItems(query string) ([]*models.Item, error) {
	args := m.Called(query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Item), args.Error(1)
}

func TestItemHandlerGetItemDetails(t *testing.T) {
	// Set up
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		itemID         string
		mockSetup      func(*MockItemService)
		expectedStatus int
		expectedError  bool
	}{
		{
			name:   "should return item details for valid item ID",
			itemID: "34", // Tritanium
			mockSetup: func(m *MockItemService) {
				m.On("GetItemByID", int32(34)).Return(&models.Item{
					TypeID:   34,
					TypeName: "Tritanium",
					GroupID:  18,
					Volume:   0.01,
					Mass:     0,
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:   "should return 400 for invalid item ID",
			itemID: "invalid",
			mockSetup: func(m *MockItemService) {
				// No mock setup needed for invalid ID
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name:   "should return 404 for non-existent item",
			itemID: "999999",
			mockSetup: func(m *MockItemService) {
				m.On("GetItemByID", int32(999999)).Return(nil, service.ErrItemNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := &MockItemService{}
			tt.mockSetup(mockService)

			handler := handlers.NewItemHandler(mockService)

			router := gin.New()
			router.GET("/api/v1/items/:item_id", handler.GetItemDetails)

			req := httptest.NewRequest("GET", "/api/v1/items/"+tt.itemID, nil)
			w := httptest.NewRecorder()

			// Act
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectedError {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "data")

				data := response["data"].(map[string]interface{})
				assert.Equal(t, float64(34), data["type_id"])
				assert.Equal(t, "Tritanium", data["type_name"])
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestItemHandlerSearchItems(t *testing.T) {
	// Set up
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		query          string
		mockSetup      func(*MockItemService)
		expectedStatus int
		expectedError  bool
	}{
		{
			name:  "should return search results for valid query",
			query: "Tritanium",
			mockSetup: func(m *MockItemService) {
				items := []*models.Item{
					{
						TypeID:   34,
						TypeName: "Tritanium",
						GroupID:  18,
						Volume:   0.01,
					},
					{
						TypeID:   35,
						TypeName: "Pyerite",
						GroupID:  18,
						Volume:   0.01,
					},
				}
				m.On("SearchItems", "Tritanium").Return(items, nil)
			},
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:  "should return 400 for empty query",
			query: "",
			mockSetup: func(m *MockItemService) {
				// No mock setup needed for empty query
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name:  "should return empty results for no matches",
			query: "NonExistentItem",
			mockSetup: func(m *MockItemService) {
				m.On("SearchItems", "NonExistentItem").Return([]*models.Item{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := &MockItemService{}
			tt.mockSetup(mockService)

			handler := handlers.NewItemHandler(mockService)

			router := gin.New()
			router.GET("/api/v1/items/search", handler.SearchItems)

			req := httptest.NewRequest("GET", "/api/v1/items/search?q="+tt.query, nil)
			w := httptest.NewRecorder()

			// Act
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectedError {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "data")
			}

			mockService.AssertExpectations(t)
		})
	}
}
