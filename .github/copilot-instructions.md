# GitHub Copilot Instructions - EVE Profit Calculator 2.0

## 🎯 Projekt-Kontext
EVE Online Trading Calculator mit Go-Backend + React-Frontend, Clean Code + TDD.

## 🏗️ Tech Stack
- **Backend:** Go + Gin + SQLite SDE + ESI API  
- **Frontend:** React 19 + TypeScript + Vite + Testing Library
- **Tests:** 135 Tests (31 Backend + 19 Frontend + 85 E2E)

## 🎯 Coding Rules

### Clean Code (aus docs/UNIVERSAL_CLEAN_CODE_GUIDELINES.md):
- **Meaningful Names:** `calculateNetProfit()` nicht `calc()`
- **Single Responsibility:** Eine Funktion = eine Aufgabe
- **Dependency Injection:** Interfaces für Testbarkeit
- **Error Handling:** Strukturierte Responses, nie unbehandelt

> **Erweiterte Patterns:** Siehe `docs/UNIVERSAL_CLEAN_CODE_GUIDELINES.md` für EVE-spezifische Business Logic

### TDD Workflow (aus docs/UNIVERSAL_TESTING_GUIDELINES.md):
- **RED:** Test schreiben (fails)
- **GREEN:** Minimal code (passes)  
- **REFACTOR:** Clean up
- **Arrange-Act-Assert** Pattern immer

> **Erweiterte Test-Patterns:** Siehe `docs/UNIVERSAL_TESTING_GUIDELINES.md` für komplexe Testing-Szenarien

## 🧪 Test-Standards
- **Tests FIRST** - niemals Code ohne Test
- **Mock Dependencies** - ESI API, Database, Services
- **All Tests Pass** - 135/135 müssen grün sein

## 🛡️ **ZERO-TOLERANCE Test Policy (VERBINDLICH)**

### **🚨 Absolute Regel: NO FEATURE DEVELOPMENT ohne 100% grüne Tests**

**VOR jeder Feature-Erweiterung oder Code-Änderung:**

1. **ALLE lokalen Tests müssen bestehen:**
   ```bash
   # Backend Tests (Go)
   cd backend && go test -v ./...
   
   # Frontend Tests (React) 
   cd frontend && npm run test:run
   
   # E2E Tests (lokal)
   npx playwright test
   ```

2. **VOLLSTÄNDIGE CI/CD-Pipeline muss grün sein:**
   ```bash
   gh run list --limit 1    # Letzte Pipeline prüfen
   gh run view <id> --log   # Bei Fehlern detailliert analysieren
   ```

3. **Quality Gates müssen erfüllt sein:**
   - ✅ Backend Tests: 31+ Tests bestehen
   - ✅ Frontend Tests: 36+ Tests bestehen  
   - ✅ E2E Tests: 85+ Tests bestehen
   - ✅ ESLint: 0 Errors
   - ✅ TypeScript: 0 Errors
   - ✅ Security Scan: Pass

### **🚫 ENTWICKLUNGS-STOPP bei roten Tests**

- **KEINE neuen Features** bis alle Tests grün
- **KEINE Code-Commits** mit failing Tests
- **KEINE Pipeline-Ignorierung** - Fehler SOFORT beheben
- **IMMER Fix-First** - repariere kaputte Tests vor neuer Arbeit

### **⚡ Schneller Test-Status-Check**
```bash
# All-in-One Test Status Check
./scripts/check-all-tests.sh   # falls vorhanden
# oder manuell:
cd backend && go test ./... && cd ../frontend && npm run test:run && cd .. && gh run list --limit 1
```

> **Development Setup:** Siehe `docs/PROJECT_SESSION_MANAGEMENT.md` für Commands und Server-Management

## 💻 Code Patterns

**Go Backend:**
```go
// Interface-based Design
type ItemService interface {
    GetItem(ctx context.Context, id int) (*models.Item, error)
}

// Structured Error Responses
if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item"})
    return
}
```

**React Frontend:**
```typescript
// Full TypeScript typing
interface ItemSearchProps {
  onItemSelect: (item: EVEItem) => void;
}

// React Testing Library
test('should render input', () => {
  render(<ItemSearch onItemSelect={vi.fn()} />);
  expect(screen.getByRole('textbox')).toBeInTheDocument();
});
```

## 🎯 EVE Domain
- **SDE:** 25.818 Items (Tritanium, Veldspar, etc.)
- **ESI:** 150 req/sec Rate Limit
- **Trading:** Buy/Sell margins, taxes, broker fees

> **API Specs:** Siehe `docs/PROJECT_API_SPECS.md` für vollständige Endpoint-Dokumentation  
> **Security:** Siehe `docs/PROJECT_SECURITY_GUIDELINES.md` für ESI Credential Management

## 🚫 Vermeide
- Code ohne Tests
- Breaking existing tests  
- Hardcoded values
- Unhandled errors
- **Feature-Development bei roten Tests** ⚠️
- **Pipeline-Ignorierung bei Fehlern** ⚠️

## 📚 Dokumentations-Regeln
- **Projekt-Status-Änderungen** IMMER in `docs/PROJECT_STATUS.md` dokumentieren
- Neue Features, abgeschlossene Phasen, Test-Updates → PROJECT_STATUS.md
- Technische Entscheidungen und Architektur-Änderungen dokumentieren

> **Extended Guidelines:** Siehe `docs/UNIVERSAL_DEVELOPMENT_GUIDELINES.md` für Projektmanagement-Standards

## ✅ Quality Gates
- All 135 Tests pass
- TypeScript: 0 errors
- ESLint: 0 warnings
- SonarQube: 0 issues
- **CI/CD Pipeline: 100% grün** 🎯

## 🔧 **Modulare CI/CD Pipeline (3-Phasen-Architektur)**

### 📋 **Pipeline-Phasen (Sequenziell)**

#### **Phase 1: CI Pipeline (`ci.yml`)**
```bash
# Schnelle Feedback-Schleife für Code-Qualität
gh run list --workflow="CI Pipeline" --limit 3
```

**Pipelines:**
- ✅ **Frontend Tests:** TypeScript, ESLint, React Tests (18s)
- ✅ **Backend Tests:** Go Tests mit SDE Database (1m30s)  
- ✅ **E2E Tests:** Playwright mit Backend+Frontend (4m)
- ✅ **Security Scan:** Trivy Vulnerability Scanner (25s)

#### **Phase 2: Docker Build (`docker-build.yml`)**
```bash
# Triggered automatisch nach erfolgreicher CI
gh run list --workflow="Docker Build" --limit 3
```

**Container Builds:**
- 🐳 **Backend Docker:** Go 1.23-alpine + CGO + SQLite (3m7s)
- 🐳 **Frontend Docker:** Node.js + nginx Production Build (27s)
- 📦 **Registry Push:** ghcr.io mit separaten Caches

#### **Phase 3: Deployment (`deploy.yml`)**
```bash
# Production Deployment (main branch only)
gh run list --workflow="Deploy" --limit 3
```

**Deployment Flow:**
- 🚀 **Production Deploy:** Automated nach Docker Build Success (20s)
- 📢 **Notifications:** Deployment Status Reporting

### 🚀 **Modulare Pipeline-Überwachung**

#### **Pipeline Status Check Workflow:**
1. **CI Phase:** `gh run list --workflow="CI Pipeline"` - Alle Tests müssen grün sein
2. **Docker Phase:** `gh run list --workflow="Docker Build"` - Container müssen erfolgreich gebaut werden  
3. **Deploy Phase:** `gh run list --workflow="Deploy"` - Production Deployment Success
4. **Bei Fehlern:** `gh run view <id> --log` für detaillierte Analyse
5. **Reparatur:** Fehler beheben BEVOR weitere Entwicklung

#### **Vorteile der Modularen Architektur:**
- **⚡ Schnellere Feedback:** CI läuft parallel (Frontend + Backend + E2E)
- **🔧 Isolierte Debugging:** Separate Workflows für einfachere Fehlerdiagnose
- **📈 Bessere Skalierung:** Docker Builds laufen nur bei CI Success
- **🚀 Flexible Deployments:** Manuelle Triggers für Emergency Deployments
- **💾 Optimierte Caches:** Separate Cache-Scopes für Backend/Frontend Docker Builds

### ⚡ **Schnelle Problemlösung**

#### **Bei ESLint Fehlern:**
```typescript
// ❌ Falsch
private handleError(error: any): void

// ✅ Richtig  
private handleError(error: unknown): void {
  const typedError = error as { message?: string };
  // ...
}
```

#### **Bei Pipeline Fehlern:**
1. **Logs prüfen:** `gh run view <id> --log`
2. **Lokale Tests:** `npm run lint && npm run type-check`
3. **Reparatur committen:** Nur Fehler-Fixes, keine neuen Features
4. **Pipeline erneut prüfen:** Bestätigung des Erfolgs

### 🚨 **Wichtige Erinnerungen**

- **NIEMALS** Pipeline-Fehler ignorieren
- **IMMER** ESLint-Errors vor Commit beheben  
- **IMMER** TypeScript-Errors vor Commit beheben
- **NIEMALS** `any` in TypeScript verwenden
- **IMMER** CI/CD Status nach Push überprüfen
- **SOFORT** Reparaturen bei roten Pipelines

> **Extended Pipeline Documentation:** Siehe `docs/CI_CD_PIPELINE_ARCHITECTURE.md` für detaillierte Architektur-Specs
