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
