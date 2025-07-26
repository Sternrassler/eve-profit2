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

---
**TDD + Clean Code + EVE Domain Knowledge + 100% Test Coverage**
