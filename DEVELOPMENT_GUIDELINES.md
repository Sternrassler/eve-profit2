# Development Guidelines - EVE Profit Calculator 2.0

## 🎯 Für neue Sessions - Sofort lesen!

Diese Datei enthält wichtige Entwicklungsrichtlinien, die bei jeder neuen Copilot-Session beachtet werden müssen.

## 📖 Obligatorische Session-Vorbereitung
**Vor jeder Entwicklungsarbeit:**
1. `PROJECT_CONTEXT.md` lesen
2. `DEVELOPMENT_GUIDELINES.md` (diese Datei) lesen
3. Aktuellen Code-Stand überprüfen

## 🔧 Technische Festlegungen

## 📚 Clean Code + TDD Methodologie

**Dieses Projekt folgt konsequent Clean Code Prinzipien und Test-Driven Development.**

### Clean Code Prinzipien (Uncle Bob Martin)

#### 1. Lesbarkeit & Verständlichkeit
```go
// ❌ Schlecht: Cryptisch und unklar
func proc(d []int) int {
    s := 0
    for _, v := range d {
        s += v
    }
    return s
}

// ✅ Gut: Selbstdokumentierend
func CalculateTotalProfitFromTrades(tradeProfits []int) int {
    totalProfit := 0
    for _, individualTradeProfit := range tradeProfits {
        totalProfit += individualTradeProfit
    }
    return totalProfit
}
```

#### 2. Funktionen - Eine Verantwortlichkeit (SRP)
```go
// ❌ Schlecht: Mehrere Verantwortlichkeiten
func ProcessMarketData(data []byte) (*MarketOrder, error) {
    // Validierung
    if len(data) == 0 { return nil, errors.New("no data") }
    
    // Parsing
    var order MarketOrder
    json.Unmarshal(data, &order)
    
    // Business Logic
    order.Profit = order.SellPrice - order.BuyPrice
    
    // Persistierung
    database.Save(order)
    
    // Logging
    log.Printf("Saved order %d", order.ID)
    
    return &order, nil
}

// ✅ Gut: Getrennte Verantwortlichkeiten
func ValidateMarketData(data []byte) error { /* nur validierung */ }
func ParseMarketOrder(data []byte) (*MarketOrder, error) { /* nur parsing */ }
func CalculateProfit(order *MarketOrder) { /* nur berechnung */ }
func SaveMarketOrder(order *MarketOrder) error { /* nur persistierung */ }
```

#### 3. Aussagekräftige Namen
```typescript
// ❌ Schlecht: Abkürzungen und unklare Namen
const calc = (p1: number, p2: number, q: number) => {
  const d = p2 - p1;
  const t = d * 0.1; // Was ist 0.1?
  return (d - t) * q;
}

// ✅ Gut: Selbsterklärende Namen
const calculateNetProfitAfterTaxes = (
  buyPrice: number,
  sellPrice: number,
  quantity: number
) => {
  const grossProfitPerUnit = sellPrice - buyPrice;
  const taxPerUnit = grossProfitPerUnit * BROKER_TAX_RATE;
  const netProfitPerUnit = grossProfitPerUnit - taxPerUnit;
  return netProfitPerUnit * quantity;
}
```

#### 4. SOLID Prinzipien

**Single Responsibility Principle (SRP):**
```go
// ✅ Jede Struktur hat eine klare Verantwortlichkeit
type ESIClient struct { /* nur ESI API calls */ }
type MarketCalculator struct { /* nur Markt-Berechnungen */ }
type CacheManager struct { /* nur Caching */ }
```

**Open/Closed Principle (OCP):**
```go
// ✅ Erweiterbar durch Interfaces
type PriceCalculator interface {
    CalculatePrice(item Item, market Market) float64
}

type BasicPriceCalculator struct {}
type AdvancedPriceCalculator struct {} // Neue Implementation ohne Änderung bestehender
```

**Dependency Inversion Principle (DIP):**
```go
// ✅ Abhängigkeiten über Interfaces
type MarketService struct {
    esiClient    ESIClientInterface    // Interface, nicht konkrete Implementierung
    cache        CacheInterface       // Interface, nicht konkrete Implementierung
    calculator   CalculatorInterface  // Interface, nicht konkrete Implementierung
}
```

### Test-Driven Development (TDD) - Red-Green-Refactor

#### Strikter TDD-Workflow:
```
🔴 RED (Test schreiben):
   1. Schreibe einen fehlschlagenden Test
   2. Test muss aus dem richtigen Grund fehlschlagen
   3. Kompiliere-Fehler zählen als "failing test"

🟢 GREEN (Minimal-Implementation):
   1. Schreibe gerade genug Code, damit der Test grün wird
   2. Keine Optimierung, keine "schönen" Lösungen
   3. Hauptziel: Test zum Laufen bringen

🔄 REFACTOR (Code verbessern):
   1. Clean Code Prinzipien anwenden
   2. Duplikation entfernen
   3. Code-Struktur verbessern
   4. Tests dürfen NICHT brechen!
```

#### TDD Beispiel-Zyklus:
```go
// 1. 🔴 RED: Test schreiben (fehlschlagend)
func TestCalculateProfit_ShouldReturnCorrectProfit(t *testing.T) {
    // Given
    buyPrice := 100.0
    sellPrice := 120.0
    quantity := 10
    
    // When
    profit := CalculateProfit(buyPrice, sellPrice, quantity)
    
    // Then
    expected := 200.0 // (120-100) * 10
    assert.Equal(t, expected, profit)
}

// 2. 🟢 GREEN: Minimal-Implementation
func CalculateProfit(buyPrice, sellPrice float64, quantity int) float64 {
    return (sellPrice - buyPrice) * float64(quantity) // Minimal, aber funktional
}

// 3. 🔄 REFACTOR: Code verbessern (Tests bleiben grün)
func CalculateProfit(buyPrice, sellPrice float64, quantity int) float64 {
    if quantity <= 0 {
        return 0
    }
    
    profitPerUnit := sellPrice - buyPrice
    return profitPerUnit * float64(quantity)
}
```

### TypeScript Konfiguration
- **Strict Mode:** Aktiviert
- **No Any:** Vermeide `any` - nutze spezifische Types
- **Import Organization:** Absolute Imports mit `@/` für src-Ordner

### React Patterns
- **Functional Components:** Nur funktionale Komponenten, keine Klassen
- **Hooks:** Custom Hooks für wiederverwendbare Logik
- **Props Interface:** Jede Komponente hat explizite Props-Interface

### EVE API Handling
- **Rate Limiting:** Implementiert mit exponential backoff
- **Caching:** 5 Minuten für Marktdaten, 24h für statische Daten
- **Error Handling:** Graceful degradation bei API-Fehlern

### Testing Standards
- **Test-Driven Development (TDD):** Pflicht für ALLE neuen Features
- **Test-First Approach:** Schreibe Tests VOR der Implementierung
- **Coverage Minimum:** 90% Code Coverage für Business Logic
- **Test Kategorien:**
  - Unit Tests: Einzelne Funktionen/Methoden (TDD)
  - Integration Tests: API Endpoints + Database
  - End-to-End Tests: Critical User Flows

## 🎨 UI-Standards

### Komponenten-Hierarchie
```
Layout
├── Header (Navigation)
├── Main Content
│   ├── Sidebar (optional)
│   └── Content Area
└── Footer (minimal)
```

### Styling-Konventionen
- **Tailwind Classes:** Utility-first Ansatz
- **Custom CSS:** Nur für komplexe Animationen
- **Responsive:** Mobile-first mit `sm:`, `md:`, `lg:` Breakpoints

### Farb-Definitionen
```css
/* Primärfarben */
--primary-blue: #1e40af
--primary-light: #3b82f6
--primary-dark: #1e3a8a

/* EVE-inspiriert */
--eve-orange: #ff6b35
--eve-gold: #fbbf24
--background-dark: #0f172a
--surface-dark: #1e293b
```

## � API-Service Pattern

### Go Clean Code Service Pattern
```go
// ✅ Service mit Dependency Injection
type MarketService struct {
    esiClient    ESIClientInterface    // Interface, nicht konkret
    cache        CacheInterface       // Interface, nicht konkret
    logger       LoggerInterface      // Interface, nicht konkret
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
```
```typescript
class BackendApiService {
  private baseURL = 'http://localhost:8080/api/v1';
  
  // Alle Backend-Calls verwenden dieses Pattern:
  async getData<T>(endpoint: string): Promise<T> {
    // Error handling, loading states, retries
  }
}
```

### Response Handling
- **Success:** Typisierte Responses (Go Structs -> TS Interfaces)
- **Error:** Einheitliche Error-Objekte mit Codes
- **Loading:** Loading-States in UI mit Skeleton Components

## 🧪 Testing-Anforderungen

## 🧪 TDD Testing-Strategie

### TDD-Workflow (Test-Driven Development):
```
1. 🔴 RED: Schreibe einen fehlschlagenden Test
2. 🟢 GREEN: Implementiere minimalen Code für Test-Pass
3. 🔄 REFACTOR: Verbessere Code ohne Tests zu brechen
4. 📝 REPEAT: Für jede neue Funktion
```

### Clean Code + TDD Vereinheitlichung:
- **Jeder neue Code:** TDD-Zyklus ZUERST
- **Refactoring-Phase:** Clean Code Prinzipien konsequent anwenden
- **Code-Review:** Sowohl Test-Qualität als auch Code-Qualität prüfen

### TDD Best Practices:

#### 1. Test-Namen als Spezifikation
```go
// ✅ Tests beschreiben das erwartete Verhalten
func TestCalculateProfit_WithPositivePriceSpread_ShouldReturnPositiveProfit(t *testing.T)
func TestCalculateProfit_WithNegativePriceSpread_ShouldReturnNegativeProfit(t *testing.T)
func TestCalculateProfit_WithZeroQuantity_ShouldReturnZero(t *testing.T)
func TestCalculateProfit_WithInvalidInput_ShouldReturnError(t *testing.T)
```

#### 2. Arrange-Act-Assert (AAA) Pattern
```go
func TestMarketService_GetBestTrade_ShouldReturnHighestProfitTrade(t *testing.T) {
    // Arrange (Given)
    mockESI := &MockESIClient{}
    mockCache := &MockCache{}
    service := NewMarketService(mockESI, mockCache)
    
    expectedTrades := []Trade{
        {Profit: 100}, {Profit: 500}, {Profit: 200},
    }
    mockESI.On("GetTrades").Return(expectedTrades, nil)
    
    // Act (When)
    bestTrade, err := service.GetBestTrade("Jita")
    
    // Assert (Then)
    assert.NoError(t, err)
    assert.Equal(t, 500.0, bestTrade.Profit)
    mockESI.AssertExpectations(t)
}
```

#### 3. Test-Kategorien mit TDD:

**Unit Tests (Micro-TDD):**
```go
// Eine Funktion, ein Test-Zyklus
func TestCalculateBrokerFee_WithStandardRate_ShouldApplyCorrectPercentage(t *testing.T) {
    // Arrange
    orderValue := 1000000.0 // 1M ISK
    brokerRate := 0.03      // 3%
    
    // Act
    fee := CalculateBrokerFee(orderValue, brokerRate)
    
    // Assert
    expected := 30000.0 // 3% of 1M
    assert.Equal(t, expected, fee)
}
```

**Integration Tests (Macro-TDD):**
```go
// Ganze API-Endpoint, ein Test-Zyklus
func TestMarketHandler_GetMarketData_ShouldReturnFormattedResponse(t *testing.T) {
    // Arrange
    testDB := setupTestDatabase(t)
    handler := NewMarketHandler(testDB)
    
    // Act
    response := handler.GetMarketData(testRequest)
    
    // Assert
    assert.Equal(t, http.StatusOK, response.Status)
    assert.NotEmpty(t, response.Data.MarketOrders)
}
```

### Test-Strategie nach Entwicklungsphase:
- **Phase 2:** SDE Client Tests (Items, Stations, Regionen) - TDD für alle neuen Repository-Methoden
- **Phase 3:** ESI API Client Tests (Rate Limiting, Parallel Calls) - TDD für HTTP-Client-Layer
- **Phase 4:** Character API Tests (OAuth, JWT, Permissions) - TDD für Authentication-Services
- **Phase 5:** Business Logic Tests (Profit Calculation, Trading Routes) - TDD für alle Berechnungs-Algorithmen
- **Phase 6:** Frontend Component Tests (React Testing Library) - TDD für UI-Komponenten

### Go Backend Testing mit Clean Code:
```go
// ✅ Clean Code + TDD Beispiel
func TestSDERepository_GetItemByID_WithValidID_ShouldReturnCorrectItem(t *testing.T) {
    // Arrange: Setup Test Database
    testRepo := setupTestSDERepository(t)
    expectedItem := createTestItem("Tritanium", 34)
    
    // Act: Execute Function Under Test
    actualItem, err := testRepo.GetItemByID(34)
    
    // Assert: Verify Behavior
    assert.NoError(t, err, "GetItemByID should not return error for valid ID")
    assert.Equal(t, expectedItem.TypeName, actualItem.TypeName)
    assert.Equal(t, expectedItem.TypeID, actualItem.TypeID)
}

// Helper functions für bessere Lesbarkeit
func setupTestSDERepository(t *testing.T) *SDERepository {
    db := setupTestDatabase(t)
    return NewSDERepository(db)
}

func createTestItem(name string, id int32) *Item {
    return &Item{
        TypeName: name,
        TypeID:   id,
        GroupID:  1, // Default test value
    }
}
```

### Frontend Testing mit Clean Code:
```typescript
// ✅ Clean Code + TDD für React Components
describe('ItemSearchComponent', () => {
  const renderItemSearchWithMocks = (mockItems: Item[] = []) => {
    const mockApiService = createMockApiService();
    mockApiService.searchItems.mockResolvedValue(mockItems);
    
    return {
      ...render(<ItemSearchComponent apiService={mockApiService} />),
      mockApiService
    };
  };
  
  it('should display search results when user types valid item name', async () => {
    // Arrange
    const expectedItems = [createTestItem('Tritanium', 34)];
    const { mockApiService } = renderItemSearchWithMocks(expectedItems);
    
    // Act
    const searchInput = screen.getByRole('textbox', { name: /search items/i });
    fireEvent.change(searchInput, { target: { value: 'Tritanium' } });
    
    // Assert
    await waitFor(() => {
      expect(screen.getByText('Tritanium')).toBeInTheDocument();
    });
    expect(mockApiService.searchItems).toHaveBeenCalledWith('Tritanium');
  });
});

// Helper functions für bessere Teststruktur
function createMockApiService(): jest.Mocked<BackendApiService> {
  return {
    searchItems: jest.fn(),
    getMarketData: jest.fn(),
    // ... andere Methoden
  } as jest.Mocked<BackendApiService>;
}

function createTestItem(name: string, id: number): Item {
  return {
    name,
    id,
    groupId: 1, // Default test value
  };
}
```

### Test-Ordnerstruktur:
```
backend/
├── internal/
│   ├── repository/
│   │   ├── sde.go
│   │   └── sde_test.go         # Unit Tests
│   ├── service/
│   │   ├── market.go
│   │   └── market_test.go      # Business Logic Tests
│   └── api/handlers/
│       ├── items.go
│       └── items_test.go       # HTTP Handler Tests
└── tests/
    ├── integration/            # Integration Tests
    └── testdata/              # Test Fixtures

frontend/
├── src/
│   ├── components/
│   │   ├── ItemSearch.tsx
│   │   └── ItemSearch.test.tsx
│   └── services/
│       ├── api.ts
│       └── api.test.ts
└── __tests__/
    └── e2e/                   # End-to-End Tests
```

### Zu testende Bereiche:
1. **Komponenten:** Visual Regression Tests
2. **Hooks:** Logic Tests
3. **Services:** API Mock Tests mit Clean Interfaces
4. **Utils:** Unit Tests für Business Logic Calculations

### Test-Patterns (Clean Code):
- **Test-Namen:** Behavior-driven (should_return_when_given)
- **AAA-Pattern:** Arrange-Act-Assert konsistent verwenden
- **Helper Functions:** DRY-Prinzip auch in Tests
- **Mock-Interfaces:** Dependency Injection auch in Tests

## 🚨 Code-Review Checklist

### Vor jedem Commit:

#### Clean Code Prinzipien:
- [ ] **Namensgebung:** Aussagekräftige Namen für Variablen, Funktionen, Klassen
- [ ] **Funktionen:** Max. 20 Zeilen, eine Verantwortlichkeit (SRP)
- [ ] **Kommentare:** Code ist selbstdokumentierend, Kommentare nur für "Warum", nicht "Was"
- [ ] **DRY-Prinzip:** Keine Code-Duplikation
- [ ] **SOLID-Prinzipien:** Besonders SRP und DIP beachtet

#### TDD-Qualität:
- [ ] **Test-First:** Alle neuen Features mit TDD entwickelt
- [ ] **Test-Coverage:** Mindestens 90% Coverage für Business Logic
- [ ] **Test-Namen:** Beschreiben das erwartete Verhalten
- [ ] **AAA-Pattern:** Arrange-Act-Assert in allen Tests
- [ ] **Test-Isolation:** Tests laufen unabhängig voneinander

#### Technische Qualität:
- [ ] TypeScript Errors behoben
- [ ] ESLint Warnings addressiert
- [ ] Komponenten-Props dokumentiert
- [ ] Performance optimiert (useMemo, useCallback wo nötig)
- [ ] Error Boundaries implementiert
- [ ] Loading States hinzugefügt

#### Code-Struktur:
- [ ] **Dependency Injection:** Services über Interfaces
- [ ] **Error Handling:** Konsistente Fehlerbehandlung
- [ ] **Logging:** Strukturiertes Logging mit angemessenen Levels
- [ ] **Configuration:** Keine Hard-coded Values

### Clean Code Metriken:
- **Cyclomatic Complexity:** < 10 pro Funktion
- **Method Length:** < 20 Zeilen
- **Class Size:** < 300 Zeilen
- **Parameter Count:** < 5 Parameter pro Funktion

## 🎯 Performance-Ziele

### Metriken:
- **First Contentful Paint:** < 1.5s
- **Largest Contentful Paint:** < 2.5s
- **Bundle Size:** < 500KB gzipped
- **API Response Time:** Cache-miss < 3s

### Optimierungen:
- React.lazy() für Code-Splitting
- Virtual Scrolling für große Listen
- Debounced Search Inputs
- Optimistic UI Updates

## 📱 Responsive Breakpoints

```css
/* Mobile First */
sm: 640px   /* Small tablets */
md: 768px   /* Tablets */
lg: 1024px  /* Small desktops */
xl: 1280px  /* Large desktops */
2xl: 1536px /* Very large screens */
```

## 🔐 Security Guidelines

### API-Keys:
- Keine API-Keys im Client-Code
- Environment Variables für Konfiguration
- CORS-Policy beachten

### Data Validation:
- Input-Sanitization für User-Eingaben
- Zod für API-Response Validation
- XSS-Protection

---

**⚠️ WICHTIG: Diese Guidelines sind bindend und müssen bei jeder Entwicklungsarbeit befolgt werden!**
