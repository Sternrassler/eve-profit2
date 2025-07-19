# Go Backend Spezifikationen - EVE Profit Calculator 2.0

## 🚀 Go Backend Architektur

### Core Design Prinzipien
- **Massive Parallelität:** Goroutines für alle ESI-Abfragen
- **Intelligentes Caching:** Multi-Layer Cache-Strategie
- **Resilience:** Circuit Breaker Pattern für ESI-Ausfälle
- **Performance:** Sub-100ms Response Times für gecachte Daten

## 🛠️ Tech Stack Details

### Go Dependencies
```go
// Core Framework
github.com/gin-gonic/gin           // HTTP Router
github.com/allegro/bigcache/v3     // In-Memory Cache

// Database
github.com/mattn/go-sqlite3        // SQLite Driver (for SDE)
github.com/jmoiron/sqlx           // SQL Extensions

// EVE ESI & HTTP
golang.org/x/time/rate            // Rate Limiting
golang.org/x/sync/errgroup        // Concurrent Error Handling

// OAuth & JWT
golang.org/x/oauth2               // OAuth2 Client
github.com/golang-jwt/jwt/v5      // JWT Token Handling

// Configuration & Utilities
github.com/spf13/viper            // Configuration
github.com/rs/zerolog             // Structured Logging
```

## 🏗️ Detaillierte Backend-Struktur

```
backend/
├── cmd/server/
│   └── main.go                   # Application Entry Point
├── internal/
│   ├── api/
│   │   ├── handlers/             # HTTP Request Handlers
│   │   │   ├── market.go        # Market Data Endpoints
│   │   │   ├── items.go         # Item Search Endpoints (SDE)
│   │   │   ├── profit.go        # Profit Calculation Endpoints
│   │   │   ├── character.go     # Character Data Endpoints
│   │   │   └── auth.go          # OAuth Authentication Endpoints
│   │   ├── middleware/          # HTTP Middleware
│   │   │   ├── cors.go          # CORS Handler
│   │   │   ├── ratelimit.go     # Rate Limiting
│   │   │   └── logger.go        # Request Logging
│   │   └── router.go            # Route Configuration
│   ├── service/
│   │   ├── market.go            # Market Data Business Logic
│   │   ├── profit.go            # Profit Calculation Logic
│   │   ├── items.go             # Item Search Logic (SDE)
│   │   ├── character.go         # Character Data Logic
│   │   ├── auth.go              # OAuth Authentication Logic
│   │   └── cache.go             # Cache Management
│   ├── repository/
│   │   ├── sde/                 # SDE SQLite Repository
│   │   │   ├── items.go         # Item Data Access
│   │   │   ├── stations.go      # Station Data Access
│   │   │   └── regions.go       # Region Data Access
│   │   └── interfaces.go        # Repository Interfaces
│   ├── cache/
│   │   ├── bigcache.go          # BigCache Implementation
│   │   └── interfaces.go        # Cache Interfaces
│   └── models/
│       ├── market.go            # Market Data Structures
│       ├── item.go              # Item Data Structures (SDE)
│       ├── character.go         # Character Data Structures
│       ├── auth.go              # OAuth Token Structures
│       └── response.go          # API Response Structures
├── pkg/
│   ├── eve/
│   │   ├── client.go            # ESI HTTP Client
│   │   ├── oauth.go             # EVE SSO OAuth Client
│   │   ├── types.go             # ESI Response Types
│   │   ├── endpoints.go         # ESI Endpoint Definitions
│   │   └── ratelimiter.go       # ESI Rate Limiting
│   ├── sde/
│   │   ├── client.go            # SDE SQLite Client
│   │   ├── queries.go           # Pre-compiled SQL Queries
│   │   └── types.go             # SDE Data Types
│   ├── config/
│   │   └── config.go            # Configuration Management
│   └── utils/
│       ├── calculations.go      # Trading Calculations
│       └── parallel.go          # Parallel Processing Utilities
├── data/
│   ├── sde.sqlite               # EVE Static Data Export Database (Fuzzwork)
│   └── sde_update.sh            # SDE Update Script (Fuzzwork Download)
├── scripts/
│   ├── build.sh                 # Build Scripts
│   └── download_sde.sh          # SDE Download Script (Fuzzwork)
├── docker/
│   ├── Dockerfile               # Production Docker Image
│   └── docker-compose.yml       # Development Environment
├── go.mod                       # Go Module Definition
├── go.sum                       # Dependency Checksums
├── .env.example                 # Environment Template
└── README.md                    # Backend Documentation
```

## ⚡ Parallelisierungs-Strategien

### 1. Worker Pool Pattern
```go
type WorkerPool struct {
    workers    int
    jobQueue   chan Job
    resultChan chan Result
    ctx        context.Context
}

// Für massive parallele ESI-Abfragen
func (wp *WorkerPool) ProcessMarketData(regions []int32, typeIDs []int32)
```

### 2. Fan-Out/Fan-In Pattern
```go
// Gleichzeitige Abfrage mehrerer Regionen
func GetMultiRegionMarketData(ctx context.Context, regions []int32, typeID int32) 
```

### 3. Pipeline Pattern
```go
// ESI Fetch -> Cache -> Transform -> Response
func MarketDataPipeline(ctx context.Context, request MarketRequest) 
```

## 🗃️ Caching-Architektur

### Simplified In-Memory Only Strategy
```go
type CacheManager struct {
    marketCache  *bigcache.BigCache    // ESI Market Data (5min TTL)
    historyCache *bigcache.BigCache    // ESI Price History (1h TTL)
    itemCache    map[int32]*Item       // SDE Items (permanent in memory)
    stationCache map[int32]*Station    // SDE Stations (permanent)
    regionCache  map[int32]*Region     // SDE Regions (permanent)
}

// Cache-Hit Prioritäten: Memory -> SDE SQLite -> ESI
```

### Cache Keys Schema
```
Market Data:    "market:{region}:{type}:{order_type}"
Price History:  "history:{region}:{type}:{date}"
SDE Data:       Loaded at startup, no cache keys needed
```

### TTL-Strategien
```go
const (
    MarketDataTTL    = 5 * time.Minute   // Live Market Orders
    HistoryDataTTL   = 1 * time.Hour     // Price History
    // SDE Data: Permanent in memory (loaded at startup)
)
```

## 🔌 API Endpoints Design

### RESTful API Structure
```
# Public Endpoints (No Auth)
GET    /api/v1/market/orders/{region}/{type}     # Market Orders (ESI + Cache)
GET    /api/v1/market/history/{region}/{type}    # Price History (ESI + Cache)
GET    /api/v1/market/compare                    # Multi-Region Compare
POST   /api/v1/profit/calculate                 # Profit Calculation
GET    /api/v1/items/search?q={query}           # Item Search (SDE)
GET    /api/v1/items/{type}                     # Item Details (SDE)
GET    /api/v1/stations                         # Stations List (SDE)
GET    /api/v1/regions                          # Regions List (SDE)
GET    /api/v1/health                           # Health Check

# Authentication Endpoints
POST   /api/v1/auth/login                       # EVE SSO Login
POST   /api/v1/auth/callback                    # OAuth Callback
POST   /api/v1/auth/refresh                     # Token Refresh
DELETE /api/v1/auth/logout                      # Logout

# Character Endpoints (Auth Required)
GET    /api/v1/character/info                   # Character Info
GET    /api/v1/character/skills                 # Character Skills
GET    /api/v1/character/assets                 # Character Assets
GET    /api/v1/character/ships                  # Character Ships
GET    /api/v1/character/wallet                 # Wallet Balance
GET    /api/v1/character/orders                 # Active Market Orders
```

### Response Format Standard
```go
type APIResponse struct {
    Success   bool        `json:"success"`
    Data      interface{} `json:"data,omitempty"`
    Error     *APIError   `json:"error,omitempty"`
    Meta      *MetaData   `json:"meta,omitempty"`
    Timestamp time.Time   `json:"timestamp"`
}
```

## 🎯 Performance-Ziele

### Response Time Targets
- **Cache Hit:** < 50ms
- **Cache Miss (Single ESI Call):** < 500ms  
- **Batch Operations:** < 2s für 10 Regionen
- **Search Queries:** < 100ms

### Throughput Targets
- **Concurrent Users:** 1000+
- **Requests/Second:** 5000+
- **ESI Efficiency:** 95%+ Cache Hit Rate

## 🛡️ Error Handling & Resilience

### Circuit Breaker Pattern
```go
type CircuitBreaker struct {
    maxFailures int
    resetTime   time.Duration
    state       CircuitState
}

// Schützt vor ESI-Ausfällen
func (cb *CircuitBreaker) Call(fn func() error) error
```

### Exponential Backoff
```go
func RetryWithBackoff(ctx context.Context, fn func() error, maxRetries int) error {
    // 1s, 2s, 4s, 8s, 16s delays
}
```

### Graceful Degradation
- ESI-Ausfall: Serve cached data mit Stale-Markierung
- Database-Ausfall: In-Memory only mode
- Redis-Ausfall: Direct database fallback

## 🔧 Development Workflow

### Environment Setup
```bash
# Development
go run cmd/server/main.go

# With hot reload  
air

# Docker development
docker-compose up -d
```

### Testing Strategy
```go
// Unit Tests
func TestMarketDataService(t *testing.T)

// Integration Tests  
func TestESIClientIntegration(t *testing.T)

// Load Tests
func BenchmarkParallelESICalls(b *testing.B)
```

## 📊 Monitoring & Observability

### Metrics to Track
- ESI Request Latency & Success Rate
- Cache Hit Ratios (L1/L2/L3)
- Goroutine Count & Memory Usage
- API Response Times per Endpoint

### Logging Structure
```go
log.Info().
    Str("endpoint", "/api/v1/market/orders").
    Int("region", 10000002).
    Dur("duration", duration).
    Bool("cache_hit", true).
    Msg("Market data request processed")
```

## 🚀 Deployment Considerations

### Production Optimizations
- **Binary Size:** < 20MB statically linked
- **Memory Usage:** < 512MB base + cache
- **CPU:** Efficient auf 2+ Cores
- **Scaling:** Horizontal via Load Balancer

### Docker Configuration
```dockerfile
FROM golang:1.21-alpine AS builder
# Multi-stage build for minimal image size
FROM alpine:latest
# Final image ~15MB
```

## 🐳 Docker Integration mit SDE

### Dockerfile mit automatischem SDE Download
```dockerfile
# backend/Dockerfile
FROM golang:1.21-alpine AS builder

# Install dependencies for downloading and decompressing
RUN apk add --no-cache curl bzip2 sqlite

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o eve-profit-server cmd/server/main.go

# Runtime image
FROM alpine:latest

RUN apk add --no-cache ca-certificates curl bzip2 sqlite
RUN adduser -D -s /bin/sh appuser

WORKDIR /app

# Copy binary
COPY --from=builder /app/eve-profit-server .
COPY --from=builder /app/scripts/download_sde.sh ./scripts/

# Create data directory
RUN mkdir -p data && chown appuser:appuser data

# Download SDE at build time (optional)
# RUN ./scripts/download_sde.sh

USER appuser

EXPOSE 8080

# Download SDE at startup if not exists
CMD ["sh", "-c", "[ ! -f data/sde.sqlite ] && ./scripts/download_sde.sh || true; ./eve-profit-server"]
```

### Docker Compose mit SDE Volume
```yaml
# docker-compose.yml
version: '3.8'

services:
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    volumes:
      - sde_data:/app/data
    environment:
      - GIN_MODE=release
      - SDE_PATH=/app/data/sde.sqlite
      - LOG_LEVEL=info
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/api/v1/health"]
      interval: 30s
      timeout: 10s
      retries: 3

  frontend:
    build: ./frontend
    ports:
      - "3000:3000"
    environment:
      - VITE_BACKEND_URL=http://localhost:8080
    depends_on:
      - backend

volumes:
  sde_data:
    driver: local
```

### SDE Startup Check
```go
// cmd/server/main.go
func main() {
    // Check if SDE exists, download if not
    sdePath := config.GetString("sde.path", "./data/sde.sqlite")
    if _, err := os.Stat(sdePath); os.IsNotExist(err) {
        log.Info().Msg("SDE not found, downloading from Fuzzwork...")
        if err := sde.DownloadLatestSDE(); err != nil {
            log.Fatal().Err(err).Msg("Failed to download SDE")
        }
    }
    
    // Initialize SDE client
    sdeClient, err := sde.NewClient(sdePath)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to initialize SDE client")
    }
    
    // Validate SDE
    if err := sdeClient.ValidateSDE(); err != nil {
        log.Fatal().Err(err).Msg("SDE validation failed")
    }
    
    // Load static data into memory
    if err := sdeClient.LoadStaticDataIntoMemory(); err != nil {
        log.Fatal().Err(err).Msg("Failed to load static data")
    }
    
    log.Info().Msg("SDE initialized successfully")
    
    // Continue with server startup...
}
```
---

**Diese Spezifikationen definieren die Go-Backend Implementierung für optimale EVE ESI Performance!**
