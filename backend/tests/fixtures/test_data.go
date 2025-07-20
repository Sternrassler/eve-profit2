package fixtures

import (
	"time"

	"eve-profit2/internal/models"
)

// TestMarketOrders provides test data for market orders
var TestMarketOrders = []models.MarketOrder{
	{
		OrderID:      123,
		TypeID:       34,       // Tritanium
		LocationID:   60003760, // Jita IV - Moon 4 - Caldari Navy Assembly Plant
		SystemID:     30000142, // Jita
		VolumeTotal:  1000,
		VolumeRemain: 1000,
		MinVolume:    1,
		Price:        5.50,
		IsBuyOrder:   false,
		Duration:     90,
		Issued:       time.Date(2025, 7, 20, 12, 0, 0, 0, time.UTC),
		Range:        "station",
	},
	{
		OrderID:      124,
		TypeID:       34, // Tritanium
		LocationID:   60003760,
		SystemID:     30000142,
		VolumeTotal:  500,
		VolumeRemain: 500,
		MinVolume:    1,
		Price:        5.00,
		IsBuyOrder:   true,
		Duration:     90,
		Issued:       time.Date(2025, 7, 20, 11, 30, 0, 0, time.UTC),
		Range:        "station",
	},
}

// TestItems provides test data for EVE items
var TestItems = []models.Item{
	{
		TypeID:       34,
		TypeName:     "Tritanium",
		GroupID:      18,
		GroupName:    "Mineral",
		CategoryID:   4,
		CategoryName: "Material",
		Volume:       0.01,
		Mass:         1.0,
		Description:  "The most common ore type in the known universe, tritanium is still one of the most useful.",
	},
	{
		TypeID:       35,
		TypeName:     "Pyerite",
		GroupID:      18,
		GroupName:    "Mineral",
		CategoryID:   4,
		CategoryName: "Material",
		Volume:       0.01,
		Mass:         1.0,
		Description:  "Probably the most widely used ore for manufacturing basic technology.",
	},
}

// TestMarketHistory provides test data for market history
var TestMarketHistory = []models.MarketHistory{
	{
		Date:       time.Date(2025, 7, 19, 0, 0, 0, 0, time.UTC),
		Average:    5.25,
		Highest:    5.75,
		Lowest:     4.80,
		OrderCount: 1500,
		Volume:     5000000,
	},
	{
		Date:       time.Date(2025, 7, 18, 0, 0, 0, 0, time.UTC),
		Average:    5.30,
		Highest:    5.80,
		Lowest:     4.85,
		OrderCount: 1600,
		Volume:     5200000,
	},
}
