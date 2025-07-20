# Universal Development Guidelines

## 🎯 Für neue Sessions - Sofort lesen!

Diese universellen Entwicklungsrichtlinien können für jedes Softwareprojekt verwendet werden und müssen bei jeder neuen Entwicklungssession beachtet werden.

## 📖 Obligatorische Session-Vorbereitung
**Vor jeder Entwicklungsarbeit:**
1. Projektdokumentation lesen (README, STATUS, etc.)
2. Diese Development Guidelines lesen
3. Testing Guidelines verstehen
4. Aktuellen Code-Stand überprüfen

## 🔧 Universelle Technische Standards

## 📚 Grundlagen - Obligatorische Referenzen

**Vor jeder Entwicklungsarbeit müssen diese beiden Dokumente gelesen werden:**

### 🎯 Clean Code Prinzipien
**Siehe: `UNIVERSAL_CLEAN_CODE_GUIDELINES.md`**

Diese umfassende Referenz enthält alle Clean Code Prinzipien:
- Meaningful Names & Selbstdokumentierender Code
- Single Responsibility Functions
- SOLID Prinzipien (SRP, OCP, LSP, ISP, DIP)
- Code-Struktur und Formatting Standards
- Universelle Patterns und Best Practices

### 🧪 Test-Driven Development (TDD)
**Siehe: `UNIVERSAL_TESTING_GUIDELINES.md`**

Diese umfassende TDD-Referenz enthält:
- Red-Green-Refactor Zyklus
- AAA Pattern (Arrange-Act-Assert)
- Test-Organisation und Naming Conventions
- Mocking Best Practices
- Framework-spezifische Beispiele (Python, TypeScript, Go, Java, C#)

## 💼 Projektmanagement Standards

### Universelle Konfigurationsstandards

#### Code-Qualität Tools
- **Linter:** ESLint (JS/TS), golangci-lint (Go), pylint (Python)
- **Static Analysis:** SonarQube (kontinuierliche Qualitätsüberwachung)
- **Formatter:** Prettier (JS/TS), gofmt (Go), black (Python)
- **Type Checking:** TypeScript strict mode, mypy (Python)

#### Git Best Practices
```bash
# Commit Message Format
type(scope): description

# Examples:
feat(auth): add user authentication
fix(api): resolve null pointer exception
refactor(service): extract user validation logic
test(calculator): add edge case tests
docs(readme): update installation instructions
```

#### Testing Patterns
- **Unit Tests:** Einzelne Funktionen/Methoden (TDD)
- **Integration Tests:** Zusammenspiel zwischen Komponenten
- **End-to-End Tests:** Kritische User-Workflows
- **Coverage Minimum:** 80-90% für Business Logic

## 🎨 Universelle Code-Struktur Standards

### Projekt-Organisation

#### Moderne Backend/Frontend-Trennung
```
project-root/
├── backend/                # Backend Services
│   ├── cmd/               # Entry Points (main.go, etc.)
│   │   └── server/        # Server Startup
│   ├── internal/          # Private Application Code
│   │   ├── api/           # API Layer (REST/GraphQL Handlers)
│   │   │   ├── handlers/  # HTTP Request Handlers
│   │   │   └── middleware/# HTTP Middleware
│   │   ├── service/       # Business Logic Services
│   │   ├── repository/    # Data Access Layer
│   │   ├── models/        # Data Models/Entities
│   │   ├── config/        # Configuration Management
│   │   └── cache/         # Caching Layer
│   ├── pkg/               # Public Libraries (shared packages)
│   ├── tests/             # Backend Tests
│   │   ├── unit/          # Unit Tests
│   │   ├── integration/   # Integration Tests
│   │   └── fixtures/      # Test Data
│   └── scripts/           # Backend Build/Deploy Scripts
├── frontend/              # Frontend Application
│   ├── src/               # Source Code
│   │   ├── components/    # UI Components
│   │   ├── pages/         # Page Components
│   │   ├── services/      # API Client Services
│   │   ├── stores/        # State Management
│   │   ├── utils/         # Helper Functions
│   │   └── types/         # TypeScript Type Definitions
│   ├── tests/             # Frontend Tests
│   │   ├── unit/          # Component Unit Tests
│   │   ├── integration/   # Component Integration Tests
│   │   └── e2e/           # End-to-End Tests
│   └── public/            # Static Assets
├── docs/                  # Project Documentation
├── config/                # Shared Configuration
└── scripts/               # Shared Scripts
```

#### Alternative: Monorepo-Struktur
```
project-root/
├── apps/                  # Applications
│   ├── api/              # Backend API
│   ├── web/              # Frontend Web App
│   └── mobile/           # Mobile App (optional)
├── packages/             # Shared Libraries
│   ├── types/            # Shared Types
│   ├── utils/            # Shared Utilities
│   └── config/           # Shared Configuration
├── tools/                # Development Tools
└── docs/                 # Documentation
```

### Naming Conventions

#### SonarQube-Konforme Naming Standards
**Kritisch für Code-Qualität:** SonarQube überwacht kontinuierlich Namenskonventionen und Code-Qualität

```typescript
// ✅ SonarQube-konforme Namenskonventionen

// Funktionen/Methoden: camelCase (KEINE snake_case!)
// ❌ Wrong: test_user_validation 
// ✅ Correct: testUserValidation
function calculateTotal() {}
function validateUserInput() {}

// Test-Funktionen: TestXxx oder testXxx (je nach Sprache)
// Go: TestXxx
func TestUserValidation(t *testing.T) {}
// JavaScript/TypeScript: testXxx 
function testUserValidation() {}

// Klassen: PascalCase
class UserService {}
class PaymentProcessor {}

// Constants: SCREAMING_SNAKE_CASE
const MAX_RETRY_ATTEMPTS = 3;
const DEFAULT_TIMEOUT_MS = 5000;

// Variables: camelCase
const userEmail = "user@example.com";
const orderTotal = 99.99;

// Files: kebab-case oder camelCase (konsistent pro Projekt)
user-service.ts / userService.ts
payment-processor.ts / paymentProcessor.ts
```

#### SonarQube Qualitätsregeln
```go
// ✅ Vermeide Duplicate String Literals
const testSDEPath = "../../data/sqlite-latest.sqlite"
// Verwende die Konstante statt Strings zu wiederholen

// ✅ Teste alle struct fields (vermeide "unused write" warnings)
func TestMarketOrderValidation(t *testing.T) {
    order := models.MarketOrder{
        OrderID:      12345,
        TypeID:       34,
        Price:        100.50,
        VolumeTotal:  int32(1000),    // Korrekte Datentypen verwenden!
        VolumeRemain: int32(500),
        MinVolume:    int32(1),
        // ... teste ALLE Felder
    }
    
    // Assertions für alle relevanten Felder
    assert.Equal(t, int64(12345), order.OrderID)
    assert.Equal(t, int32(34), order.TypeID)
    assert.Equal(t, 100.50, order.Price)
    assert.Equal(t, int32(1000), order.VolumeTotal)
    // ...
}
```

## 🔍 Kontinuierliche Code-Qualitätsüberwachung

### SonarQube Integration - Obligatorisch

**SonarQube überwacht permanent Code-Qualität und muss bei JEDER Verletzung korrigiert werden:**

#### Kritische SonarQube-Regeln:
```bash
# Naming Convention Violations
- Function names: camelCase (NICHT snake_case!)
- Test functions: TestXxx (Go) / testXxx (JS/TS)
- Classes: PascalCase
- Constants: SCREAMING_SNAKE_CASE

# Code Quality Violations
- Duplicate string literals → Konstanten verwenden
- Unused variables/parameters → Entfernen oder verwenden
- Incorrect data types → Exakte Typen verwenden (int32 vs int64)
- Missing test assertions → Alle struct fields testen

# Security Vulnerabilities
- SQL Injection risks → Prepared Statements
- XSS vulnerabilities → Input sanitization
- Hardcoded credentials → Environment variables
```

#### SonarQube Workflow:
```bash
# 1. Code schreiben/ändern
git add .
git commit -m "feat: implement new feature"

# 2. SonarQube-Analyse automatisch ausgeführt
# (CI/CD Pipeline oder IDE Integration)

# 3. Bei Violations: SOFORT korrigieren
# ❌ Niemals mit Violations committen!
# ✅ Alle Issues beheben vor dem nächsten Commit

# 4. Re-Analyse bis Clean Code erreicht
# Ziel: 0 Bugs, 0 Vulnerabilities, 0 Code Smells
```

#### Automatisierte Qualitätsgates:
```yaml
# .github/workflows/quality-gate.yml
name: Quality Gate
on: [push, pull_request]

jobs:
  sonarqube:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: SonarQube Scan
        uses: sonarqube-quality-gate-action@master
        with:
          scanMetadataReportFile: target/sonar/report-task.txt
        env:
          SONAR_TOKEN: ${{ secrets.SONAR_TOKEN }}
      
      # ❌ Build fails wenn SonarQube Quality Gate nicht bestanden
      - name: Quality Gate Check
        run: |
          if [ "${{ steps.sonarqube.outputs.quality-gate-status }}" != "PASSED" ]; then
            echo "Quality Gate failed!"
            exit 1
          fi
```

#### SonarQube Best Practices:
- **Immediate Fix:** Violations sofort korrigieren (nie akkumulieren)
- **Prevention:** IDE-Integration für Live-Feedback
- **Zero Tolerance:** 0 Bugs, 0 Vulnerabilities als Standard
- **Continuous Monitoring:** Automatische Scans bei jedem Commit
- **Team Training:** Alle Entwickler SonarQube-Standards kennen

**Detaillierte TDD-Guidelines siehe: `UNIVERSAL_TESTING_GUIDELINES.md`**

### Kurz-Referenz - TDD-Workflow:
```
1. 🔴 RED: Test schreiben (fehlschlagend)
2. 🟢 GREEN: Minimale Implementation für Test-Pass
3. 🔄 REFACTOR: Code verbessern ohne Tests zu brechen
4. 📝 REPEAT: Für jede neue Funktion
```

### Test-Coverage Ziele:
- **Unit Tests:** Business Logic 90%+ abgedeckt
- **Integration Tests:** Kritische Workflows getestet
- **End-to-End Tests:** User-Journeys vollständig abgedeckt

## 🚨 Universal Code-Review Checklist

### Vor jedem Commit:

#### Code-Qualität (Details: `UNIVERSAL_CLEAN_CODE_GUIDELINES.md`):
- [ ] **Clean Code Prinzipien:** Meaningful Names, SRP, SOLID befolgt
- [ ] **Selbstdokumentierender Code:** Code ist verständlich ohne Kommentare
- [ ] **DRY-Prinzip:** Keine Code-Duplikation
- [ ] **Dependency Injection:** Services über Interfaces

#### SonarQube Quality Gate (Kritisch):
- [ ] **0 Bugs:** Alle SonarQube Bug-Reports behoben
- [ ] **0 Vulnerabilities:** Alle Sicherheitslücken geschlossen
- [ ] **0 Code Smells:** Alle Code-Qualitätsprobleme behoben
- [ ] **Naming Conventions:** camelCase für Funktionen (NICHT snake_case!)
- [ ] **Duplicate Literals:** String-Konstanten statt Wiederholung
- [ ] **Data Types:** Korrekte Typen (int32 vs int64, etc.)

#### Test-Qualität (Details: `UNIVERSAL_TESTING_GUIDELINES.md`):
- [ ] **TDD-Workflow:** Red-Green-Refactor befolgt
- [ ] **Test-Coverage:** Business Logic 90%+ abgedeckt
- [ ] **Test-Independence:** Tests laufen unabhängig voneinander
- [ ] **AAA-Pattern:** Arrange-Act-Assert in allen Tests
- [ ] **SonarQube Test Rules:** Alle struct fields getestet (keine "unused write" warnings)

#### Technische Qualität:
- [ ] **Compiler-Errors:** Alle Fehler behoben
- [ ] **Linter-Warnings:** Alle Warnungen addressiert
- [ ] **Performance:** Kritische Pfade optimiert
- [ ] **Error Handling:** Konsistente Fehlerbehandlung
- [ ] **Documentation:** Komplexe Logik dokumentiert

#### Code-Struktur:
- [ ] **Dependency Injection:** Services über Interfaces
- [ ] **Configuration:** Keine Hard-coded Values
- [ ] **Logging:** Strukturiertes Logging mit angemessenen Levels
- [ ] **Security:** Eingabe-Validierung und sichere Defaults

## 🎯 Universal Best Practices

### Error Handling Patterns:
```go
// ✅ Go Error Handling
func ProcessData(data []byte) (*Result, error) {
    if len(data) == 0 {
        return nil, errors.New("data cannot be empty")
    }
    
    result, err := parseData(data)
    if err != nil {
        return nil, fmt.Errorf("failed to parse data: %w", err)
    }
    
    return result, nil
}
```

```typescript
// ✅ TypeScript Error Handling
async function processData(data: string): Promise<Result> {
    if (!data) {
        throw new Error('Data cannot be empty');
    }
    
    try {
        const parsed = await parseData(data);
        return transformed(parsed);
    } catch (error) {
        throw new Error(`Failed to process data: ${error.message}`);
    }
}
```

```python
# ✅ Python Error Handling
def process_data(data: str) -> Result:
    if not data:
        raise ValueError("Data cannot be empty")
    
    try:
        parsed = parse_data(data)
        return transform(parsed)
    except ParseError as e:
        raise ProcessingError(f"Failed to process data: {e}") from e
```

### Logging Standards:
```javascript
// ✅ Structured Logging
logger.info('User registration completed', {
    userId: user.id,
    email: user.email,
    timestamp: new Date().toISOString(),
    duration: processingTime
});

logger.error('Payment processing failed', {
    orderId: order.id,
    amount: order.total,
    error: error.message,
    stackTrace: error.stack
});
```

### Configuration Management:
```yaml
# ✅ Environment-based Config
development:
  database:
    host: localhost
    port: 5432
  api:
    timeout: 30s
    retries: 3

production:
  database:
    host: ${DB_HOST}
    port: ${DB_PORT}
  api:
    timeout: 10s
    retries: 5
```

---

## 🎯 **Universelle Entwicklungsstandards**

**Diese Guidelines sind ein Überblick und Management-Framework. Die technischen Details sind in den spezialisierten Dokumenten:**

- **`UNIVERSAL_CLEAN_CODE_GUIDELINES.md`** - Vollständige Clean Code und SOLID Prinzipien
- **`UNIVERSAL_TESTING_GUIDELINES.md`** - Umfassende TDD-Patterns und Test-Strategien

**Ziel:** Struktur und Prozess für professionelle Softwareentwicklung - unabhängig von Technologie oder Domain.**
