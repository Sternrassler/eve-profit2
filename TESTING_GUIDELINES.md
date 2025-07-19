# Testing Guidelines - EVE Profit Calculator 2.0

## 🧪 Clean Code + Test-Driven Development (TDD)

**Dieses Projekt verwendet strikt TDD kombiniert mit Clean Code Prinzipien.**

### **TDD + Clean Code Philosophie:**
- **Test-First:** Alle neuen Features beginnen mit einem Test
- **Clean Tests:** Tests sind genauso wichtig wie Production Code
- **Living Documentation:** Tests dokumentieren das erwartete Verhalten
- **Refactoring Safety:** Tests ermöglichen sicheres Refactoring

## 🔄 TDD Red-Green-Refactor Cycle (Clean Code Edition)

### **1. 🔴 RED Phase - Test schreiben:**
```go
// ✅ Clean Test: Beschreibt Verhalten klar und präzise
func TestCalculateProfit_WithValidInputs_ShouldReturnCorrectProfit(t *testing.T) {
    // Arrange: Setup with meaningful test data
    buyPrice := 100.0
    sellPrice := 120.0
    quantity := 10
    
    // Act: Call the function under test
    actualProfit := CalculateProfit(buyPrice, sellPrice, quantity)
    
    // Assert: Verify expected behavior
    expectedProfit := 200.0 // (120-100) * 10
    assert.Equal(t, expectedProfit, actualProfit, 
        "Profit should be (sellPrice - buyPrice) * quantity")
}
```

**RED Phase Checklist:**
- [ ] Test-Name beschreibt das erwartete Verhalten
- [ ] Test ist selbstdokumentierend
- [ ] Arrange-Act-Assert Pattern verwendet
- [ ] Test kompiliert aber schlägt fehl

### **2. 🟢 GREEN Phase - Minimal-Implementation:**
```go
// ✅ Minimal aber funktional
func CalculateProfit(buyPrice, sellPrice float64, quantity int) float64 {
    return (sellPrice - buyPrice) * float64(quantity)
}
```

**GREEN Phase Checklist:**
- [ ] Gerade genug Code für grünen Test
- [ ] Keine Optimierung oder "schöne" Lösungen
- [ ] Test läuft erfolgreich durch

### **3. 🔄 REFACTOR Phase - Clean Code anwenden:**
```go
// ✅ Clean Code: Verbesserte Lesbarkeit und Robustheit
func CalculateProfit(buyPrice, sellPrice float64, quantity int) float64 {
    if quantity <= 0 {
        return 0.0
    }
    
    if buyPrice < 0 || sellPrice < 0 {
        return 0.0
    }
    
    profitPerUnit := sellPrice - buyPrice
    totalProfit := profitPerUnit * float64(quantity)
    
    return totalProfit
}
```

**REFACTOR Phase Checklist:**
- [ ] Clean Code Prinzipien angewendet
- [ ] Aussagekräftige Variablennamen
- [ ] Edge Cases behandelt
- [ ] Alle Tests bleiben grün

## 🎯 Clean Code Testing Standards

### **Test-Namen als Spezifikation:**
```go
// ✅ Test-Namen folgen: [Method]_[Scenario]_[ExpectedBehavior]
func TestCalculateProfit_WithZeroQuantity_ShouldReturnZero(t *testing.T)
func TestCalculateProfit_WithNegativePrices_ShouldReturnZero(t *testing.T) 
func TestCalculateProfit_WithLargePriceSpread_ShouldReturnHighProfit(t *testing.T)
```

### **Arrange-Act-Assert (AAA) Pattern:**
```go
func TestMarketService_GetBestDeals_ShouldReturnSortedByProfit(t *testing.T) {
    // Arrange: Setup test conditions
    mockESI := setupMockESIClient()
    mockCache := setupMockCache()
    service := NewMarketService(mockESI, mockCache)
    
    expectedDeals := []Deal{
        {Item: "Tritanium", Profit: 500},
        {Item: "Pyerite", Profit: 300},
        {Item: "Mexallon", Profit: 100},
    }
    mockESI.ExpectGetMarketData().Return(expectedDeals, nil)
    
    // Act: Execute the function under test
    actualDeals, err := service.GetBestDeals("Jita", 3)
    
    // Assert: Verify the expected behavior
    require.NoError(t, err)
    assert.Len(t, actualDeals, 3)
    assert.Equal(t, 500.0, actualDeals[0].Profit, "First deal should have highest profit")
    assert.Equal(t, 300.0, actualDeals[1].Profit, "Second deal should have medium profit")
    assert.Equal(t, 100.0, actualDeals[2].Profit, "Third deal should have lowest profit")
}
```

### **Test-Isolation und Helper Functions:**
```go
// ✅ Clean Test Helpers für bessere Lesbarkeit
func setupTestMarketService(t *testing.T) (*MarketService, *MockESI, *MockCache) {
    mockESI := &MockESI{}
    mockCache := &MockCache{}
    service := NewMarketService(mockESI, mockCache)
    return service, mockESI, mockCache
}

func createTestDeal(itemName string, profit float64) Deal {
    return Deal{
        Item:   itemName,
        Profit: profit,
        Volume: 1000, // Default test volume
    }
}

func assertDealsAreSortedByProfitDesc(t *testing.T, deals []Deal) {
    for i := 1; i < len(deals); i++ {
        assert.GreaterOrEqual(t, deals[i-1].Profit, deals[i].Profit,
            "Deals should be sorted by profit in descending order")
    }
}
```

## 🎯 Test-Strategien nach Entwicklungsphase

### **Phase 2: SDE Client (SQLite Integration) - TDD**
```go
// ✅ TDD Examples für SDE Client
func TestSDERepository_GetItemByID_WithValidID_ShouldReturnCorrectItem(t *testing.T)
func TestSDERepository_GetItemByID_WithInvalidID_ShouldReturnNotFoundError(t *testing.T)
func TestSDERepository_SearchItems_WithPartialName_ShouldReturnMatchingItems(t *testing.T)
func TestSDERepository_SearchItems_WithEmptyQuery_ShouldReturnEmptyResult(t *testing.T)
```

**TDD-Focus Phase 2:**
- Database Connection & Query Performance
- Data Integrity & Type Mapping
- Error Handling (Item not found, DB connection issues)
- Cache Integration Tests

### **Phase 3: ESI API Client (EVE External API) - TDD**
```go
// ✅ TDD Examples für ESI Client
func TestESIClient_GetMarketOrders_WithValidRegion_ShouldReturnOrdersData(t *testing.T)
func TestESIClient_GetMarketOrders_WithRateLimit_ShouldRespectLimits(t *testing.T)
func TestESIClient_ParallelRequests_WithMultipleRegions_ShouldHandleConcurrency(t *testing.T)
func TestESIClient_NetworkError_ShouldRetryWithExponentialBackoff(t *testing.T)
```
func TestESIClient_ErrorHandling(t *testing.T)
```

**Test-Focus:**
- HTTP Client Mocking (ESI API Responses)
- Rate Limiting Compliance (max requests per second)
- Parallel Request Management (Goroutines)
- Network Error Recovery & Retries

### **Phase 4: Character API (OAuth + JWT)**
```go
// Test Examples für Character API
func TestOAuth_LoginFlow(t *testing.T)
func TestJWT_TokenValidation(t *testing.T)
func TestCharacterService_GetAssets(t *testing.T)
func TestCharacterService_GetSkills(t *testing.T)
```

**Test-Focus:**
- OAuth Flow Simulation
- JWT Token Lifecycle (create, validate, refresh)
- Permission & Scope Validation
- Character Data Integration

### **Phase 5: Business Logic (Profit Calculation)**
```go
// Test Examples für Business Logic
func TestProfitCalculator_BasicProfit(t *testing.T)
func TestProfitCalculator_TradingFees(t *testing.T)
func TestProfitCalculator_SkillBonuses(t *testing.T)
func TestTradingRoutes_FindBestRoutes(t *testing.T)
```

**Test-Focus:**
- Mathematical Accuracy (Profit formulas)
- Edge Cases (Zero profit, negative margins)
- Performance (Large dataset calculations)
- Integration with Character Skills

## 🏗️ Test-Struktur & Organisation

### **Ordner-Struktur:**
```
backend/
├── internal/
│   ├── repository/
│   │   ├── sde.go
│   │   └── sde_test.go           # Unit Tests
│   ├── service/
│   │   ├── market.go
│   │   └── market_test.go        # Service Layer Tests
│   └── api/handlers/
│       ├── items.go
│       └── items_test.go         # HTTP Handler Tests
├── tests/
│   ├── integration/              # Cross-component Tests
│   │   ├── api_integration_test.go
│   │   └── database_integration_test.go
│   ├── testdata/                 # Test Fixtures & Mock Data
│   │   ├── sde_test.sqlite
│   │   └── mock_esi_responses.json
│   └── mocks/                    # Generated Mocks
│       ├── mock_sde_repo.go
│       └── mock_esi_client.go
└── cmd/
    └── test-runner/              # Custom Test Runner
```

## 🔧 Test-Tools & Dependencies

### **Go Testing Stack:**
```go
// Testing Framework
import "testing"

// Assertions & Test Utilities
import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/mock"
import "github.com/stretchr/testify/suite"

// HTTP Testing
import "net/http/httptest"

// Database Testing
import "github.com/DATA-DOG/go-sqlmock"

// Context & Timeout Testing
import "context"
import "time"
```

### **Mock Generation:**
```bash
# Install mockery for automatic mock generation
go install github.com/vektra/mockery/v2@latest

# Generate mocks for interfaces
mockery --dir=internal/repository --name=SDERepository --output=tests/mocks
mockery --dir=internal/service --name=ESIClient --output=tests/mocks
```

## 📊 Test-Coverage & Quality

### **Coverage-Ziele:**
- **Unit Tests:** 90%+ für Business Logic
- **Integration Tests:** 80%+ für API Endpoints
- **End-to-End Tests:** Critical User Flows (100%)

### **Coverage-Messung:**
```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...

# View coverage report
go tool cover -html=coverage.out

# Coverage summary
go tool cover -func=coverage.out
```

### **Quality-Checks:**
```bash
# Run all tests
go test ./...

# Race condition detection
go test -race ./...

# Benchmarks for performance-critical code
go test -bench=. ./internal/service/

# Memory leak detection
go test -memprofile=mem.prof ./...
```

## 🎭 Test-Patterns & Best Practices

### **1. Arrange-Act-Assert Pattern:**
```go
func TestMarketService_CalculateProfit(t *testing.T) {
    // Arrange
    service := NewMarketService(mockCache)
    buyPrice := 100.0
    sellPrice := 150.0
    quantity := 1000
    
    // Act
    profit, err := service.CalculateProfit(buyPrice, sellPrice, quantity)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, 50000.0, profit)
}
```

### **2. Table-Driven Tests:**
```go
func TestProfitCalculation_VariousScenarios(t *testing.T) {
    tests := []struct {
        name       string
        buyPrice   float64
        sellPrice  float64
        quantity   int
        expected   float64
        expectErr  bool
    }{
        {"basic profit", 100, 150, 1000, 50000, false},
        {"zero profit", 100, 100, 1000, 0, false},
        {"loss scenario", 150, 100, 1000, -50000, false},
        {"invalid quantity", 100, 150, -1, 0, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            profit, err := CalculateProfit(tt.buyPrice, tt.sellPrice, tt.quantity)
            
            if tt.expectErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.expected, profit)
            }
        })
    }
}
```

### **3. Mock-basierte Tests:**
```go
func TestMarketService_WithMockedESI(t *testing.T) {
    // Arrange
    mockESI := &mocks.ESIClient{}
    mockESI.On("GetMarketOrders", mock.Anything, 10000002, 34).
           Return([]MarketOrder{testOrder}, nil)
    
    service := NewMarketService(mockESI)
    
    // Act
    orders, err := service.GetOrders(10000002, 34)
    
    // Assert
    assert.NoError(t, err)
    assert.Len(t, orders, 1)
    mockESI.AssertExpectations(t)
}
```

## 🚀 Test-Ausführung & CI/CD

### **Lokale Test-Ausführung:**
```bash
# Alle Tests
make test

# Spezifische Package Tests
go test ./internal/repository/

# Verbose Output
go test -v ./...

# Nur spezifische Tests
go test -run TestSDERepository ./internal/repository/
```

### **Test-Automation Script:**
```bash
#!/bin/bash
# test-all.sh

echo "🧪 Running EVE Profit Calculator Tests..."

# Unit Tests
echo "📋 Unit Tests..."
go test ./internal/... -cover

# Integration Tests  
echo "🔗 Integration Tests..."
go test ./tests/integration/...

# Race Condition Tests
echo "🏃 Race Condition Tests..."
go test -race ./...

# Performance Benchmarks
echo "⚡ Performance Benchmarks..."
go test -bench=. ./internal/service/

echo "✅ All tests completed!"
```

## 📋 Test-Checkliste für neue Features

### **Vor Implementation:**
- [ ] Test-Cases definiert (Happy Path + Edge Cases)
- [ ] Mock-Dependencies identifiziert
- [ ] Test-Data/Fixtures vorbereitet

### **Während TDD:**
- [ ] Red: Test schreibt und schlägt fehl
- [ ] Green: Minimale Implementation funktioniert
- [ ] Refactor: Code optimiert, Tests bestehen weiterhin

### **Nach Implementation:**
- [ ] Coverage-Report geprüft (>80%)
- [ ] Integration Tests hinzugefügt
- [ ] Performance-Tests bei kritischen Pfaden
- [ ] Documentation für komplexe Test-Scenarios

---

## 🎯 **TDD-Implementierung ab Phase 2**

**Test-First Development ist ab Phase 2 Pflicht für alle neuen Backend-Features. Jede Funktion startet mit einem Test, dann folgt die minimale Implementierung.**

**Ziel:** Höhere Code-Qualität, bessere Dokumentation durch Tests, und sichere Refactoring-Möglichkeiten für langfristige Projektentwicklung.
