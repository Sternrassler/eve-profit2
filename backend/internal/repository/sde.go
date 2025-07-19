package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// SDEItem represents an item from SDE
type SDEItem struct {
	TypeID      int32   `json:"typeId"`
	TypeName    string  `json:"typeName"`
	GroupID     int32   `json:"groupId"`
	Volume      float64 `json:"volume"`
	MarketGroup int32   `json:"marketGroupID,omitempty"`
	Published   bool    `json:"published"`
}

// SDEStation represents a station from SDE
type SDEStation struct {
	StationID   int64  `json:"stationId"`
	StationName string `json:"stationName"`
	SystemID    int32  `json:"systemId"`
	RegionID    int32  `json:"regionId"`
	TypeID      int32  `json:"typeId"`
}

// SDERegion represents a region from SDE
type SDERegion struct {
	RegionID   int32  `json:"regionId"`
	RegionName string `json:"regionName"`
}

// SDERepository handles SDE SQLite database operations
type SDERepository struct {
	db *sql.DB
}

func NewSDERepository(dbPath string) (*SDERepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &SDERepository{db: db}, nil
}

func (r *SDERepository) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}

func (r *SDERepository) Ping() error {
	if r.db == nil {
		return fmt.Errorf("database connection is nil")
	}
	return r.db.Ping()
}

// GetItemByID retrieves a single item by TypeID
func (r *SDERepository) GetItemByID(typeID int32) (*SDEItem, error) {
	query := `
		SELECT typeID, typeName, groupID, volume, marketGroupID, published
		FROM invTypes 
		WHERE typeID = ?
	`

	var item SDEItem
	var marketGroup sql.NullInt32

	err := r.db.QueryRow(query, typeID).Scan(
		&item.TypeID,
		&item.TypeName,
		&item.GroupID,
		&item.Volume,
		&marketGroup,
		&item.Published,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("item not found: typeID %d", typeID)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	if marketGroup.Valid {
		item.MarketGroup = marketGroup.Int32
	}

	return &item, nil
}

// SearchItems searches for items by name
func (r *SDERepository) SearchItems(searchTerm string, limit int) ([]*SDEItem, error) {
	query := `
		SELECT typeID, typeName, groupID, volume, marketGroupID, published
		FROM invTypes 
		WHERE typeName LIKE ? AND published = 1
		ORDER BY typeName
		LIMIT ?
	`

	rows, err := r.db.Query(query, "%"+searchTerm+"%", limit)
	if err != nil {
		return nil, fmt.Errorf("failed to search items: %w", err)
	}
	defer rows.Close()

	var items []*SDEItem
	for rows.Next() {
		var item SDEItem
		var marketGroup sql.NullInt32

		err := rows.Scan(
			&item.TypeID,
			&item.TypeName,
			&item.GroupID,
			&item.Volume,
			&marketGroup,
			&item.Published,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}

		if marketGroup.Valid {
			item.MarketGroup = marketGroup.Int32
		}

		items = append(items, &item)
	}

	return items, nil
}

// GetStationsBySystem retrieves all stations in a system
func (r *SDERepository) GetStationsBySystem(systemID int32) ([]*SDEStation, error) {
	query := `
		SELECT stationID, stationName, solarSystemID, regionID, stationTypeID
		FROM staStations 
		WHERE solarSystemID = ?
	`

	rows, err := r.db.Query(query, systemID)
	if err != nil {
		return nil, fmt.Errorf("failed to get stations: %w", err)
	}
	defer rows.Close()

	var stations []*SDEStation
	for rows.Next() {
		var station SDEStation

		err := rows.Scan(
			&station.StationID,
			&station.StationName,
			&station.SystemID,
			&station.RegionID,
			&station.TypeID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan station: %w", err)
		}

		stations = append(stations, &station)
	}

	return stations, nil
}

// GetRegions retrieves all regions with limit
func (r *SDERepository) GetRegions(limit int) ([]*SDERegion, error) {
	query := `
		SELECT regionID, regionName
		FROM mapRegions
		ORDER BY regionName
		LIMIT ?
	`

	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get regions: %w", err)
	}
	defer rows.Close()

	var regions []*SDERegion
	for rows.Next() {
		var region SDERegion

		err := rows.Scan(
			&region.RegionID,
			&region.RegionName,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan region: %w", err)
		}

		regions = append(regions, &region)
	}

	return regions, nil
}
