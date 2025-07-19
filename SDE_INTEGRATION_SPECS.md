# SDE Integration Spezifikationen - EVE Profit Calculator 2.0

## 📊 EVE Static Data Export (SDE) Integration

### Was ist die SDE?
Die **EVE Static Data Export** ist eine SQLite-Datenbank mit allen statischen EVE Online Daten:
- **Items/Types:** Alle Schiffe, Module, Rohstoffe, etc.
- **Stations:** Alle Handelsstationen mit Standorten
- **Regions:** Alle EVE-Regionen mit IDs
- **Market Groups:** Kategorisierung der Items
- **Dogma:** Item-Attribute (Größe, Volumen, etc.)

## 🗄️ SDE SQLite Schema (Wichtigste Tabellen)

### invTypes (Items/Ships/Modules)
```sql
CREATE TABLE invTypes (
    typeID INTEGER PRIMARY KEY,
    groupID INTEGER,
    typeName VARCHAR(100),
    description TEXT,
    mass FLOAT,
    volume FLOAT,
    capacity FLOAT,
    published BOOLEAN
);
```

### invGroups (Item-Kategorien)
```sql
CREATE TABLE invGroups (
    groupID INTEGER PRIMARY KEY,
    categoryID INTEGER,
    groupName VARCHAR(100),
    published BOOLEAN
);
```

### staStations (Handelsstationen)
```sql
CREATE TABLE staStations (
    stationID INTEGER PRIMARY KEY,
    stationName VARCHAR(100),
    stationTypeID INTEGER,
    solarSystemID INTEGER,
    corporationID INTEGER,
    regionID INTEGER
);
```

### mapRegions (EVE Regionen)
```sql
CREATE TABLE mapRegions (
    regionID INTEGER PRIMARY KEY,
    regionName VARCHAR(100)
);
```

### mapSolarSystems (Sonnensysteme)
```sql
CREATE TABLE mapSolarSystems (
    solarSystemID INTEGER PRIMARY KEY,
    regionID INTEGER,
    solarSystemName VARCHAR(100),
    security FLOAT
);
```

## 🔧 Go SDE Client Implementation

### SDE Client Struktur
```go
// pkg/sde/client.go
type SDEClient struct {
    db     *sqlx.DB
    cache  map[string]interface{} // Pre-loaded data
    logger zerolog.Logger
}

type Item struct {
    TypeID      int32   `db:"typeID" json:"typeId"`
    GroupID     int32   `db:"groupID" json:"groupId"`
    TypeName    string  `db:"typeName" json:"typeName"`
    Description string  `db:"description" json:"description"`
    Mass        float64 `db:"mass" json:"mass"`
    Volume      float64 `db:"volume" json:"volume"`
    Published   bool    `db:"published" json:"published"`
}

type Station struct {
    StationID     int32  `db:"stationID" json:"stationId"`
    StationName   string `db:"stationName" json:"stationName"`
    SolarSystemID int32  `db:"solarSystemID" json:"solarSystemId"`
    RegionID      int32  `db:"regionID" json:"regionId"`
}

type Region struct {
    RegionID   int32  `db:"regionID" json:"regionId"`
    RegionName string `db:"regionName" json:"regionName"`
}
```

### Pre-compiled Queries
```go
// pkg/sde/queries.go
const (
    // Item Search Queries
    SearchItemsByName = `
        SELECT typeID, groupID, typeName, description, mass, volume, published
        FROM invTypes 
        WHERE typeName LIKE ? AND published = 1
        ORDER BY typeName
        LIMIT ?`
    
    GetItemByID = `
        SELECT typeID, groupID, typeName, description, mass, volume, published
        FROM invTypes 
        WHERE typeID = ?`
    
    GetItemsByGroup = `
        SELECT t.typeID, t.groupID, t.typeName, t.description, t.mass, t.volume, t.published
        FROM invTypes t
        JOIN invGroups g ON t.groupID = g.groupID
        WHERE g.groupName = ? AND t.published = 1
        ORDER BY t.typeName`
    
    // Station Queries
    GetStationsByRegion = `
        SELECT stationID, stationName, solarSystemID, regionID
        FROM staStations 
        WHERE regionID = ?
        ORDER BY stationName`
    
    GetStationByID = `
        SELECT stationID, stationName, solarSystemID, regionID
        FROM staStations 
        WHERE stationID = ?`
    
    // Region Queries
    GetAllRegions = `
        SELECT regionID, regionName
        FROM mapRegions
        ORDER BY regionName`
    
    GetRegionByID = `
        SELECT regionID, regionName
        FROM mapRegions 
        WHERE regionID = ?`
    
    // Trading Hub Queries (Major Trading Stations)
    GetTradingHubs = `
        SELECT s.stationID, s.stationName, s.solarSystemID, s.regionID, 
               ss.solarSystemName, r.regionName
        FROM staStations s
        JOIN mapSolarSystems ss ON s.solarSystemID = ss.solarSystemID
        JOIN mapRegions r ON s.regionID = r.regionID
        WHERE s.stationID IN (60003760, 60008494, 60004588, 60005686, 60011866)
        ORDER BY s.stationName`
)
```

### SDE Client Methods
```go
// Item Operations
func (c *SDEClient) SearchItems(query string, limit int) ([]Item, error)
func (c *SDEClient) GetItemByID(typeID int32) (*Item, error)
func (c *SDEClient) GetItemsByGroup(groupName string) ([]Item, error)

// Station Operations  
func (c *SDEClient) GetStationsByRegion(regionID int32) ([]Station, error)
func (c *SDEClient) GetStationByID(stationID int32) (*Station, error)
func (c *SDEClient) GetTradingHubs() ([]Station, error)

// Region Operations
func (c *SDEClient) GetAllRegions() ([]Region, error)
func (c *SDEClient) GetRegionByID(regionID int32) (*Region, error)

// Cache Operations (Startup)
func (c *SDEClient) LoadStaticDataIntoMemory() error
```

## 🚀 Startup Data Loading

### Memory Pre-loading Strategy
```go
// Beim Server-Start werden kritische Daten in den Speicher geladen
func (c *SDEClient) LoadStaticDataIntoMemory() error {
    // 1. Alle Regionen laden (klein, ~100 Einträge)
    regions, err := c.GetAllRegions()
    if err != nil {
        return err
    }
    c.cache["regions"] = regions
    
    // 2. Trading Hubs laden (sehr klein, ~5 Einträge)
    hubs, err := c.GetTradingHubs()
    if err != nil {
        return err
    }
    c.cache["trading_hubs"] = hubs
    
    // 3. Beliebte Item-Gruppen pre-loaden
    popularGroups := []string{"Battleship", "Cruiser", "Frigate", "PLEX"}
    for _, group := range popularGroups {
        items, err := c.GetItemsByGroup(group)
        if err != nil {
            continue // Non-fatal
        }
        c.cache["group_"+group] = items
    }
    
    return nil
}
```

## 📋 SDE Update Process

### Automatische SDE Updates (Fuzzwork SQLite)
```bash
#!/bin/bash
# scripts/download_sde.sh

# Fuzzwork bietet bereits konvertierte SQLite SDE
SDE_URL="https://www.fuzzwork.co.uk/dump/sqlite-latest.sqlite.bz2"
DATA_DIR="./data"

echo "Downloading latest SDE from Fuzzwork..."
curl -L -o /tmp/sde.sqlite.bz2 $SDE_URL

echo "Extracting SQLite database..."
bunzip2 -c /tmp/sde.sqlite.bz2 > $DATA_DIR/sde.sqlite

echo "Validating SDE database..."
sqlite3 $DATA_DIR/sde.sqlite "SELECT COUNT(*) FROM invTypes;"

echo "SDE Update completed!"
echo "Database size: $(du -h $DATA_DIR/sde.sqlite | cut -f1)"
```

### Fuzzwork SDE Vorteile
- **✅ Bereits konvertiert:** Keine YAML/CSV → SQLite Konvertierung nötig
- **✅ Regelmäßige Updates:** Automatisch nach EVE-Patches aktualisiert
- **✅ Optimiert:** Vordefinierte Indizes für bessere Performance
- **✅ Kompakt:** bz2-komprimiert, ~50MB Download
- **✅ Verlässlich:** Weit verbreitet in der EVE-Community

### SDE Download Integration
```go
// pkg/sde/updater.go
const (
    FuzzworkSQLiteURL = "https://www.fuzzwork.co.uk/dump/sqlite-latest.sqlite.bz2"
    LocalSDEPath      = "./data/sde.sqlite"
    TempDownloadPath  = "/tmp/sde.sqlite.bz2"
)

func DownloadLatestSDE() error {
    log.Info().Msg("Downloading latest SDE from Fuzzwork...")
    
    // Download bz2 file
    resp, err := http.Get(FuzzworkSQLiteURL)
    if err != nil {
        return fmt.Errorf("failed to download SDE: %w", err)
    }
    defer resp.Body.Close()
    
    // Save to temp file
    tempFile, err := os.Create(TempDownloadPath)
    if err != nil {
        return err
    }
    defer tempFile.Close()
    
    _, err = io.Copy(tempFile, resp.Body)
    if err != nil {
        return err
    }
    
    // Decompress bz2
    return decompressBZ2(TempDownloadPath, LocalSEDPath)
}
```

### SDE Validierung
```go
func (c *SDEClient) ValidateSDE() error {
    // Check essential tables exist and have data
    tables := map[string]int{
        "invTypes":        50000,   // ~50k items
        "invGroups":       1500,    // ~1.5k groups  
        "staStations":     4000,    // ~4k stations
        "mapRegions":      100,     // ~100 regions
        "mapSolarSystems": 8000,    // ~8k systems
    }
    
    for table, expectedMin := range tables {
        var count int
        err := c.db.Get(&count, fmt.Sprintf("SELECT COUNT(*) FROM %s", table))
        if err != nil {
            return fmt.Errorf("table %s missing or inaccessible: %w", table, err)
        }
        if count < expectedMin {
            return fmt.Errorf("table %s has only %d rows, expected at least %d", table, count, expectedMin)
        }
        c.logger.Info().Str("table", table).Int("count", count).Msg("SDE table validated")
    }
    
    // Test critical queries
    var testItem Item
    err := c.db.Get(&testItem, "SELECT typeID, typeName FROM invTypes WHERE typeID = 34 LIMIT 1") // Tritanium
    if err != nil {
        return fmt.Errorf("failed to query test item: %w", err)
    }
    
    return nil
}
```

## 🔍 Item Search Optimierung

### Search Performance
```go
// Fuzzy Search Implementation
func (c *SDEClient) FuzzySearchItems(query string, limit int) ([]Item, error) {
    // 1. Exact Match zuerst
    exactQuery := `typeName = ? AND published = 1`
    
    // 2. Prefix Match
    prefixQuery := `typeName LIKE ? AND published = 1`
    
    // 3. Contains Match
    containsQuery := `typeName LIKE ? AND published = 1`
    
    // Kombiniere Results mit Priorität
}

// Auto-complete für Item Search
func (c *SDEClient) GetItemSuggestions(query string) ([]string, error) {
    suggestions := []string{}
    rows, err := c.db.Query(`
        SELECT DISTINCT typeName 
        FROM invTypes 
        WHERE typeName LIKE ? AND published = 1
        ORDER BY typeName 
        LIMIT 10`, query+"%")
    // ...
    return suggestions, nil
}
```

## 🎯 Integration Points

### Service Layer Integration
```go
// internal/service/items.go
type ItemService struct {
    sdeClient *sde.SDEClient
    cache     *bigcache.BigCache
}

func (s *ItemService) SearchItems(query string, limit int) ([]models.Item, error) {
    // 1. Check cache first
    cacheKey := fmt.Sprintf("search:%s:%d", query, limit)
    if cached, err := s.cache.Get(cacheKey); err == nil {
        var items []models.Item
        json.Unmarshal(cached, &items)
        return items, nil
    }
    
    // 2. Query SDE
    sdeItems, err := s.sdeClient.SearchItems(query, limit)
    if err != nil {
        return nil, err
    }
    
    // 3. Convert & Cache
    items := convertSDEItems(sdeItems)
    if data, err := json.Marshal(items); err == nil {
        s.cache.Set(cacheKey, data)
    }
    
    return items, nil
}
```

### API Handler Integration
```go
// internal/api/handlers/items.go
func (h *ItemHandler) SearchItems(c *gin.Context) {
    query := c.Query("q")
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
    
    items, err := h.itemService.SearchItems(query, limit)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(200, gin.H{
        "success": true,
        "data": items,
        "meta": gin.H{
            "query": query,
            "count": len(items),
        },
    })
}
```

## 📊 Performance Considerations

### SDE Performance Metrics
- **Startup Load Time:** < 2s für kritische Daten
- **Item Search:** < 50ms für gecachte Queries
- **SDE Queries:** < 10ms für Index-optimierte Abfragen
- **Memory Usage:** ~50MB für pre-loaded Daten

### Indexing Strategy
```sql
-- Wichtige Indizes für Performance
CREATE INDEX idx_invTypes_typeName ON invTypes(typeName);
CREATE INDEX idx_invTypes_groupID ON invTypes(groupID);
CREATE INDEX idx_staStations_regionID ON staStations(regionID);
CREATE INDEX idx_mapSolarSystems_regionID ON mapSolarSystems(regionID);
```

---

**Die SDE Integration bietet blitzschnelle statische Datenabfragen ohne externe API-Abhängigkeiten!**
