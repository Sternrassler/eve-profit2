# EVE Profit Calculator 2.0 - Projektkontext

## 📋 Projektübersicht
**Name:** EVE Profit Calculator 2.0  
**Zweck:** Moderne Webanwendung für EVE Online Handelsanalysen und Profit-Berechnungen  
**Erstellt:** 19. Juli 2025  
**Entwickler:** Karsten Flache  

## 🎯 Projektziele
- Marktdatenanalyse zwischen verschiedenen EVE Online Stationen
- Profit-Berechnungen für Trading-Routen
- Item-Lookup mit Live-Preisdaten
- Trend-Analysen und Preishistorie
- Benutzerfreundliche, moderne Oberfläche

## 🛠️ Technologie-Stack
### Frontend
- **Framework:** React 18+ mit TypeScript
- **Build Tool:** Vite
- **Styling:** Tailwind CSS
- **Icons:** Lucide React oder Heroicons
- **State Management:** React Hooks + Context API

### Backend/API
- **Backend:** Go (Golang) REST API
- **Datenquelle:** EVE ESI (EVE Swagger Interface) 
- **HTTP Client:** Go net/http mit Goroutines für Parallelisierung
- **Caching:** In-Memory Cache (BigCache)
- **Database:** SDE SQLite (EVE Static Data Export)
- **OAuth:** EVE SSO für Character Authentication
- **API Format:** JSON REST API

### Entwicklung
- **Package Manager:** npm
- **Linting:** ESLint + Prettier
- **Testing:** Vitest + React Testing Library

## 🔧 Architektur-Entscheidungen
### Backend-Struktur (Go)
```
backend/
├── cmd/
│   └── server/          # Main Application Entry
├── internal/
│   ├── api/            # HTTP Handlers & Routes
│   ├── service/        # Business Logic Layer
│   ├── repository/     # Data Access Layer (SDE SQLite)
│   ├── cache/          # In-Memory Caching Implementation
│   └── models/         # Data Models & Types
├── pkg/
│   ├── eve/            # EVE ESI Client
│   ├── sde/            # SDE SQLite Client
│   ├── config/         # Configuration Management
│   └── utils/          # Shared Utilities
├── data/
│   └── sde.sqlite      # EVE Static Data Export Database
└── migrations/         # SDE Update Scripts
```

### Frontend-Struktur
```
frontend/src/
├── components/          # Wiederverwendbare UI-Komponenten
├── pages/              # Hauptseiten/Views
├── hooks/              # Custom React Hooks
├── services/           # API-Services (Backend Calls)
├── types/              # TypeScript Type-Definitionen
├── utils/              # Hilfsfunktionen
├── context/            # React Context für State
└── styles/             # Globale Styles
```

### Naming Conventions
**Backend (Go):**
- **Packages:** lowercase (z.B. `eveapi`, `marketdata`)
- **Functions:** PascalCase für öffentliche, camelCase für private
- **Structs:** PascalCase (z.B. `MarketOrder`, `ItemData`)
- **Interfaces:** PascalCase mit -er Suffix (z.B. `CacheProvider`)

**Frontend (TypeScript):**
- **Komponenten:** PascalCase (z.B. `MarketAnalysis.tsx`)
- **Hooks:** camelCase mit "use" Prefix (z.B. `useMarketData.ts`)
- **Services:** camelCase (z.B. `backendApiService.ts`)
- **Types:** PascalCase mit Type-Suffix (z.B. `MarketDataType.ts`)

## 🌐 EVE Online API Integration
### Go Backend - ESI Client Design
```go
// Parallelisierte ESI-Abfragen mit Worker Pools
type ESIClient struct {
    httpClient   *http.Client
    rateLimiter  *rate.Limiter
    cache        CacheProvider
    workerPool   int
    authClient   *OAuthClient  // Für Character API calls
}

// Beispiel für parallele Marktdaten-Abfrage
func (c *ESIClient) GetMarketDataParallel(regions []int32, typeID int32) 

// Character API mit OAuth Token
func (c *ESIClient) GetCharacterSkills(characterID int32, token string) (*CharacterSkills, error)
func (c *ESIClient) GetCharacterAssets(characterID int32, token string) ([]Asset, error)
```

### ESI Endpoints (Wichtigste)
**Market Data:**
- Market Orders: `/markets/{region_id}/orders/`
- Market History: `/markets/{region_id}/history/`
- Market Groups: `/markets/groups/`
- Universe Types: `/universe/types/`
- Universe Stations: `/universe/stations/`

**Character Data (OAuth required):**
- Character Info: `/characters/{character_id}/`
- Character Skills: `/characters/{character_id}/skills/`
- Character Assets: `/characters/{character_id}/assets/`
- Character Ships: `/characters/{character_id}/ships/`
- Character Wallet: `/characters/{character_id}/wallet/`
- Character Orders: `/characters/{character_id}/orders/`

### Wichtige EVE-Konzepte
- **Region IDs:** Numerische IDs für Handelsregionen
- **Type IDs:** Numerische IDs für Items/Schiffe
- **Station IDs:** Numerische IDs für Handelsstationen
- **Security Status:** High-Sec (≥0.5), Low-Sec (0.1-0.4), Null-Sec (≤0.0)

## 📊 Kerntfunktionen
### 1. Marktanalyse
- Preisvergleiche zwischen Regionen
- Buy/Sell Order Analyse
- Volumen-Analysen

### 2. Profit-Berechnung
- Trading-Margin Berechnung
- Transport-Kosten berücksichtigen
- Steuer- und Broker-Gebühren

### 3. Item-Suche
- Autocomplete Item-Suche
- Kategorie-basierte Filterung
- Favoriten-System

### 4. Benutzereinstellungen
- Bevorzugte Handelsregionen
- Standard-Schiffstypen für Transport
- Steuer-/Broker-Settings

### 5. Charakter-Integration
- EVE Character Login (OAuth)
- Skill-basierte Trading-Optimierungen
- Asset-Tracking (verfügbare Schiffe & Fracht)
- Wallet-Balance Integration
- Aktuelle Market Orders des Charakters

## 🎨 UI/UX Richtlinien
### Design-Prinzipien
- **Dark Theme:** EVE Online inspiriertes dunkles Design
- **Responsiv:** Mobile-first Ansatz
- **Performance:** Schnelle Ladezeiten, effiziente Updates
- **Accessibility:** WCAG 2.1 AA konform

### Farbschema
- **Primary:** Blau-töne (EVE-typisch)
- **Secondary:** Orange/Gold für Highlights
- **Background:** Dunkle Grau-töne
- **Text:** Helle Farben für Kontrast

## 🔒 Sicherheit & Performance
### Go Backend - Parallelisierung
- **Goroutines:** Worker Pools für ESI-Abfragen (Standard: 20 Workers)
- **Rate Limiting:** Token Bucket für 150 req/sec ESI-Limit
- **Context Timeout:** 30s für einzelne ESI-Calls, 2min für Batch-Operationen

### Caching-Strategie
- **BigCache:** In-Memory Cache für alle ESI-Daten
- **SDE SQLite:** Statische EVE-Daten (Items, Stations, Regionen)
- **TTL:** 5min Marktdaten, 1h Preishistorie
- **Cache Keys:** `market:{region}:{type}`, `item:{typeID}`, `station:{stationID}`

### Ratenlimits
- EVE ESI: Max. 150 Requests/Sekunde
- Backend API: Rate Limiting pro Client-IP
- Error Handling für API-Ausfälle mit Exponential Backoff

### Datenvalidation
- **Go:** Struct Tags für Validation
- **TypeScript:** Compile-time Checks
- **Zod:** Runtime-Validation im Frontend
- **Input-Sanitization:** SQL Injection Prevention

## 📝 Entwicklungsstandards
### Code-Qualität
- Alle Komponenten in TypeScript
- Minimum 80% Test-Abdeckung
- ESLint + Prettier für konsistente Formatierung
- Komponenten-Dokumentation mit JSDoc

### Git-Workflow
- Feature-Branches für neue Funktionen
- Conventional Commits für aussagekräftige Messages
- Pull Request Reviews vor Merge

## 🚀 Deployment-Ziele
- **Hosting:** Vercel oder Netlify für Frontend
- **CI/CD:** GitHub Actions
- **Monitoring:** Einfache Analytics
- **Performance:** Lighthouse Score >90

## 📋 Nächste Schritte
1. Go Backend Setup (Gin + BigCache + SDE SQLite)
2. SDE Database Integration & Item-Lookup
3. ESI Client mit Goroutines implementieren  
4. React Frontend Setup (Vite + TypeScript + Tailwind)
5. Basis-Komponenten entwickeln
6. Marktdaten-Integration & UI Implementation

---
**Wichtig:** Diese Datei sollte bei jeder Session als Referenz gelesen werden, um Konsistenz in Entwicklungsentscheidungen zu gewährleisten.
