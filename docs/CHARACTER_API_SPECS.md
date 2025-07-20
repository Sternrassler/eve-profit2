# Character API Integration - EVE Profit Calculator 2.0

## 🔐 EVE SSO OAuth Integration (Phase 4)

**WICHTIG:** Diese Datei ist für Phase 4 - API Handlers Implementation.

### OAuth Flow Overview
```
1. User clicks "Login with EVE" → Frontend redirects to EVE SSO
2. User authorizes app → EVE SSO redirects to callback
3. Backend exchanges code for access/refresh tokens
4. Backend validates token and fetches character info
5. JWT issued to frontend with character session
```

### EVE Application Configuration (Production-Ready)
- **Client ID:** `0928b4bcd20242aeb9b8be10f5451094`
- **Client Secret:** `AQPjLZ3VYAewR59J5jStZs52dY7jISGVLwXv5NA`
- **Callback URL:** `http://localhost:9000/callback`

### Required Scopes
```go
var RequiredScopes = []string{
    "esi-skills.read_skills.v1",           // Character Skills
    "esi-assets.read_assets.v1",           // Character Assets  
    "esi-wallet.read_character_wallet.v1", // Wallet Balance
    "esi-location.read_location.v1",       // Current Location
    "esi-location.read_ship_type.v1",      // Current Ship
}
```

## 📊 Clean Code Character Models

### Domain Models (TDD Implementation)
```go
// Character represents an EVE Online character
type Character struct {
    ID              int32   `json:"characterId"`
    Name            string  `json:"characterName"`
    CorporationID   int32   `json:"corporationId"`
    SecurityStatus  float64 `json:"securityStatus"`
    WalletBalance   float64 `json:"walletBalance"`
    Skills          []Skill `json:"skills"`
}

// Skill represents a character skill
type Skill struct {
    SkillID          int32 `json:"skillId"`
    ActiveSkillLevel int32 `json:"activeSkillLevel"`
    SkillPointsInSkill int64 `json:"skillPointsInSkill"`
}

// TradingSkills contains relevant trading skills
type TradingSkills struct {
    BrokerRelations     int32 `json:"brokerRelations"`     // Reduces broker fees
    Accounting          int32 `json:"accounting"`          // Reduces sales tax
    MarketingSkill      int32 `json:"marketing"`           // Increases market order range
}
```

## 🧪 TDD Implementation Plan (Phase 4)

### Test-First Development for Character API:
```go
// 1. 🔴 RED: Write failing tests first
func TestCharacterHandler_AuthCallback_WithValidCode_ShouldReturnJWT(t *testing.T)
func TestCharacterService_GetCharacterSkills_WithValidToken_ShouldReturnSkills(t *testing.T)
func TestOAuthClient_ExchangeCodeForToken_WithValidCode_ShouldReturnToken(t *testing.T)

// 2. 🟢 GREEN: Implement minimal functionality
// 3. 🔄 REFACTOR: Apply Clean Code principles
```

### API Endpoints to Implement (TDD):
- `POST /api/v1/auth/login` - EVE SSO Login redirect
- `GET /api/v1/auth/callback` - OAuth callback handler
- `GET /api/v1/character/profile` - Character information
- `GET /api/v1/character/skills` - Character skills (trading relevant)
- `GET /api/v1/character/wallet` - Wallet balance
- `POST /api/v1/auth/logout` - Session cleanup

---

**Phase 4 Ready:** All EVE SSO configuration is complete. Start TDD implementation of Character API handlers.
    Birthday        time.Time `json:"birthday"`
    LastLogin       time.Time `json:"lastLogin"`
    TotalSP         int64     `json:"totalSp"`
}

type CharacterSkills struct {
    Skills      []Skill `json:"skills"`
    TotalSP     int64   `json:"totalSp"`
    UnallocSP   int64   `json:"unallocatedSp"`
}

type Skill struct {
    SkillID           int32 `json:"skillId"`
    SkillName         string `json:"skillName"`        // From SDE
    TrainedSkillLevel int32 `json:"trainedSkillLevel"`
    SkillpointsInSkill int64 `json:"skillpointsInSkill"`
    ActiveSkillLevel  int32 `json:"activeSkillLevel"`
}
```

### Character Assets
```go
type Asset struct {
    ItemID       int64  `json:"itemId"`
    TypeID       int32  `json:"typeId"`
    TypeName     string `json:"typeName"`        // From SDE
    LocationID   int64  `json:"locationId"`
    LocationName string `json:"locationName"`    // From SDE/ESI
    LocationFlag string `json:"locationFlag"`    // Hangar, CargoHold, etc.
    Quantity     int32  `json:"quantity"`
    IsSingleton  bool   `json:"isSingleton"`
    IsBPC        bool   `json:"isBlueprintCopy"`
}

type Ship struct {
    ItemID     int64  `json:"itemId"`
    TypeID     int32  `json:"typeId"`
    ShipName   string `json:"shipName"`        // From SDE
    LocationID int64  `json:"locationId"`
    LocationName string `json:"locationName"`
    CargoHold  float64 `json:"cargoHold"`      // m³
    MaxCargo   float64 `json:"maxCargo"`       // m³ from SDE
}

type WalletBalance struct {
    Balance   float64   `json:"balance"`       // ISK
    LastUpdate time.Time `json:"lastUpdate"`
}
```

### Character Market Orders
```go
type CharacterOrder struct {
    OrderID      int64     `json:"orderId"`
    TypeID       int32     `json:"typeId"`
    TypeName     string    `json:"typeName"`     // From SDE
    RegionID     int32     `json:"regionId"`
    LocationID   int64     `json:"locationId"`
    LocationName string    `json:"locationName"` // From SDE
    IsBuyOrder   bool      `json:"isBuyOrder"`
    Price        float64   `json:"price"`
    VolumeTotal  int32     `json:"volumeTotal"`
    VolumeRemain int32     `json:"volumeRemain"`
    MinVolume    int32     `json:"minVolume"`
    Duration     int32     `json:"duration"`
    Issued       time.Time `json:"issued"`
    Range        string    `json:"range"`
    State        string    `json:"state"`        // open, closed, expired
}
```

## 🚀 Service Layer Implementation

### Character Service
```go
// internal/service/character.go
type CharacterService struct {
    esiClient *eve.ESIClient
    sdeClient *sde.SDEClient
    cache     *bigcache.BigCache
    logger    zerolog.Logger
}

func (s *CharacterService) GetCharacterInfo(characterID int32, token string) (*Character, error) {
    // 1. Check cache first
    cacheKey := fmt.Sprintf("char:info:%d", characterID)
    if cached, err := s.cache.Get(cacheKey); err == nil {
        var char Character
        json.Unmarshal(cached, &char)
        return &char, nil
    }
    
    // 2. Fetch from ESI
    esiChar, err := s.esiClient.GetCharacterInfo(characterID, token)
    if err != nil {
        return nil, err
    }
    
    // 3. Enrich with additional data
    char := &Character{
        CharacterID:    esiChar.CharacterID,
        CharacterName:  esiChar.Name,
        CorporationID:  esiChar.CorporationID,
        AllianceID:     esiChar.AllianceID,
        SecurityStatus: esiChar.SecurityStatus,
        Birthday:       esiChar.Birthday,
    }
    
    // 4. Cache result (1 hour TTL)
    if data, err := json.Marshal(char); err == nil {
        s.cache.Set(cacheKey, data)
    }
    
    return char, nil
}

func (s *CharacterService) GetCharacterSkills(characterID int32, token string) (*CharacterSkills, error) {
    // Similar pattern: Cache -> ESI -> SDE enrichment -> Cache
}

func (s *CharacterService) GetCharacterAssets(characterID int32, token string) ([]Asset, error) {
    // 1. Fetch assets from ESI
    esiAssets, err := s.esiClient.GetCharacterAssets(characterID, token)
    if err != nil {
        return nil, err
    }
    
    // 2. Enrich with SDE data (type names, etc.)
    assets := make([]Asset, len(esiAssets))
    for i, esiAsset := range esiAssets {
        // Get type name from SDE
        item, _ := s.sdeClient.GetItemByID(esiAsset.TypeID)
        
        assets[i] = Asset{
            ItemID:      esiAsset.ItemID,
            TypeID:      esiAsset.TypeID,
            TypeName:    item.TypeName,
            LocationID:  esiAsset.LocationID,
            Quantity:    esiAsset.Quantity,
            IsSingleton: esiAsset.IsSingleton,
        }
    }
    
    return assets, nil
}

func (s *CharacterService) GetCharacterShips(characterID int32, token string) ([]Ship, error) {
    // Filter assets for ships, add cargo capacity from SDE
}
```

## 🔧 ESI Client Character Methods

### OAuth-Protected ESI Calls
```go
// pkg/eve/client.go
func (c *ESIClient) GetCharacterInfo(characterID int32, token string) (*ESICharacter, error) {
    url := fmt.Sprintf("https://esi.evetech.net/latest/characters/%d/", characterID)
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("User-Agent", "eve-profit-calculator/1.0")
    
    resp, err := c.doRequest(req)
    if err != nil {
        return nil, err
    }
    
    var char ESICharacter
    err = json.Unmarshal(resp, &char)
    return &char, err
}

func (c *ESIClient) GetCharacterSkills(characterID int32, token string) (*ESISkills, error) {
    url := fmt.Sprintf("https://esi.evetech.net/latest/characters/%d/skills/", characterID)
    return c.makeAuthenticatedRequest(url, token, &ESISkills{})
}

func (c *ESIClient) GetCharacterAssets(characterID int32, token string) ([]ESIAsset, error) {
    url := fmt.Sprintf("https://esi.evetech.net/latest/characters/%d/assets/", characterID)
    return c.makeAuthenticatedRequest(url, token, &[]ESIAsset{})
}

func (c *ESIClient) GetCharacterWallet(characterID int32, token string) (float64, error) {
    url := fmt.Sprintf("https://esi.evetech.net/latest/characters/%d/wallet/", characterID)
    var balance float64
    _, err := c.makeAuthenticatedRequest(url, token, &balance)
    return balance, err
}

func (c *ESIClient) GetCharacterOrders(characterID int32, token string) ([]ESIOrder, error) {
    url := fmt.Sprintf("https://esi.evetech.net/latest/characters/%d/orders/", characterID)
    return c.makeAuthenticatedRequest(url, token, &[]ESIOrder{})
}
```

## 🔐 JWT Token Management

### JWT Session Handling
```go
// internal/service/auth.go
type AuthService struct {
    oauthClient *eve.EVEAuthClient
    jwtSecret   []byte
    logger      zerolog.Logger
}

type JWTClaims struct {
    CharacterID   int32  `json:"character_id"`
    CharacterName string `json:"character_name"`
    AccessToken   string `json:"-"`           // Don't include in JWT
    RefreshToken  string `json:"-"`           // Store securely server-side
    jwt.RegisteredClaims
}

func (s *AuthService) CreateJWTToken(characterID int32, characterName string) (string, error) {
    claims := JWTClaims{
        CharacterID:   characterID,
        CharacterName: characterName,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "eve-profit-calculator",
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(s.jwtSecret)
}

func (s *AuthService) ValidateJWTToken(tokenString string) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return s.jwtSecret, nil
    })
    
    if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, err
}
```

## 🎯 Trading-Optimierungen mit Character Data

### Skill-basierte Berechnungen
```go
// Trading-relevante Skills für Berechnungen
const (
    BrokerRelationsSkillID = 3446  // Reduziert Broker Fees
    AccountingSkillID     = 16622  // Reduziert Sales Tax
    MarginTradingSkillID  = 16597  // Ermöglicht Buy Orders mit weniger ISK
    DaytradingSkillID     = 16598  // Reduziert Modify Order Range
)

func (s *CharacterService) CalculateTradingFees(characterID int32, token string) (*TradingFees, error) {
    skills, err := s.GetCharacterSkills(characterID, token)
    if err != nil {
        return nil, err
    }
    
    // Standard Fees
    brokerFee := 0.03   // 3%
    salesTax := 0.08    // 8%
    
    // Skill-basierte Reduktionen
    for _, skill := range skills.Skills {
        switch skill.SkillID {
        case BrokerRelationsSkillID:
            brokerFee -= float64(skill.TrainedSkillLevel) * 0.001 // -0.1% pro Level
        case AccountingSkillID:
            salesTax -= float64(skill.TrainedSkillLevel) * 0.001  // -0.1% pro Level
        }
    }
    
    return &TradingFees{
        BrokerFee: brokerFee,
        SalesTax:  salesTax,
    }, nil
}
```

### Asset-basierte Transportberechnung
```go
func (s *CharacterService) GetAvailableCargoCapacity(characterID int32, token string, locationID int64) (float64, error) {
    ships, err := s.GetCharacterShips(characterID, token)
    if err != nil {
        return 0, err
    }
    
    // Find ships at location
    var totalCargo float64
    for _, ship := range ships {
        if ship.LocationID == locationID {
            totalCargo += ship.MaxCargo - ship.CargoHold // Available space
        }
    }
    
    return totalCargo, nil
}
```

## 🔄 Caching Strategy für Character Data

### Cache TTL Configuration
```go
const (
    CharacterInfoTTL  = 1 * time.Hour     // Character basic info
    CharacterSkillsTTL = 6 * time.Hour    // Skills (change rarely)
    CharacterAssetsTTL = 5 * time.Minute  // Assets (can change quickly)
    CharacterWalletTTL = 1 * time.Minute  // Wallet balance
    CharacterOrdersTTL = 2 * time.Minute  // Market orders
)
```

### Cache Keys
```
Character Info:    "char:info:{characterId}"
Character Skills:  "char:skills:{characterId}"
Character Assets:  "char:assets:{characterId}"
Character Wallet:  "char:wallet:{characterId}"
Character Orders:  "char:orders:{characterId}"
Trading Fees:      "char:fees:{characterId}"
```

## 🛡️ Security Considerations

### Token Storage
- **Access Tokens:** Kurze TTL (20min), nur für API-Calls
- **Refresh Tokens:** Sichere serverseitige Speicherung
- **JWT Tokens:** Nur Character-ID und Name, keine sensiblen Daten

### Rate Limiting
- **Character API Calls:** Separates Rate Limiting von Market Data
- **Token Refresh:** Exponential Backoff bei Fehlern
- **Cache-First:** Minimiere ESI-Calls durch intelligentes Caching

---

**Character Integration ermöglicht personalisierte Trading-Kalkulationen basierend auf Skills, Assets und aktueller Situation!**
