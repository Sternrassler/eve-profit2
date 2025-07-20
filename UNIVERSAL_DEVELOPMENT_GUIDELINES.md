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
- **Linter:** ESLint (JS/TS), golangci-lint (Go), pylint (Python), SonarQube
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
```
project-root/
├── src/                    # Source Code
│   ├── components/         # UI Components (Frontend)
│   ├── services/          # Business Logic
│   ├── repositories/      # Data Access Layer
│   ├── models/           # Data Models/Entities
│   └── utils/            # Helper Functions
├── tests/                 # Test Files
│   ├── unit/             # Unit Tests (Einzelne Funktionen/Klassen)
│   │   ├── services/     # Tests für src/services/
│   │   ├── models/       # Tests für src/models/
│   │   └── utils/        # Tests für src/utils/
│   ├── integration/      # Integration Tests (Komponenten-Zusammenspiel)
│   │   ├── api/          # API Integration Tests
│   │   ├── database/     # Database Integration Tests
│   │   └── services/     # Service Integration Tests
│   ├── e2e/              # End-to-End Tests (User Workflows)
│   │   ├── auth/         # Authentication Workflows
│   │   ├── checkout/     # Purchase/Checkout Workflows
│   │   └── registration/ # User Registration Workflows
│   └── fixtures/         # Test Data
│       ├── users.json    # User Test Data
│       ├── products.json # Product Test Data
│       └── orders.json   # Order Test Data
├── docs/                 # Documentation
├── config/               # Configuration Files
└── scripts/              # Build/Deploy Scripts
```

### Naming Conventions
```typescript
// ✅ Consistent Naming Patterns

// Classes: PascalCase
class UserService {}
class PaymentProcessor {}

// Functions/Methods: camelCase
function calculateTotal() {}
function validateUserInput() {}

// Constants: SCREAMING_SNAKE_CASE
const MAX_RETRY_ATTEMPTS = 3;
const DEFAULT_TIMEOUT_MS = 5000;

// Variables: camelCase
const userEmail = "user@example.com";
const orderTotal = 99.99;

// Files: kebab-case or camelCase (consistent per project)
user-service.ts / userService.ts
payment-processor.ts / paymentProcessor.ts
```

## 🧪 Universal Testing-Strategy

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

#### Test-Qualität (Details: `UNIVERSAL_TESTING_GUIDELINES.md`):
- [ ] **TDD-Workflow:** Red-Green-Refactor befolgt
- [ ] **Test-Coverage:** Business Logic 90%+ abgedeckt
- [ ] **Test-Independence:** Tests laufen unabhängig voneinander
- [ ] **AAA-Pattern:** Arrange-Act-Assert in allen Tests

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
