# Development Guidelines - EVE Profit Calculator 2.0

## 🎯 Für neue Sessions - Sofort lesen!

Diese Datei enthält wichtige Entwicklungsrichtlinien, die bei jeder neuen Copilot-Session beachtet werden müssen.

## 📖 Obligatorische Session-Vorbereitung
**Vor jeder Entwicklungsarbeit:**
1. `PROJECT_CONTEXT.md` lesen
2. `DEVELOPMENT_GUIDELINES.md` (diese Datei) lesen
3. Aktuellen Code-Stand überprüfen

## 🔧 Technische Festlegungen

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
- **Test-Driven Development (TDD):** Pflicht ab Phase 2
- **Test-First Approach:** Schreibe Tests vor der Implementierung
- **Coverage Minimum:** 80% Code Coverage für Business Logic
- **Test Kategorien:**
  - Unit Tests: Einzelne Funktionen/Methoden
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

## 📁 Datei-Organisation

### Backend (Go) - Zwingend einzuhalten:
```
backend/
├── cmd/server/          # Application Entry Point
├── internal/
│   ├── api/handlers/    # HTTP Request Handlers
│   ├── service/         # Business Logic Layer
│   ├── repository/      # Data Access Layer
│   ├── cache/          # Caching Implementation
│   └── models/         # Data Models & Types
├── pkg/
│   ├── eve/            # EVE ESI Client
│   ├── config/         # Configuration Management
│   └── utils/          # Shared Utilities
└── migrations/         # Database Migrations
```

### Frontend (React) - Zwingend einzuhalten:
```
frontend/src/
├── components/
│   ├── ui/              # Basis-UI Komponenten (Button, Input, etc.)
│   ├── features/        # Feature-spezifische Komponenten
│   └── layout/          # Layout-Komponenten
├── hooks/
│   ├── api/             # Backend API Hooks
│   └── utils/           # Utility Hooks
├── services/
│   └── backend-api/     # Backend API Service Layer
├── types/
│   ├── api/             # Backend API Response Types
│   └── app/             # App-spezifische Types
└── utils/
    ├── calculations/    # Frontend-Berechnungen
    └── formatters/      # Daten-Formatierung
```

## 🔄 API-Service Pattern

### Go Backend Service Structure
```go
// EVE ESI Client mit Parallelisierung
type ESIClient struct {
    httpClient   *http.Client
    rateLimiter  *rate.Limiter
    cache        CacheProvider
    workerPool   *WorkerPool
}

// Beispiel: Parallele Marktdaten-Abfrage
func (c *ESIClient) GetMarketDataParallel(
    ctx context.Context, 
    regions []int32, 
    typeID int32,
) ([]MarketData, error) {
    // Worker Pool + Goroutines implementation
}
```

### Frontend API Service Structure
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

## 🧪 Testing-Anforderungen

### TDD-Workflow (Test-Driven Development):
```
1. 🔴 RED: Schreibe einen fehlschlagenden Test
2. 🟢 GREEN: Implementiere minimalen Code für Test-Pass
3. 🔄 REFACTOR: Verbessere Code ohne Tests zu brechen
4. 📝 REPEAT: Für jede neue Funktion
```

### Test-Strategie nach Entwicklungsphase:
- **Phase 2:** SDE Client Tests (Items, Stations, Regionen)
- **Phase 3:** ESI API Client Tests (Rate Limiting, Parallel Calls)
- **Phase 4:** Character API Tests (OAuth, JWT, Permissions)
- **Phase 5:** Business Logic Tests (Profit Calculation, Trading Routes)
- **Phase 6:** Frontend Component Tests (React Testing Library)

### Go Backend Testing:
```go
// Unit Test Beispiel
func TestSDERepository_GetItemByID(t *testing.T) {
    // Given: Test Database Setup
    repo := setupTestSDERepo(t)
    
    // When: Execute Function
    item, err := repo.GetItemByID(34) // Tritanium
    
    // Then: Verify Results
    assert.NoError(t, err)
    assert.Equal(t, "Tritanium", item.TypeName)
    assert.Equal(t, int32(34), item.TypeID)
}
```

### Frontend Testing:
```typescript
// React Component Test Beispiel
describe('ItemSearchComponent', () => {
  it('should display search results when API returns data', async () => {
    // Given: Mock API Response
    mockApiService.searchItems.mockResolvedValue(mockItems);
    
    // When: User types search query
    render(<ItemSearchComponent />);
    fireEvent.change(screen.getByRole('textbox'), { 
      target: { value: 'Tritanium' } 
    });
    
    // Then: Results should be displayed
    await waitFor(() => {
      expect(screen.getByText('Tritanium')).toBeInTheDocument();
    });
  });
});
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
3. **Services:** API Mock Tests
4. **Utils:** Unit Tests für Berechnungen

### Test-Datei Konvention:
- `ComponentName.test.tsx` für Komponenten
- `hookName.test.ts` für Hooks
- `serviceName.test.ts` für Services

## 🚨 Code-Review Checklist

### Vor jedem Commit:
- [ ] TypeScript Errors behoben
- [ ] ESLint Warnings addressiert
- [ ] Komponenten-Props dokumentiert
- [ ] Performance optimiert (useMemo, useCallback wo nötig)
- [ ] Error Boundaries implementiert
- [ ] Loading States hinzugefügt

## 💡 EVE Online Spezifika

### Wichtige Konzepte:
- **ISK:** EVE-Währung (Inter Stellar Kredits)
- **Jita:** Haupt-Handelshub (System-ID: 30000142)
- **Market Orders:** Buy-Orders (bids) vs Sell-Orders (asks)
- **Jump Distance:** Wichtig für Transportkosten

### Business Logic:
```typescript
// Profit-Berechnung Template
const calculateProfit = (
  buyPrice: number,
  sellPrice: number,
  quantity: number,
  taxes: TaxSettings
) => {
  // Broker-Gebühren, Verkaufssteuern, Transportkosten
}
```

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
