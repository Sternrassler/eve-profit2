# EVE Profit Calculator 2.0 - Projektkontext

## 📋 Projektübersicht
**Name:** EVE Profit Calculator 2.0  
**Zweck:** Moderne Webanwendung für EVE Online Handelsanalysen und Profit-Berechnungen  
**Erstellt:** 19. Juli 2025  
**Entwickler:** Karsten Flache  
**Entwicklungsmethodik:** Clean Code + Test-Driven Development (TDD)

## 🎯 Projektziele
- Marktdatenanalyse zwischen verschiedenen EVE Online Stationen
- Profit-Berechnungen für Trading-Routen
- Item-Lookup mit Live-Preisdaten
- Trend-Analysen und Preishistorie
- Benutzerfreundliche, moderne Oberfläche
- **Code-Qualität:** Maintainable, self-documenting, highly tested codebase

## 📚 Code-Standards
**Vollständige Entwicklungsstandards siehe:**
- `UNIVERSAL_CLEAN_CODE_GUIDELINES.md` - Clean Code + SOLID Prinzipien
- `UNIVERSAL_TESTING_GUIDELINES.md` - TDD Workflows + Test-Patterns
- `UNIVERSAL_DEVELOPMENT_GUIDELINES.md` - Projektmanagement + Code-Review

**Ziel:** Maintainable, self-documenting, highly tested codebase

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
- **Server Port:** 9000 (Backend), 3000 (Frontend Dev Server)

### EVE ESI Application Settings
- **Client ID:** 0928b4bcd20242aeb9b8be10f5451094
- **Client Secret:** AQPjLZ3VYAewR59J5jStZs52dY7jISGVLwXv5NA
- **Callback URL:** http://localhost:9000/callback
- **Scopes:**
  - publicData
  - esi-location.read_location.v1
  - esi-location.read_ship_type.v1
  - esi-skills.read_skills.v1
  - esi-wallet.read_character_wallet.v1
  - esi-universe.read_structures.v1
  - esi-assets.read_assets.v1
  - esi-fittings.read_fittings.v1
  - esi-characters.read_standings.v1

### Entwicklung
- **Package Manager:** npm
- **Linting:** ESLint + Prettier
- **Testing:** Vitest + React Testing Library

## 🎯 EVE Online Business Context

### Wichtige EVE-Konzepte für das Projekt
- **ISK:** EVE-Währung (Inter Stellar Kredits)
- **Jita:** Haupt-Handelshub (System-ID: 30000142)
- **Market Orders:** Buy-Orders (bids) vs Sell-Orders (asks)
- **Station Trading:** Profit durch Kauf/Verkauf in derselben Station
- **Inter-Region Trading:** Profit durch Transport zwischen Regionen
- **Broker Fees:** 3% NPC-Station Gebühren
- **Sales Tax:** 1-8% basierend auf Character Skills

### EVE ESI API Integration
- **Market Data:** Live Orders & Historical Price Data
- **Character Data:** Skills, Assets, Wallet (OAuth required)
- **Universe Data:** Stations, Regions, Item Types
- **Rate Limiting:** 150 requests/second maximum

### Gewinn-Berechnungslogik
```
Total Profit = (Sell Price - Buy Price) * Quantity 
             - Broker Fees 
             - Sales Tax 
             - Transport Costs
```
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
