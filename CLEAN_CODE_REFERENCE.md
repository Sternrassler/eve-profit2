# Clean Code Reference - EVE Profit Calculator 2.0

## 📚 Clean Code Prinzipien (Robert C. Martin)

Dieses Projekt folgt strikt den Clean Code Prinzipien. Diese Datei dient als schnelle Referenz.

## 🎯 Kernprinzipien

### 1. Meaningful Names (Aussagekräftige Namen)

```go
// ❌ Schlecht: Abkürzungen und unklare Namen
func calc(p1, p2 float64, q int) float64 {
    return (p2 - p1) * float64(q)
}

// ✅ Gut: Selbsterklärende Namen
func CalculateNetProfitFromTrade(buyPrice, sellPrice float64, quantity int) float64 {
    profitPerUnit := sellPrice - buyPrice
    return profitPerUnit * float64(quantity)
}
```

**Regeln:**
- Verwende aussagekräftige und aussprechbare Namen
- Verwende suchbare Namen für wichtige Konzepte
- Vermeide mentale Zuordnung und Codierung
- Ein Konzept pro Wort

### 2. Functions (Funktionen)

```go
// ❌ Schlecht: Zu viele Verantwortlichkeiten
func ProcessMarketOrder(data []byte) error {
    // Validierung
    if len(data) == 0 { return errors.New("no data") }
    
    // Parsing
    var order MarketOrder
    json.Unmarshal(data, &order)
    
    // Business Logic
    order.Profit = order.SellPrice - order.BuyPrice
    
    // Persistierung
    database.Save(order)
    
    // Logging
    log.Printf("Processed order %d", order.ID)
    
    return nil
}

// ✅ Gut: Single Responsibility
func ValidateMarketOrderData(data []byte) error {
    if len(data) == 0 {
        return errors.New("market order data cannot be empty")
    }
    return nil
}

func ParseMarketOrder(data []byte) (*MarketOrder, error) {
    var order MarketOrder
    if err := json.Unmarshal(data, &order); err != nil {
        return nil, fmt.Errorf("failed to parse market order: %w", err)
    }
    return &order, nil
}

func CalculateOrderProfit(order *MarketOrder) {
    order.Profit = order.SellPrice - order.BuyPrice
}
```

**Regeln:**
- Funktionen sollten klein sein (< 20 Zeilen)
- Funktionen sollten eine Sache tun
- Ein Abstraktionslevel pro Funktion
- Verwende deskriptive Namen

### 3. Comments (Kommentare)

```go
// ❌ Schlecht: Kommentare erklären WAS
func CalculateTax(amount float64) float64 {
    // Multiply amount by 0.1 to get 10% tax
    return amount * 0.1
}

// ✅ Gut: Kommentare erklären WARUM (wenn überhaupt nötig)
func CalculateBrokerFee(orderValue float64) float64 {
    // EVE Online broker fee is fixed at 3% for NPC stations
    // This rate cannot be reduced and varies for player structures
    const NPC_BROKER_FEE_RATE = 0.03
    return orderValue * NPC_BROKER_FEE_RATE
}

// ✅ Besser: Selbstdokumentierender Code
const EVE_NPC_BROKER_FEE_RATE = 0.03

func CalculateBrokerFee(orderValue float64) float64 {
    return orderValue * EVE_NPC_BROKER_FEE_RATE
}
```

**Regeln:**
- Kommentare sind ein Versagen des Codes, sich selbst zu erklären
- Erkläre WARUM, nicht WAS
- Halte Kommentare aktuell oder lösche sie

### 4. Formatting (Formatierung)

```go
// ✅ Konsistente Formatierung
type MarketService struct {
    esiClient    ESIClientInterface
    cache        CacheInterface
    logger       LoggerInterface
}

func NewMarketService(
    esiClient ESIClientInterface,
    cache CacheInterface,
    logger LoggerInterface,
) *MarketService {
    return &MarketService{
        esiClient: esiClient,
        cache:     cache,
        logger:    logger,
    }
}

func (s *MarketService) GetBestTrades(
    regionID int32,
    itemTypes []int32,
    maxResults int,
) ([]Trade, error) {
    if len(itemTypes) == 0 {
        return nil, errors.New("item types cannot be empty")
    }
    
    trades := make([]Trade, 0, maxResults)
    
    for _, itemType := range itemTypes {
        marketData, err := s.esiClient.GetMarketData(regionID, itemType)
        if err != nil {
            s.logger.Warn("Failed to get market data", 
                "region", regionID, 
                "item", itemType, 
                "error", err)
            continue
        }
        
        trade := s.calculateBestTrade(marketData)
        trades = append(trades, trade)
    }
    
    return s.sortTradesByProfit(trades), nil
}
```

## 🏗️ SOLID Prinzipien

### S - Single Responsibility Principle (SRP)

```go
// ✅ Jede Struktur hat eine klare Verantwortlichkeit
type MarketDataFetcher struct {
    esiClient ESIClientInterface
}

type ProfitCalculator struct {
    taxSettings TaxSettings
}

type TradeOptimizer struct {
    calculator ProfitCalculator
    fetcher    MarketDataFetcher
}
```

### O - Open/Closed Principle (OCP)

```go
// ✅ Erweiterbar ohne Modifikation
type PriceCalculator interface {
    CalculatePrice(item Item, market Market) float64
}

type BasicPriceCalculator struct{}
func (c *BasicPriceCalculator) CalculatePrice(item Item, market Market) float64 {
    return market.LowestSellOrder
}

type AdvancedPriceCalculator struct{}
func (c *AdvancedPriceCalculator) CalculatePrice(item Item, market Market) float64 {
    // Komplexere Berechnung mit Volumen, History, etc.
    return calculateWeightedAveragePrice(market)
}
```

### L - Liskov Substitution Principle (LSP)

```go
// ✅ Subtypes müssen ersetzbar sein
type Storage interface {
    Save(data interface{}) error
    Load(id string) (interface{}, error)
}

type DatabaseStorage struct{}
func (d *DatabaseStorage) Save(data interface{}) error { /* DB implementation */ }
func (d *DatabaseStorage) Load(id string) (interface{}, error) { /* DB implementation */ }

type CacheStorage struct{}
func (c *CacheStorage) Save(data interface{}) error { /* Cache implementation */ }
func (c *CacheStorage) Load(id string) (interface{}, error) { /* Cache implementation */ }

// Beide sind vollständig austauschbar
func ProcessData(storage Storage, data interface{}) error {
    return storage.Save(data) // Funktioniert mit beiden Implementierungen
}
```

### I - Interface Segregation Principle (ISP)

```go
// ❌ Schlecht: Zu große Interface
type MarketService interface {
    GetMarketData(regionID int32) ([]MarketOrder, error)
    CalculateProfit(buyPrice, sellPrice float64) float64
    SendNotification(message string) error
    LogActivity(activity string) error
}

// ✅ Gut: Getrennte, spezifische Interfaces
type MarketDataProvider interface {
    GetMarketData(regionID int32) ([]MarketOrder, error)
}

type ProfitCalculator interface {
    CalculateProfit(buyPrice, sellPrice float64) float64
}

type NotificationSender interface {
    SendNotification(message string) error
}

type ActivityLogger interface {
    LogActivity(activity string) error
}
```

### D - Dependency Inversion Principle (DIP)

```go
// ✅ Abhängigkeiten über Interfaces
type MarketService struct {
    dataProvider   MarketDataProvider    // Interface, nicht konkrete Implementierung
    calculator     ProfitCalculator      // Interface, nicht konkrete Implementierung
    cache          CacheInterface        // Interface, nicht konkrete Implementierung
    logger         LoggerInterface       // Interface, nicht konkrete Implementierung
}

func NewMarketService(
    dataProvider MarketDataProvider,
    calculator ProfitCalculator,
    cache CacheInterface,
    logger LoggerInterface,
) *MarketService {
    return &MarketService{
        dataProvider: dataProvider,
        calculator:   calculator,
        cache:        cache,
        logger:       logger,
    }
}
```

## 📋 Clean Code Checkliste

### Funktionen:
- [ ] Funktion hat eine einzige Verantwortlichkeit
- [ ] Funktionsname beschreibt genau was sie tut
- [ ] Maximal 20 Zeilen
- [ ] Maximal 3-4 Parameter
- [ ] Keine Side Effects

### Variablen:
- [ ] Aussagekräftige Namen
- [ ] Keine Abkürzungen oder Codierungen
- [ ] Konstanten für Magic Numbers
- [ ] Scope so klein wie möglich

### Klassen/Structs:
- [ ] Single Responsibility
- [ ] Kleine, cohäsive Klassen
- [ ] Dependency Injection verwendet
- [ ] Interface-basierte Abhängigkeiten

### Allgemein:
- [ ] Keine Code-Duplikation (DRY)
- [ ] Konsistente Formatierung
- [ ] Selbstdokumentierender Code
- [ ] Alle Tests grün

## 🚀 EVE Profit Calculator Spezifische Clean Code Patterns

### Trading Domain Model:
```go
// ✅ Saubere Domain-Modelle
type Trade struct {
    BuyOrder  MarketOrder
    SellOrder MarketOrder
    Item      Item
    Quantity  int
}

func (t *Trade) CalculateProfit() float64 {
    return t.calculateGrossProfit() - t.calculateFees()
}

func (t *Trade) calculateGrossProfit() float64 {
    priceSpread := t.SellOrder.Price - t.BuyOrder.Price
    return priceSpread * float64(t.Quantity)
}

func (t *Trade) calculateFees() float64 {
    brokerFees := t.calculateBrokerFees()
    salesTax := t.calculateSalesTax()
    return brokerFees + salesTax
}
```

### Service Layer Pattern:
```go
// ✅ Clean Service Implementation
type TradingService struct {
    marketData   MarketDataProvider
    profitCalc   ProfitCalculator
    itemRepo     ItemRepository
    logger       Logger
}

func (s *TradingService) FindBestTrades(
    ctx context.Context,
    criteria TradeCriteria,
) ([]Trade, error) {
    items, err := s.itemRepo.FindByCategory(criteria.Category)
    if err != nil {
        return nil, fmt.Errorf("failed to find items: %w", err)
    }
    
    var bestTrades []Trade
    for _, item := range items {
        trade, err := s.findBestTradeForItem(ctx, item, criteria)
        if err != nil {
            s.logger.Warn("Failed to find trade for item", 
                "item", item.Name, 
                "error", err)
            continue
        }
        
        if trade.CalculateProfit() >= criteria.MinProfit {
            bestTrades = append(bestTrades, trade)
        }
    }
    
    return s.sortTradesByProfit(bestTrades), nil
}
```

---

**💡 Denke daran: Clean Code ist nicht nur über Syntax - es geht um Kommunikation und Verständlichkeit!**
