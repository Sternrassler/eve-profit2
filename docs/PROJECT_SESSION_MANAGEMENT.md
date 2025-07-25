# EVE Profit Calculator 2.0 - Session Management

> **Schnelle Wiederaufnahme:** PROJECT_STATUS.md lesen + entsprechenden Command verwenden

**Universelle Session Guidelines siehe:** `UNIVERSAL_SESSION_MANAGEMENT_GUIDELINES.md`

## � EVE-spezifische Wiederaufnahme Commands

### Phase 4 fortsetzen (API Handlers)
```
"Lies PROJECT_STATUS.md. Phase 3 ESI Integration ist abgeschlossen. 
Implementiere jetzt Phase 4 API Handlers mit ESI Service Integration."
```

### Frontend starten (Phase 5)
```
"Lies PROJECT_STATUS.md. Backend läuft auf Port 9000 mit ESI Integration. 
Erstelle React + TypeScript Frontend mit Backend API Integration."
```

### Character API implementieren
```
"Lies PROJECT_API_SPECS.md. Implementiere EVE SSO OAuth Flow 
und Character API Handlers mit TDD."
```

## 🛠️ EVE Development Commands

### Server starten & testen
```bash
cd backend
go run cmd/server/main.go     # Server auf Port 9000
go test ./...                 # Alle Tests (18/18 ✅)
curl http://localhost:9000/api/v1/esi/test  # ESI Test
```

### Git Status & Commit
```bash
git status                    # EVE Projekt Änderungen anzeigen
git add -A                    # Alle EVE Änderungen stagen
git commit -m "feat(eve): ..."  # EVE-spezifische Commit Message
git push origin main          # Push to EVE-profit2 remote
git log --oneline -5          # Letzte EVE Commits anzeigen
```

### Development Dependencies
```bash
# Go Version Check (EVE Backend Requirements)
go version                    # Benötigt: Go 1.21+ für EVE Backend

# SQLite3 Check (EVE SDE Database)
sqlite3 --version            # Für EVE SDE Database (529MB)

# cURL Check (EVE ESI API Testing)
curl --version               # Für EVE ESI + Market API Tests

# EVE Dependencies Install
cd backend && go mod download # EVE Go Dependencies
```

## 📋 EVE Session Checklist

### EVE Session Start:
- [ ] `PROJECT_STATUS.md` gelesen - aktueller EVE Entwicklungsstand
- [ ] EVE Backend läuft: `go run cmd/server/main.go` (Port 9000)
- [ ] EVE Tests grün: `go test ./...` (18/18 ✅)
- [ ] EVE ESI Connection: `curl http://localhost:9000/api/v1/esi/test`
- [ ] Git Status sauber für EVE Projekt

### During EVE Development:
- [ ] TDD Red-Green-Refactor für EVE API Handlers
- [ ] Clean Code Standards für EVE Business Logic
- [ ] ESI Rate Limiting beachten (150 req/sec)
- [ ] EVE-spezifische Tests für Market/Character APIs
- [ ] SDE Database Performance optimiert

### EVE Session End:
- [ ] Alle EVE-Änderungen committed
- [ ] `PROJECT_STATUS.md` aktualisiert (Phase-Progress)
- [ ] EVE Tests grün: `go test ./...`
- [ ] ESI Integration funktional
- [ ] Nächste EVE Phase dokumentiert

## � EVE Production Commands

## 🛠️ EVE Development Commands

### Server starten & ESI testen
```bash
cd backend
go run cmd/server/main.go     # Server auf Port 9000
go test ./...                 # Alle Tests (18/18 ✅)
go test -v ./...              # Verbose Test Output
go test -cover ./...          # Test Coverage Report
curl http://localhost:9000/api/v1/esi/test  # ESI Connectivity Test
curl http://localhost:9000/api/v1/health   # Health Check
```

### EVE Konfiguration & Setup
```bash
cd backend
cp .env.example .env          # EVE ESI Settings vorkonfiguriert
./download-sde.sh            # SDE Database Download (529MB)
go mod tidy                   # EVE Dependencies aufräumen
```

### EVE Code Quality & Formatting
```bash
cd backend
go fmt ./...                  # Go Code formatieren
go vet ./...                  # Static Analysis für EVE Backend
go mod tidy                   # Go Modules cleanup
go build ./cmd/server         # EVE Server Binary erstellen
```

**💡 API Testing Commands siehe: `PROJECT_API_SPECS.md`**


## � Development Commands

### Server starten & testen
```bash
cd backend
go run cmd/server/main.go     # Server auf Port 9000
go test ./...                 # Alle Tests (18/18 ✅)
curl http://localhost:9000/api/v1/esi/test  # ESI Test
```

### Git Status & Commit
```bash
git status                    # Änderungen anzeigen
git add -A                    # Alle Änderungen stagen
git commit -m "..."          # Commit erstellen
```

## � EVE Production Commands

### EVE Backend Build & Deploy
```bash
cd backend
go build -o eve-server ./cmd/server    # EVE Server Binary
./eve-server                          # Production EVE Server
```

### EVE Performance & Monitoring
```bash
# EVE Server Memory Usage
ps aux | grep eve-server

# EVE SDE Database Size
du -h backend/data/sqlite-latest.sqlite  # ~529MB

# EVE ESI Rate Limiting Check
curl -I http://localhost:9000/api/v1/esi/test  # Headers prüfen
```

### EVE Docker Commands (geplant Phase 6)
```bash
# EVE Container Build (zukünftig)
docker build -t eve-profit2 .

# EVE Production Deploy (zukünftig)  
docker-compose up -d eve-backend
```

### ✅ Phase 3 Abgeschlossen (Aktuell)
- **ESI Integration:** Vollständig implementiert mit Rate Limiting
- **Server:** Port 9000, Production-ready mit Gin Framework
- **Tests:** 18/18 bestehen, ~90% Coverage
- **SDE:** SQLite Integration für Item Data
- **Config:** EVE Application Settings (Client ID/Secret)

### 🎯 Phase 4 Bereit (Nächster Schritt)
- **Ziel:** Market Data & Item Search API Handlers
- **ESI Endpoints:** Market Orders, Item Search, Character Data
- **Basis:** ESI Client + Service Layer + SDE Repository fertig
- **Character Auth:** EVE SSO OAuth Flow Implementation

### 🚀 Phase 5 Geplant (Frontend)
- **React + TypeScript:** Trading Dashboard UI
- **EVE Integration:** Character Login, Market Analysis
- **Real-time Updates:** WebSocket für Live Market Data

## 📁 EVE-spezifische Dokumentation

### EVE Business Logic
- `PROJECT_STATUS.md` - **ZENTRAL:** Entwicklungsstand + EVE API Integration
- `PROJECT_CONTEXT.md` - EVE Online Business Context + Trading Logic
- `CHARACTER_API_SPECS.md` - EVE SSO OAuth + Character API Details

### Universal Guidelines (wiederverwendbar)
- `UNIVERSAL_SESSION_MANAGEMENT_GUIDELINES.md` - Session Management für alle Projekte
- `UNIVERSAL_DEVELOPMENT_GUIDELINES.md` - Projektmanagement + Entwicklungsprozesse
- `UNIVERSAL_TESTING_GUIDELINES.md` - TDD-Patterns + Test-Strategien
- `UNIVERSAL_CLEAN_CODE_GUIDELINES.md` - Clean Code + SOLID Prinzipien

---

**💡 Session Management:** Nutze die Universal Guidelines für allgemeine Patterns, diese Datei für EVE-spezifische Commands und PROJECT_STATUS.md für aktuellen Entwicklungsstand! 🚀
