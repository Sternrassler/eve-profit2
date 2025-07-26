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

## 🔧 **CI/CD Pipeline Rules (VERBINDLICH)**

### 📋 **Pflichtprüfungen vor jedem Commit**

#### 1. **CI/CD Pipeline Status prüfen**
```bash
# IMMER ausführen vor Code-Änderungen
gh run list --limit 3
gh run view <run-id> --log  # bei Fehlern
```

**Regel:** Jede Code-Änderung muss erst nach erfolgreicher Pipeline-Prüfung committed werden.

#### 2. **Frontend Linting (Obligatorisch)**
```bash
cd frontend
npm run lint          # ESLint Prüfung
npm run type-check     # TypeScript Validation
```

**Regel:** Alle ESLint-Fehler müssen VOR dem Commit behoben werden. Keine Ausnahmen.

#### 3. **Backend Testing (Obligatorisch)**
```bash
cd backend
go test -v ./...      # Go Tests
go vet ./...          # Go Static Analysis
```

**Regel:** Alle Tests müssen grün sein. Failing Tests blockieren den Commit.

### 🚀 **CI/CD Pipeline Monitoring**

#### **Pipeline Status Check Workflow:**
1. **Nach jedem Push:** `gh run list` ausführen
2. **Bei Fehlern:** Detaillierte Logs mit `gh run view <id> --log` prüfen
3. **Reparatur:** Fehler beheben BEVOR weitere Entwicklung
4. **Validierung:** Pipeline muss grün sein vor dem nächsten Feature

#### **Häufige Pipeline-Fehler beheben:**
- **ESLint Errors:** TypeScript `any` durch konkrete Typen ersetzen  
- **Empty Interfaces:** Leere Interfaces entfernen oder erweitern
- **Deprecated Actions:** GitHub Actions auf neueste Versionen aktualisieren
- **YAML Syntax:** GitHub Actions Workflow-Dateien validieren

### 📊 **Automatisierte Quality Gates**

Das Projekt verwendet folgende automatisierte Prüfungen:

1. **Backend Tests (Go):** 31 Tests müssen bestehen
2. **Frontend Tests (React):** 36 Tests müssen bestehen  
3. **E2E Tests (Playwright):** 85 Tests müssen bestehen
4. **Security Scanning:** Trivy Vulnerability Scanner
5. **Docker Build:** Multi-stage Builds für Backend + Frontend
6. **Code Coverage:** Codecov Integration für Metriken

**Ziel:** 152/153 Tests bestehen (99.3% Success Rate)

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

---
**TDD + Clean Code + EVE Domain Knowledge + 100% Test Coverage**
