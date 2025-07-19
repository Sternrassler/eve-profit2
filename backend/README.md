# EVE Profit Calculator 2.0 - Backend

Go Backend für den EVE Online Profit Calculator mit Gin Framework.

## 🚀 Quick Start

### Voraussetzungen
- Go 1.21+
- SQLite3 (für SDE-Verifikation)
- curl oder wget (für SDE-Download)

### Installation & Start

```bash
# 1. SDE Database herunterladen (einmalig)
./download-sde.sh

# 2. Dependencies installieren
go mod tidy

# 3. Server starten
go run cmd/server/main.go
```

Server läuft auf http://localhost:8080

## 🛠️ API Endpoints

### Health Check
```bash
curl http://localhost:8080/health
```

### API v1 (Stubs)
```bash
# Market Data
curl http://localhost:8080/api/v1/market/prices/34

# Items Search  
curl http://localhost:8080/api/v1/items/search?q=tritanium

# Profit Calculation
curl -X POST http://localhost:8080/api/v1/profit/calculate
```

## 📊 SDE Database

- **Source:** Fuzzwork SQLite Export
- **Size:** ~544MB (entpackt)
- **Items:** 25.818 veröffentlichte EVE Items
- **Stations:** 5.154 Stationen
- **Systems:** 8.437 Sonnensysteme

### SDE Update
```bash
./download-sde.sh --force  # Neuer Download erzwingen
```

## 🏗️ Architektur

```
cmd/server/          # Main Application
internal/
├── api/handlers/    # HTTP Request Handlers  
├── api/middleware/  # HTTP Middleware (CORS, Auth, etc.)
├── service/         # Business Logic Layer
├── repository/      # Data Access Layer (SDE SQLite)
├── cache/          # In-Memory Caching (BigCache)
└── models/         # Data Models & Types
```

## 📋 Status

✅ **Implementiert:**
- Gin Router + API Structure
- Health Check Endpoint  
- Middleware (CORS, Logging, Recovery)
- SDE SQLite Download + Verifikation
- Grundlegende API Endpoint Stubs

🚧 **In Entwicklung:**
- SDE Client Implementation  
- BigCache Integration
- ESI API Client
- Market Data Services

## 🎯 Nächste Schritte

1. **SDE Client:** SQLite Integration für Items, Stations, Regionen
2. **Cache Integration:** BigCache Setup für Performance
3. **ESI Client:** EVE API Integration mit Rate Limiting
4. **Business Logic:** Profit Calculation Engine

## 🔄 Git Workflow

Nach jeder abgeschlossenen Entwicklungsphase:
```bash
cd .. && ./commit-phase.sh 2 "SDE Client Implementation" "SQLite integration completed"
```

## 📊 Development Tracking

- **Repository:** https://github.com/Sternrassler/eve-profit2
- **Git-Historie:** Jeder Commit = Abgeschlossene Phase
- **Status:** Siehe [../STATUS.md](../STATUS.md) für aktuellen Stand

---

Siehe [../STATUS.md](../STATUS.md) für detaillierten Entwicklungsstand.
