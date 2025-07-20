# Universal Session Management Guidelines

> **Für jedes Softwareprojekt:** Schnelle Wiederaufnahme nach Entwicklungspausen

## 🔄 Universelle Wiederaufnahme Commands

### Projekt fortsetzen (Backend Development)
```
"Lies STATUS.md. [Aktuelle Phase] ist abgeschlossen. 
Implementiere jetzt [Nächste Phase] mit [Spezifische Technologie/Integration]."
```

### Frontend Development starten
```
"Lies STATUS.md. Backend läuft auf Port [PORT] mit [API/Service] Integration. 
Erstelle [Frontend Framework] mit Backend API Integration."
```

### Schneller Status-Check
```
"Zeige aktuellen Projekt-Status und nächste Schritte"
```

### Code-Review und Testing
```
"Führe Code-Review durch. Prüfe Clean Code Standards und Test-Coverage."
```

### Deployment vorbereiten
```
"Lies STATUS.md. Bereite Deployment vor: Docker, CI/CD, Production-Config."
```

## 🛠️ Universelle Development Commands

### Server & Testing (Go)
```bash
cd backend
go run cmd/server/main.go     # Server starten
go test ./...                 # Alle Tests ausführen
go test -v ./...              # Tests mit Details
go test -cover ./...          # Test Coverage anzeigen
```

### Server & Testing (Node.js)
```bash
npm start                     # Server starten
npm test                      # Tests ausführen
npm run test:watch            # Tests im Watch-Mode
npm run test:coverage         # Test Coverage
```

### Server & Testing (Python)
```bash
python -m uvicorn main:app --reload  # FastAPI Server
python -m pytest             # Tests ausführen
python -m pytest --cov       # Test Coverage
python -m pytest -v          # Tests mit Details
```

### API Testing (Universal)
```bash
curl http://localhost:[PORT]/api/health           # Health Check
curl http://localhost:[PORT]/api/v1/[endpoint]    # API Test
curl -X POST http://localhost:[PORT]/api/[endpoint] \
  -H "Content-Type: application/json" \
  -d '{"key": "value"}'                           # POST Request
```

### Git Workflow
```bash
git status                    # Änderungen anzeigen
git add -A                    # Alle Änderungen stagen
git commit -m "feat: [feature description]"  # Commit erstellen
git push                      # Änderungen pushen
git log --oneline -10         # Letzte Commits anzeigen
```

### Database Commands (Universal)
```bash
# PostgreSQL
psql -d [database] -c "SELECT version();"

# SQLite
sqlite3 [database.db] ".tables"

# Docker Database
docker exec -it [container] psql -U [user] -d [database]
```

## 📊 Universelle Projekt-Checkpoints Template

### ✅ [Phase Name] Abgeschlossen (Aktuell)
- **[Hauptfeature]:** Vollständig implementiert
- **Server:** Port [PORT], Production-ready
- **Tests:** [X]/[Y] bestehen, ~[Z]% Coverage
- **Config:** [Konfiguration] integriert

### 🎯 [Nächste Phase] Bereit (Nächster Schritt)
- **Ziel:** [Hauptziel] mit [Technologie/Pattern]
- **Basis:** [Voraussetzungen] fertig
- **Geschätzt:** [Zeitschätzung] mit Clean Code + TDD

### 🚀 [Deployment Phase] Vorbereitung
- **Infrastructure:** Docker, CI/CD Pipeline
- **Performance:** Optimierung und Load Testing
- **Security:** Security Audit und Penetration Testing

## 📁 Universelle Dokumentationsstruktur

### Kern-Dokumentation (Empfohlene Struktur)
- `STATUS.md` - **ZENTRAL:** Entwicklungsstand + Architektur + Phasen
- `UNIVERSAL_DEVELOPMENT_GUIDELINES.md` - Projektmanagement + Entwicklungsprozesse
- `UNIVERSAL_TESTING_GUIDELINES.md` - TDD-Patterns + Test-Strategien
- `UNIVERSAL_CLEAN_CODE_GUIDELINES.md` - Clean Code + SOLID Prinzipien
- `PROJECT_CONTEXT.md` - Business Context + Domain-Spezifika
- `[FEATURE]_SPECS.md` - Spezifische Feature-Spezifikationen

### Technische Dokumentation
- `README.md` - Setup, Installation, Quick Start
- `ARCHITECTURE.md` - System-Architektur und Design Decisions
- `API_DOCS.md` - API Endpoints und Specifications
- `DEPLOYMENT.md` - Deployment Guide und Infrastructure

### Development Support
- `TROUBLESHOOTING.md` - Häufige Probleme und Lösungen
- `CONTRIBUTING.md` - Development Workflow für Teams
- `CHANGELOG.md` - Versionshistorie und Breaking Changes

## 🔄 Session Transition Patterns

### Nach längerer Pause (> 1 Woche)
1. **STATUS.md lesen** - Aktueller Stand verstehen
2. **Tests ausführen** - System-Integrität prüfen
3. **Dependencies prüfen** - Updates und Security Patches
4. **Entwicklungsumgebung validieren** - Tools und Konfiguration

### Zwischen Features wechseln
1. **Current Branch committen** - Aktueller Stand sichern
2. **STATUS.md aktualisieren** - Fortschritt dokumentieren
3. **Neue Branch erstellen** - Feature-spezifischer Branch
4. **Relevante Specs lesen** - Feature-Anforderungen verstehen

### Code Review Session
1. **Clean Code Guidelines prüfen** - Standards-Compliance
2. **Test Coverage analysieren** - Qualitätskontrolle
3. **Security Review** - Vulnerability Assessment
4. **Performance Review** - Bottleneck Identifikation

### Deployment Session
1. **Alle Tests grün** - Quality Gate
2. **Security Audit** - Production Readiness
3. **Performance Testing** - Load und Stress Tests
4. **Rollback Plan** - Disaster Recovery

## 📋 Universal Session Checklist

### Session Start:
- [ ] STATUS.md gelesen und verstanden
- [ ] Entwicklungsumgebung funktionsfähig
- [ ] Alle Tests laufen grün
- [ ] Git Status sauber (oder bewusst staged)
- [ ] Aktuelle Phase und Ziele klar

### During Development:
- [ ] TDD Red-Green-Refactor befolgt
- [ ] Clean Code Standards eingehalten
- [ ] Regelmäßige Commits mit aussagekräftigen Messages
- [ ] Test Coverage maintained (>90% für Business Logic)

### Session End:
- [ ] Alle Änderungen committed
- [ ] STATUS.md aktualisiert (falls nötig)
- [ ] Tests laufen grün
- [ ] Nächste Schritte dokumentiert
- [ ] Session Notes für Wiederaufnahme

## 🎯 Language-Specific Quick Commands

### Go Projects
```bash
go mod tidy                   # Dependencies aufräumen
go fmt ./...                  # Code formatieren
go vet ./...                  # Static Analysis
go build ./cmd/server         # Binary erstellen
```

### Node.js/TypeScript Projects
```bash
npm ci                        # Clean Install
npm run lint                  # Linting
npm run format                # Code formatieren
npm run build                 # Production Build
```

### Python Projects
```bash
pip install -r requirements.txt  # Dependencies installieren
black .                       # Code formatieren
flake8 .                      # Linting
mypy .                        # Type Checking
```

### Docker Projects
```bash
docker-compose up -d          # Services starten
docker-compose logs -f        # Logs verfolgen
docker-compose down           # Services stoppen
docker system prune          # Cleanup
```

---

## 🎯 **Universelle Session Management**

**Diese Guidelines funktionieren für jedes Softwareprojekt und jede Technologie. Das Ziel ist effizienter Kontext-Switch zwischen Entwicklungssessions.**

**Anpassung:** Ersetze Platzhalter wie `[PORT]`, `[Phase]`, `[Technologie]` mit projektspezifischen Werten in deinem lokalen Setup. Die Referenzen zu `STATUS.md` und `SESSION_MANAGEMENT.md` sollten in deinem Projekt zu `PROJECT_STATUS.md` und `PROJECT_SESSION_MANAGEMENT.md` werden.**
