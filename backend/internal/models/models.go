package models

import "time"

// MarketOrder represents a market order from ESI
type MarketOrder struct {
	OrderID      int64     `json:"order_id"`
	TypeID       int32     `json:"type_id"`
	LocationID   int64     `json:"location_id"`
	SystemID     int32     `json:"system_id"`
	VolumeTotal  int32     `json:"volume_total"`
	VolumeRemain int32     `json:"volume_remain"`
	MinVolume    int32     `json:"min_volume"`
	Price        float64   `json:"price"`
	IsBuyOrder   bool      `json:"is_buy_order"`
	Duration     int32     `json:"duration"`
	Issued       time.Time `json:"issued"`
	Range        string    `json:"range"`
}

// MarketHistory represents historical market data
type MarketHistory struct {
	Date       time.Time `json:"date"`
	Average    float64   `json:"average"`
	Highest    float64   `json:"highest"`
	Lowest     float64   `json:"lowest"`
	OrderCount int64     `json:"order_count"`
	Volume     int64     `json:"volume"`
}

// ItemPrice represents current market prices for an item
type ItemPrice struct {
	TypeID      int32     `json:"type_id"`
	BuyMax      float64   `json:"buy_max"`
	SellMin     float64   `json:"sell_min"`
	BuyVolume   int64     `json:"buy_volume"`
	SellVolume  int64     `json:"sell_volume"`
	LastUpdated time.Time `json:"last_updated"`
}

// Item represents an EVE item from SDE
type Item struct {
	TypeID       int32   `json:"type_id" db:"typeID"`
	TypeName     string  `json:"type_name" db:"typeName"`
	GroupID      int32   `json:"group_id" db:"groupID"`
	GroupName    string  `json:"group_name"`
	CategoryID   int32   `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Volume       float64 `json:"volume" db:"volume"`
	Mass         float64 `json:"mass" db:"mass"`
	Description  string  `json:"description" db:"description"`
}

// Station represents a station/structure from SDE
type Station struct {
	StationID   int64   `json:"station_id" db:"stationID"`
	StationName string  `json:"station_name" db:"stationName"`
	SystemID    int32   `json:"system_id" db:"solarSystemID"`
	SystemName  string  `json:"system_name"`
	RegionID    int32   `json:"region_id"`
	RegionName  string  `json:"region_name"`
	Security    float64 `json:"security"`
}

// ProfitRoute represents a trading route calculation
type ProfitRoute struct {
	FromStation  Station `json:"from_station"`
	ToStation    Station `json:"to_station"`
	Item         Item    `json:"item"`
	BuyPrice     float64 `json:"buy_price"`
	SellPrice    float64 `json:"sell_price"`
	Profit       float64 `json:"profit"`
	ProfitMargin float64 `json:"profit_margin"`
	Volume       int64   `json:"volume"`
	Investment   float64 `json:"investment"`
}

// Character represents EVE character data
type Character struct {
	CharacterID    int32     `json:"character_id"`
	CharacterName  string    `json:"character_name"`
	CorporationID  int32     `json:"corporation_id"`
	AllianceID     int32     `json:"alliance_id,omitempty"`
	Birthday       time.Time `json:"birthday"`
	SecurityStatus float64   `json:"security_status"`
}

// CharacterAsset represents character assets
type CharacterAsset struct {
	ItemID       int64  `json:"item_id"`
	TypeID       int32  `json:"type_id"`
	LocationID   int64  `json:"location_id"`
	LocationFlag string `json:"location_flag"`
	LocationType string `json:"location_type"`
	Quantity     int32  `json:"quantity"`
	IsSingleton  bool   `json:"is_singleton"`
}

// CharacterWallet represents wallet balance
type CharacterWallet struct {
	Balance float64 `json:"balance"`
}

// CharacterOrder represents character market orders
type CharacterOrder struct {
	OrderID      int64     `json:"order_id"`
	TypeID       int32     `json:"type_id"`
	RegionID     int32     `json:"region_id"`
	LocationID   int64     `json:"location_id"`
	Range        string    `json:"range"`
	IsBuyOrder   bool      `json:"is_buy_order"`
	Price        float64   `json:"price"`
	VolumeTotal  int32     `json:"volume_total"`
	VolumeRemain int32     `json:"volume_remain"`
	Issued       time.Time `json:"issued"`
	State        string    `json:"state"`
	MinVolume    int32     `json:"min_volume"`
	Duration     int32     `json:"duration"`
	Escrow       float64   `json:"escrow,omitempty"`
}

// CharacterSkill represents character skills
type CharacterSkill struct {
	SkillID            int32 `json:"skill_id"`
	TrainedSkillLevel  int32 `json:"trained_skill_level"`
	SkillpointsInSkill int64 `json:"skillpoints_in_skill"`
	ActiveSkillLevel   int32 `json:"active_skill_level"`
}

// AuthToken represents OAuth tokens
type AuthToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
	CharacterID  int32     `json:"character_id"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

// CacheEntry represents a cached item
type CacheEntry struct {
	Key       string    `json:"key"`
	Value     []byte    `json:"value"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
