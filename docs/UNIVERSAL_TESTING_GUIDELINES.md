# Universal Testing Guidelines

## 🎯 Universelle Test-Driven Development (TDD) Guidelines

Diese Testing-Richtlinien können für jedes Softwareprojekt und jede Programmiersprache verwendet werden.

## 📖 TDD Fundamentals

### Test-First Entwicklung - Der TDD-Zyklus

```
🔴 RED (Test schreiben):
   1. Schreibe einen fehlschlagenden Test
   2. Test definiert das erwartete Verhalten
   3. Kompiliere-Fehler = failing test

🟢 GREEN (Minimal-Code):
   1. Schreibe minimalen Code für Test-Pass
   2. Keine "schöne" Lösung, nur funktional
   3. Test muss grün werden

🔄 REFACTOR (Verbesserung):
   1. Code-Qualität verbessern
   2. Clean Code Prinzipien anwenden
   3. Tests MÜSSEN grün bleiben!
```

## 🧪 Universelle Test-Patterns

### 1. Arrange-Act-Assert (AAA) Pattern

**Das AAA-Pattern funktioniert in allen Programmiersprachen:**

```python
# Python Example
def test_calculate_discount_should_return_20_percent_of_price():
    # Arrange
    price = 100.0
    discount_rate = 0.20
    expected_discount = 20.0
    calculator = DiscountCalculator()
    
    # Act
    actual_discount = calculator.calculate_discount(price, discount_rate)
    
    # Assert
    assert actual_discount == expected_discount
```

```typescript
// TypeScript Example
describe('DiscountCalculator', () => {
    test('should return 20% discount of price', () => {
        // Arrange
        const price = 100;
        const discountRate = 0.20;
        const expectedDiscount = 20;
        const calculator = new DiscountCalculator();
        
        // Act
        const actualDiscount = calculator.calculateDiscount(price, discountRate);
        
        // Assert
        expect(actualDiscount).toBe(expectedDiscount);
    });
});
```

```go
// Go Example
func TestDiscountCalculator_CalculateDiscount_ShouldReturn20PercentOfPrice(t *testing.T) {
    // Arrange
    price := 100.0
    discountRate := 0.20
    expectedDiscount := 20.0
    calculator := NewDiscountCalculator()
    
    // Act
    actualDiscount := calculator.CalculateDiscount(price, discountRate)
    
    // Assert
    if actualDiscount != expectedDiscount {
        t.Errorf("Expected %f, got %f", expectedDiscount, actualDiscount)
    }
}
```

```csharp
// C# Example
[Test]
public void CalculateDiscount_With20Percent_ShouldReturn20PercentOfPrice()
{
    // Arrange
    decimal price = 100m;
    decimal discountRate = 0.20m;
    decimal expectedDiscount = 20m;
    var calculator = new DiscountCalculator();
    
    // Act
    decimal actualDiscount = calculator.CalculateDiscount(price, discountRate);
    
    // Assert
    Assert.AreEqual(expectedDiscount, actualDiscount);
}
```

```java
// Java Example
@Test
public void calculateDiscount_With20Percent_ShouldReturn20PercentOfPrice() {
    // Arrange
    double price = 100.0;
    double discountRate = 0.20;
    double expectedDiscount = 20.0;
    DiscountCalculator calculator = new DiscountCalculator();
    
    // Act
    double actualDiscount = calculator.calculateDiscount(price, discountRate);
    
    // Assert
    assertEquals(expectedDiscount, actualDiscount, 0.01);
}
```

### 2. Given-When-Then (BDD Style)

```python
def test_order_total_calculation():
    # Given
    order = Order()
    order.add_item(Product("Laptop", 1000.00), quantity=2)
    order.add_item(Product("Mouse", 25.00), quantity=1)
    
    # When
    total = order.calculate_total()
    
    # Then
    expected_total = 2025.00
    assert total == expected_total
```

```javascript
describe('Order Total Calculation', () => {
    test('should calculate correct total for multiple items', () => {
        // Given
        const order = new Order();
        order.addItem(new Product('Laptop', 1000.00), 2);
        order.addItem(new Product('Mouse', 25.00), 1);
        
        // When
        const total = order.calculateTotal();
        
        // Then
        expect(total).toBe(2025.00);
    });
});
```

## 🏗️ Test-Struktur Guidelines

### Test-Organisation (Universell)

#### Projekt-Struktur-Mapping
**📖 Siehe: `UNIVERSAL_DEVELOPMENT_GUIDELINES.md` - Projekt-Organisation**

Tests folgen der gleichen Backend/Frontend-Trennung wie die Projekt-Struktur:

```
project-root/
├── backend/
│   ├── internal/          # Source Code
│   └── tests/             # Backend Tests (spiegelt internal/ wider)
│       ├── unit/          # Unit Tests
│       ├── integration/   # Integration Tests  
│       └── fixtures/      # Test Data
├── frontend/
│   ├── src/               # Source Code
│   └── tests/             # Frontend Tests (spiegelt src/ wider)
│       ├── unit/          # Unit Tests
│       ├── integration/   # Integration Tests
│       ├── e2e/           # End-to-End Tests
│       └── fixtures/      # Test Data
└── docs/                  # Documentation
```

#### Test-Struktur-Prinzipien:
- **Mirror Source Structure:** Test-Verzeichnisse spiegeln Source-Code-Struktur wider
- **Co-location:** Tests befinden sich nahe dem zu testenden Code
- **Separation of Concerns:** Backend- und Frontend-Tests sind getrennt
- **Test Types:** Unit → Integration → E2E (von innen nach außen)

### Test-Naming Conventions

#### Englische Convention (Empfohlen für internationale Teams):
```python
# Pattern: test_[method]_[scenario]_[expected_result]
def test_calculate_discount_with_valid_percentage_should_return_correct_amount()
def test_calculate_discount_with_negative_price_should_raise_error()
def test_calculate_discount_with_zero_percentage_should_return_zero()
```

```typescript
// Pattern: should_[expected_behavior]_when_[scenario]
describe('UserValidator', () => {
    test('should return true when email format is valid', () => {});
    test('should return false when email is missing at symbol', () => {});
    test('should throw error when email is null', () => {});
});
```

#### Deutsche Convention:
```python
def test_rabatt_berechnung_mit_gueltigem_prozentsatz_sollte_korrekten_betrag_zurueckgeben()
def test_rabatt_berechnung_mit_negativem_preis_sollte_fehler_werfen()
```

### Test Data Management

#### Test Fixtures (Language-agnostic):
```json
// fixtures/users.json
{
    "validUser": {
        "email": "test@example.com",
        "firstName": "John",
        "lastName": "Doe",
        "age": 25
    },
    "invalidUser": {
        "email": "invalid-email",
        "firstName": "",
        "lastName": null,
        "age": -5
    }
}
```

#### Factory Pattern für Test Data:
```python
# Python Factory
class UserFactory:
    @staticmethod
    def create_valid_user(**kwargs):
        defaults = {
            'email': 'test@example.com',
            'first_name': 'John',
            'last_name': 'Doe',
            'age': 25
        }
        defaults.update(kwargs)
        return User(**defaults)
    
    @staticmethod
    def create_invalid_user():
        return User(
            email='invalid-email',
            first_name='',
            last_name=None,
            age=-5
        )
```

```typescript
// TypeScript Factory
class UserFactory {
    static createValidUser(overrides: Partial<User> = {}): User {
        const defaults = {
            email: 'test@example.com',
            firstName: 'John',
            lastName: 'Doe',
            age: 25
        };
        return new User({ ...defaults, ...overrides });
    }
    
    static createInvalidUser(): User {
        return new User({
            email: 'invalid-email',
            firstName: '',
            lastName: null,
            age: -5
        });
    }
}
```

```go
// Go Factory
type UserFactory struct{}

func (f *UserFactory) CreateValidUser() *User {
    return &User{
        Email:     "test@example.com",
        FirstName: "John",
        LastName:  "Doe",
        Age:       25,
    }
}

func (f *UserFactory) CreateValidUserWithEmail(email string) *User {
    user := f.CreateValidUser()
    user.Email = email
    return user
}

func (f *UserFactory) CreateInvalidUser() *User {
    return &User{
        Email:     "invalid-email",
        FirstName: "",
        LastName:  "",
        Age:       -5,
    }
}
```

## 🎭 Test Types & Strategies

### 1. Unit Tests - Isolierte Komponenten

#### Testing Pure Functions:
```python
# Pure Function - Einfach zu testen
def calculate_tax(amount: float, rate: float) -> float:
    return amount * rate

def test_calculate_tax_with_10_percent_rate():
    # Arrange
    amount = 100.0
    rate = 0.10
    
    # Act
    tax = calculate_tax(amount, rate)
    
    # Assert
    assert tax == 10.0
```

#### Testing Classes with Dependencies (Mocking):
```typescript
// Service mit Dependencies
class OrderService {
    constructor(
        private paymentGateway: PaymentGateway,
        private emailService: EmailService
    ) {}
    
    async processOrder(order: Order): Promise<void> {
        await this.paymentGateway.charge(order.total);
        await this.emailService.sendConfirmation(order.customerEmail);
    }
}

// Test mit Mocks
describe('OrderService', () => {
    test('should charge payment and send email when processing order', async () => {
        // Arrange
        const mockPaymentGateway = {
            charge: jest.fn().mockResolvedValue(undefined)
        };
        const mockEmailService = {
            sendConfirmation: jest.fn().mockResolvedValue(undefined)
        };
        
        const service = new OrderService(mockPaymentGateway, mockEmailService);
        const order = new Order('customer@example.com', 100.00);
        
        // Act
        await service.processOrder(order);
        
        // Assert
        expect(mockPaymentGateway.charge).toHaveBeenCalledWith(100.00);
        expect(mockEmailService.sendConfirmation).toHaveBeenCalledWith('customer@example.com');
    });
});
```

```go
// Go Interface und Mock
type PaymentGateway interface {
    Charge(amount float64) error
}

type OrderService struct {
    paymentGateway PaymentGateway
}

func (s *OrderService) ProcessOrder(order *Order) error {
    return s.paymentGateway.Charge(order.Total)
}

// Mock Implementation
type MockPaymentGateway struct {
    ChargeFunc func(amount float64) error
    CallLog   []float64
}

func (m *MockPaymentGateway) Charge(amount float64) error {
    m.CallLog = append(m.CallLog, amount)
    if m.ChargeFunc != nil {
        return m.ChargeFunc(amount)
    }
    return nil
}

// Test
func TestOrderService_ProcessOrder_ShouldChargeCorrectAmount(t *testing.T) {
    // Arrange
    mockGateway := &MockPaymentGateway{}
    service := &OrderService{paymentGateway: mockGateway}
    order := &Order{Total: 100.0}
    
    // Act
    err := service.ProcessOrder(order)
    
    // Assert
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if len(mockGateway.CallLog) != 1 {
        t.Errorf("Expected 1 charge call, got %d", len(mockGateway.CallLog))
    }
    if mockGateway.CallLog[0] != 100.0 {
        t.Errorf("Expected charge amount 100.0, got %f", mockGateway.CallLog[0])
    }
}
```

### 2. Integration Tests - Komponenten-Zusammenspiel

```python
# Database Integration Test
def test_user_repository_save_and_find():
    # Arrange
    db = create_test_database()  # Test DB
    repository = UserRepository(db)
    user = User(email="test@example.com", name="John")
    
    # Act
    saved_user = repository.save(user)
    found_user = repository.find_by_email("test@example.com")
    
    # Assert
    assert saved_user.id is not None
    assert found_user.email == "test@example.com"
    assert found_user.name == "John"
    
    # Cleanup
    db.cleanup()
```

```typescript
// API Integration Test
describe('User API Integration', () => {
    test('POST /users should create user and return 201', async () => {
        // Arrange
        const userData = {
            email: 'test@example.com',
            firstName: 'John',
            lastName: 'Doe'
        };
        
        // Act
        const response = await request(app)
            .post('/users')
            .send(userData)
            .expect(201);
        
        // Assert
        expect(response.body.id).toBeDefined();
        expect(response.body.email).toBe('test@example.com');
        
        // Verify in database
        const savedUser = await userRepository.findById(response.body.id);
        expect(savedUser).toBeTruthy();
    });
});
```

### 3. End-to-End Tests - User Workflows

```python
# Selenium E2E Test
def test_user_registration_complete_workflow(browser):
    # Arrange
    browser.get("http://localhost:3000/register")
    
    # Act - Fill Registration Form
    browser.find_element(By.ID, "email").send_keys("test@example.com")
    browser.find_element(By.ID, "password").send_keys("securepassword")
    browser.find_element(By.ID, "firstName").send_keys("John")
    browser.find_element(By.ID, "lastName").send_keys("Doe")
    browser.find_element(By.ID, "submitBtn").click()
    
    # Assert - Check Success
    success_message = browser.find_element(By.CLASS_NAME, "success-message")
    assert "Registration successful" in success_message.text
    
    # Verify redirect to dashboard
    assert "/dashboard" in browser.current_url
```

```javascript
// Playwright E2E Test
test('user can complete purchase workflow', async ({ page }) => {
    // Arrange
    await page.goto('/products');
    
    // Act - Add item to cart
    await page.click('[data-testid="product-1-add-button"]');
    await page.click('[data-testid="cart-icon"]');
    
    // Proceed to checkout
    await page.click('[data-testid="checkout-button"]');
    
    // Fill checkout form
    await page.fill('[data-testid="email"]', 'test@example.com');
    await page.fill('[data-testid="credit-card"]', '4111111111111111');
    await page.click('[data-testid="complete-purchase"]');
    
    // Assert
    await expect(page.locator('[data-testid="success-message"]')).toBeVisible();
    await expect(page).toHaveURL(/.*\/order-confirmation/);
});
```

## 🧪 Test-Driven Development Examples

### TDD Workflow - Complete Example

#### Schritt 1: 🔴 RED - Test schreiben
```python
# test_email_validator.py
import pytest
from email_validator import EmailValidator

def test_is_valid_email_with_proper_format_should_return_true():
    # Arrange
    validator = EmailValidator()
    email = "user@example.com"
    
    # Act
    result = validator.is_valid_email(email)
    
    # Assert
    assert result is True

# Dieser Test schlägt fehl, da EmailValidator noch nicht existiert
```

#### Schritt 2: 🟢 GREEN - Minimal-Implementation
```python
# email_validator.py
class EmailValidator:
    def is_valid_email(self, email):
        return "@" in email  # Minimal Implementation für Test-Pass
```

#### Schritt 3: Weiterer Test (🔴 RED)
```python
def test_is_valid_email_without_at_symbol_should_return_false():
    # Arrange
    validator = EmailValidator()
    email = "userexample.com"
    
    # Act
    result = validator.is_valid_email(email)
    
    # Assert
    assert result is False

# Test läuft durch mit aktueller Implementation
```

#### Schritt 4: Mehr Tests (🔴 RED)
```python
def test_is_valid_email_without_domain_should_return_false():
    # Arrange
    validator = EmailValidator()
    email = "user@"
    
    # Act
    result = validator.is_valid_email(email)
    
    # Assert
    assert result is False

# Jetzt schlägt der Test fehl!
```

#### Schritt 5: 🟢 GREEN - Erweiterte Implementation
```python
# email_validator.py
class EmailValidator:
    def is_valid_email(self, email):
        if "@" not in email:
            return False
        
        parts = email.split("@")
        if len(parts) != 2:
            return False
        
        username, domain = parts
        if not username or not domain:
            return False
        
        return True
```

#### Schritt 6: 🔄 REFACTOR - Clean Code
```python
# email_validator.py
import re

class EmailValidator:
    def __init__(self):
        self._email_pattern = re.compile(
            r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
        )
    
    def is_valid_email(self, email: str) -> bool:
        if not email or not isinstance(email, str):
            return False
        
        return bool(self._email_pattern.match(email))
```

#### Vollständige Test-Suite:
```python
# test_email_validator.py
import pytest
from email_validator import EmailValidator

class TestEmailValidator:
    def setup_method(self):
        self.validator = EmailValidator()
    
    def test_is_valid_email_with_proper_format_should_return_true(self):
        # Valid emails
        valid_emails = [
            "user@example.com",
            "test.email@domain.org",
            "user+tag@subdomain.example.co.uk"
        ]
        
        for email in valid_emails:
            assert self.validator.is_valid_email(email), f"Failed for: {email}"
    
    def test_is_valid_email_with_invalid_format_should_return_false(self):
        # Invalid emails
        invalid_emails = [
            "userexample.com",       # No @
            "user@",                 # No domain
            "@domain.com",           # No username
            "user@@domain.com",      # Double @
            "user@domain",           # No TLD
            "",                      # Empty
            None                     # None
        ]
        
        for email in invalid_emails:
            assert not self.validator.is_valid_email(email), f"Should be invalid: {email}"
```

## 🎯 Testing Best Practices (Universal)

### 1. Test Independence
```python
# ❌ Schlecht: Tests hängen voneinander ab
class TestUserService:
    user_id = None
    
    def test_create_user(self):
        user = self.service.create_user("test@example.com")
        self.user_id = user.id  # ⚠️ Seiteneffekt
        assert user.id is not None
    
    def test_get_user(self):
        user = self.service.get_user(self.user_id)  # ⚠️ Abhängigkeit
        assert user is not None

# ✅ Gut: Jeder Test ist unabhängig
class TestUserService:
    def test_create_user_should_return_user_with_id(self):
        # Arrange
        email = "test@example.com"
        
        # Act
        user = self.service.create_user(email)
        
        # Assert
        assert user.id is not None
        assert user.email == email
        
        # Cleanup
        self.service.delete_user(user.id)
    
    def test_get_user_should_return_existing_user(self):
        # Arrange
        existing_user = self.service.create_user("existing@example.com")
        
        # Act
        retrieved_user = self.service.get_user(existing_user.id)
        
        # Assert
        assert retrieved_user is not None
        assert retrieved_user.id == existing_user.id
        
        # Cleanup
        self.service.delete_user(existing_user.id)
```

### 2. Test Isolation (Database Tests)
```python
# ✅ Transaction Rollback Pattern
import pytest
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

@pytest.fixture
def db_session():
    # Create test database connection
    engine = create_engine("sqlite:///:memory:")
    Session = sessionmaker(bind=engine)
    session = Session()
    
    # Start transaction
    transaction = session.begin()
    
    yield session
    
    # Rollback after test
    transaction.rollback()
    session.close()

def test_user_creation_with_db_isolation(db_session):
    # Test uses session, changes are rolled back automatically
    user = User(email="test@example.com")
    db_session.add(user)
    db_session.commit()
    
    assert user.id is not None
    # No cleanup needed - transaction will be rolled back
```

### 3. Mocking Best Practices
```typescript
// ✅ Mock nur externe Dependencies
class UserService {
    constructor(
        private userRepository: UserRepository,      // Mock this
        private emailService: EmailService,         // Mock this
        private logger: Logger                       // Mock this
    ) {}
    
    async registerUser(userData: UserData): Promise<User> {
        // This method should be tested without mocking
        const user = this.validateAndCreateUser(userData);
        await this.userRepository.save(user);
        await this.emailService.sendWelcomeEmail(user.email);
        this.logger.info(`User registered: ${user.id}`);
        return user;
    }
    
    private validateAndCreateUser(userData: UserData): User {
        // Diese private Methode wird nicht gemockt
        // Wird durch Tests der public Methode getestet
    }
}

// Test
describe('UserService', () => {
    test('should register user and send welcome email', async () => {
        // Arrange - Mock nur externe Dependencies
        const mockRepository = { save: jest.fn() };
        const mockEmailService = { sendWelcomeEmail: jest.fn() };
        const mockLogger = { info: jest.fn() };
        
        const service = new UserService(mockRepository, mockEmailService, mockLogger);
        const userData = { email: 'test@example.com', name: 'John' };
        
        // Act
        const result = await service.registerUser(userData);
        
        // Assert
        expect(result.email).toBe('test@example.com');
        expect(mockRepository.save).toHaveBeenCalledWith(expect.any(User));
        expect(mockEmailService.sendWelcomeEmail).toHaveBeenCalledWith('test@example.com');
    });
});
```

### 4. Test Coverage Guidelines
```bash
# Minimum Coverage Targets (Universal)
- Business Logic: 90-100%
- Service Layer: 85-95%
- Controllers/Handlers: 70-85%
- Utilities: 90-100%
- Models/DTOs: 60-80% (meist nur Validierung)

# Coverage Commands (language-specific)
# Python
pytest --cov=src --cov-report=html

# JavaScript/TypeScript
npm test -- --coverage

# Go
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# C#
dotnet test --collect:"XPlat Code Coverage"
```

## 🚨 Universal Testing Checklist

### Vor jedem Commit:

#### TDD-Qualität:
- [ ] **Red-Green-Refactor:** Alle neuen Features mit TDD entwickelt
- [ ] **Test-First:** Tests vor Implementation geschrieben
- [ ] **Test-Namen:** Beschreiben erwartetes Verhalten klar
- [ ] **AAA/Given-When-Then:** Klare Test-Struktur
- [ ] **Test-Isolation:** Tests laufen unabhängig voneinander

#### Test-Coverage:
- [ ] **Unit Tests:** Business Logic 90%+ abgedeckt
- [ ] **Integration Tests:** Kritische Workflows getestet
- [ ] **Edge Cases:** Grenzfälle und Fehlerfälle abgedeckt
- [ ] **Happy Path:** Normale Workflows vollständig getestet

#### Test-Qualität:
- [ ] **Fast Tests:** Unit Tests unter 1ms, Integration Tests unter 100ms
- [ ] **Deterministic:** Tests geben immer gleiche Ergebnisse
- [ ] **Readable:** Test-Intent ist klar verständlich
- [ ] **Maintainable:** Tests sind einfach zu ändern

#### Mocking-Strategy:
- [ ] **External Dependencies:** Nur externe Services gemockt
- [ ] **No Over-Mocking:** Interne Logik nicht gemockt
- [ ] **Clear Interfaces:** Dependencies über Interfaces injiziert
- [ ] **Mock Verification:** Interactions verifiziert wo nötig

## 🎯 Framework-spezifische Patterns

### Python (pytest)
```python
# Fixtures für Setup/Teardown
@pytest.fixture
def user_service():
    service = UserService()
    yield service
    service.cleanup()

# Parametrized Tests
@pytest.mark.parametrize("email,expected", [
    ("test@example.com", True),
    ("invalid-email", False),
    ("", False)
])
def test_email_validation(email, expected):
    assert validate_email(email) == expected

# Async Tests
@pytest.mark.asyncio
async def test_async_operation():
    result = await async_function()
    assert result is not None
```

### JavaScript/TypeScript (Jest)
```typescript
// Setup and Teardown
describe('UserService', () => {
    let service: UserService;
    
    beforeEach(() => {
        service = new UserService();
    });
    
    afterEach(() => {
        service.cleanup();
    });
    
    // Test Suites
    describe('when user is valid', () => {
        test('should create user successfully', () => {
            // Test implementation
        });
    });
    
    // Async Tests
    test('should handle async operations', async () => {
        const result = await service.asyncMethod();
        expect(result).toBeDefined();
    });
});
```

### Go (testing package)
```go
// Table-Driven Tests
func TestEmailValidation(t *testing.T) {
    tests := []struct {
        name     string
        email    string
        expected bool
    }{
        {"valid email", "test@example.com", true},
        {"invalid email", "invalid-email", false},
        {"empty email", "", false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ValidateEmail(tt.email)
            if result != tt.expected {
                t.Errorf("ValidateEmail(%q) = %v, want %v", tt.email, result, tt.expected)
            }
        })
    }
}

// Benchmarks
func BenchmarkEmailValidation(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ValidateEmail("test@example.com")
    }
}
```

### Java (JUnit 5)
```java
// Lifecycle Methods
@ExtendWith(MockitoExtension.class)
class UserServiceTest {
    
    @Mock
    private UserRepository userRepository;
    
    @InjectMocks
    private UserService userService;
    
    @BeforeEach
    void setUp() {
        // Setup before each test
    }
    
    @ParameterizedTest
    @ValueSource(strings = {"test@example.com", "user@domain.org"})
    void shouldValidateValidEmails(String email) {
        assertTrue(EmailValidator.isValid(email));
    }
    
    @Test
    @DisplayName("Should create user when valid data provided")
    void shouldCreateUserWhenValidDataProvided() {
        // Test implementation
    }
}
```

### C# (.NET)
```csharp
[TestFixture]
public class UserServiceTests
{
    private UserService _userService;
    private Mock<IUserRepository> _mockRepository;
    
    [SetUp]
    public void SetUp()
    {
        _mockRepository = new Mock<IUserRepository>();
        _userService = new UserService(_mockRepository.Object);
    }
    
    [TestCase("test@example.com", true)]
    [TestCase("invalid-email", false)]
    [TestCase("", false)]
    public void ValidateEmail_ShouldReturnExpectedResult(string email, bool expected)
    {
        // Act
        var result = _userService.ValidateEmail(email);
        
        // Assert
        Assert.AreEqual(expected, result);
    }
    
    [Test]
    public async Task CreateUserAsync_ShouldReturnUser_WhenValidDataProvided()
    {
        // Arrange
        var userData = new UserData { Email = "test@example.com" };
        
        // Act
        var result = await _userService.CreateUserAsync(userData);
        
        // Assert
        Assert.IsNotNull(result);
        Assert.AreEqual("test@example.com", result.Email);
    }
}
```

## 🎭 End-to-End Testing mit Playwright

### 📚 Playwright E2E Testing Standards

**Projektstruktur-Referenz:** Siehe [UNIVERSAL_DEVELOPMENT_GUIDELINES.md](./UNIVERSAL_DEVELOPMENT_GUIDELINES.md) für die vollständige Projektorganisation.

#### Full-Stack E2E-Verzeichnisstruktur (Root-Level für Backend + Frontend Tests)
```
tests/e2e/                 # Full-Stack E2E Tests Root
├── pages/                 # Page Object Models
│   ├── base-page.ts      # Base Page Class
│   ├── login-page.ts     # Login Page Object
│   ├── dashboard-page.ts # Dashboard Page Object
│   └── api-health-page.ts# API Testing Page Object
├── fixtures/              # Test Data & Setup
│   ├── auth.json         # Authentication State
│   ├── users.json        # User Test Data
│   └── test-data.ts      # Data Factories
├── tests/                 # Test Scenarios
│   ├── auth/             # Authentication Tests
│   ├── workflows/        # Business Workflow Tests
│   └── api/              # API E2E Tests (Backend)
├── utils/                 # Test Utilities
│   ├── auth-helper.ts    # Authentication Helpers
│   └── data-helper.ts    # Data Manipulation
└── config/                # Playwright Configuration
    └── environments.ts    # Environment-specific configs
playwright.config.ts       # Main Playwright Configuration (Root-Level)
```

**Rationale für Root-Level E2E:**
- **Full-Stack Testing:** Tests validieren Backend + Frontend zusammen
- **Einheitliche Konfiguration:** Ein Playwright-Setup für alle E2E-Tests
- **Realitätsnahe Tests:** E2E sollte das komplette System testen
- **Weniger Duplikation:** Ein Setup statt getrennte Backend/Frontend E2E

### 🚀 Playwright Setup & Configuration

#### Installation & Setup
```bash
# Installation
npm init playwright@latest

# Konfiguration für mehrere Browser
npx playwright install
```

#### Playwright Konfiguration (playwright.config.ts)
```typescript
import { defineConfig, devices } from '@playwright/test';

export default defineConfig({
  // Test-Verzeichnis
  testDir: './tests/e2e',
  
  // Parallel Testing
  fullyParallel: true,
  
  // Retry-Strategie
  retries: process.env.CI ? 2 : 0,
  
  // Workers für parallele Ausführung
  workers: process.env.CI ? 1 : undefined,
  
  // Reporter
  reporter: [
    ['html'],
    ['json', { outputFile: 'test-results.json' }],
    ['junit', { outputFile: 'test-results.xml' }]
  ],
  
  // Global Setup
  use: {
    // Base URL
    baseURL: process.env.BASE_URL || 'http://localhost:3000',
    
    // Screenshots bei Fehlern
    screenshot: 'only-on-failure',
    
    // Video bei Fehlern
    video: 'retain-on-failure',
    
    // Trace bei Fehlern
    trace: 'on-first-retry',
    
    // Navigation Timeout
    navigationTimeout: 30000,
    
    // Action Timeout
    actionTimeout: 10000
  },
  
  // Browser-Konfiguration
  projects: [
    {
      name: 'chromium',
      use: { ...devices['Desktop Chrome'] },
    },
    {
      name: 'firefox',
      use: { ...devices['Desktop Firefox'] },
    },
    {
      name: 'webkit',
      use: { ...devices['Desktop Safari'] },
    },
    // Mobile Testing
    {
      name: 'Mobile Chrome',
      use: { ...devices['Pixel 5'] },
    },
    {
      name: 'Mobile Safari',
      use: { ...devices['iPhone 12'] },
    },
  ],
  
  // Dev Server (optional)
  webServer: {
    command: 'npm run dev',
    url: 'http://localhost:3000',
    reuseExistingServer: !process.env.CI,
  },
});
```

### 🏗️ Page Object Model Pattern

#### Base Page Class
```typescript
// pages/base-page.ts
import { Page, Locator } from '@playwright/test';

export abstract class BasePage {
  readonly page: Page;
  
  constructor(page: Page) {
    this.page = page;
  }
  
  // Common navigation methods
  async goto(path: string): Promise<void> {
    await this.page.goto(path);
  }
  
  async waitForPageLoad(): Promise<void> {
    await this.page.waitForLoadState('domcontentloaded');
  }
  
  // Common interaction methods
  async clickElement(selector: string): Promise<void> {
    await this.page.click(selector);
  }
  
  async fillInput(selector: string, text: string): Promise<void> {
    await this.page.fill(selector, text);
  }
  
  async getText(selector: string): Promise<string> {
    return await this.page.textContent(selector) || '';
  }
  
  // Wait helpers
  async waitForElement(selector: string): Promise<Locator> {
    return this.page.waitForSelector(selector);
  }
  
  async waitForUrl(urlPattern: string | RegExp): Promise<void> {
    await this.page.waitForURL(urlPattern);
  }
  
  // Screenshot helper
  async takeScreenshot(name: string): Promise<void> {
    await this.page.screenshot({ 
      path: `test-results/screenshots/${name}.png`,
      fullPage: true 
    });
  }
}
```

#### Specific Page Objects
```typescript
// pages/login-page.ts
import { expect, Page } from '@playwright/test';
import { BasePage } from './base-page';

export class LoginPage extends BasePage {
  // Locators
  readonly emailInput = this.page.locator('[data-testid="email-input"]');
  readonly passwordInput = this.page.locator('[data-testid="password-input"]');
  readonly loginButton = this.page.locator('[data-testid="login-button"]');
  readonly errorMessage = this.page.locator('[data-testid="error-message"]');
  readonly forgotPasswordLink = this.page.locator('[data-testid="forgot-password"]');
  
  constructor(page: Page) {
    super(page);
  }
  
  // Navigation
  async navigate(): Promise<void> {
    await this.goto('/login');
    await this.waitForPageLoad();
  }
  
  // Actions
  async login(email: string, password: string): Promise<void> {
    await this.emailInput.fill(email);
    await this.passwordInput.fill(password);
    await this.loginButton.click();
  }
  
  async loginWithValidCredentials(): Promise<void> {
    await this.login('test@example.com', 'password123');
  }
  
  async clickForgotPassword(): Promise<void> {
    await this.forgotPasswordLink.click();
  }
  
  // Assertions
  async expectLoginSuccess(): Promise<void> {
    await expect(this.page).toHaveURL(/.*\/dashboard/);
  }
  
  async expectLoginError(expectedMessage: string): Promise<void> {
    await expect(this.errorMessage).toBeVisible();
    await expect(this.errorMessage).toContainText(expectedMessage);
  }
  
  async expectToBeOnLoginPage(): Promise<void> {
    await expect(this.page).toHaveURL(/.*\/login/);
    await expect(this.loginButton).toBeVisible();
  }
}
```

```typescript
// pages/dashboard-page.ts
import { expect, Page, Locator } from '@playwright/test';
import { BasePage } from './base-page';

export class DashboardPage extends BasePage {
  // Locators
  readonly welcomeMessage = this.page.locator('[data-testid="welcome-message"]');
  readonly userMenu = this.page.locator('[data-testid="user-menu"]');
  readonly logoutButton = this.page.locator('[data-testid="logout-button"]');
  readonly statsCards = this.page.locator('[data-testid="stats-card"]');
  readonly navigationMenu = this.page.locator('[data-testid="nav-menu"]');
  
  constructor(page: Page) {
    super(page);
  }
  
  // Navigation
  async navigate(): Promise<void> {
    await this.goto('/dashboard');
    await this.waitForPageLoad();
  }
  
  // Actions
  async logout(): Promise<void> {
    await this.userMenu.click();
    await this.logoutButton.click();
  }
  
  async navigateToSection(sectionName: string): Promise<void> {
    await this.navigationMenu.locator(`text=${sectionName}`).click();
  }
  
  // Getters
  async getWelcomeMessage(): Promise<string> {
    return await this.welcomeMessage.textContent() || '';
  }
  
  async getStatsCount(): Promise<number> {
    return await this.statsCards.count();
  }
  
  // Assertions
  async expectDashboardLoaded(): Promise<void> {
    await expect(this.welcomeMessage).toBeVisible();
    await expect(this.navigationMenu).toBeVisible();
  }
  
  async expectUserLoggedIn(userName: string): Promise<void> {
    await expect(this.welcomeMessage).toContainText(userName);
  }
}
```

### 🧪 E2E Test Examples

#### Authentication Tests
```typescript
// tests/auth/login.spec.ts
import { test, expect } from '@playwright/test';
import { LoginPage } from '../../pages/login-page';
import { DashboardPage } from '../../pages/dashboard-page';

test.describe('User Authentication', () => {
  let loginPage: LoginPage;
  let dashboardPage: DashboardPage;
  
  test.beforeEach(async ({ page }) => {
    loginPage = new LoginPage(page);
    dashboardPage = new DashboardPage(page);
    await loginPage.navigate();
  });
  
  test('should login with valid credentials', async ({ page }) => {
    // Arrange
    const validEmail = 'test@example.com';
    const validPassword = 'password123';
    
    // Act
    await loginPage.login(validEmail, validPassword);
    
    // Assert
    await loginPage.expectLoginSuccess();
    await dashboardPage.expectDashboardLoaded();
    await dashboardPage.expectUserLoggedIn('Test User');
  });
  
  test('should show error with invalid credentials', async ({ page }) => {
    // Arrange
    const invalidEmail = 'invalid@example.com';
    const invalidPassword = 'wrongpassword';
    
    // Act
    await loginPage.login(invalidEmail, invalidPassword);
    
    // Assert
    await loginPage.expectLoginError('Invalid credentials');
    await loginPage.expectToBeOnLoginPage();
  });
  
  test('should navigate to forgot password', async ({ page }) => {
    // Act
    await loginPage.clickForgotPassword();
    
    // Assert
    await expect(page).toHaveURL(/.*\/forgot-password/);
  });
});
```

#### Business Workflow Tests
```typescript
// tests/workflows/user-registration.spec.ts
import { test, expect } from '@playwright/test';
import { RegistrationPage } from '../../pages/registration-page';
import { DashboardPage } from '../../pages/dashboard-page';
import { generateUniqueUser } from '../../fixtures/test-data';

test.describe('User Registration Workflow', () => {
  test('complete user registration flow', async ({ page }) => {
    // Arrange
    const registrationPage = new RegistrationPage(page);
    const dashboardPage = new DashboardPage(page);
    const userData = generateUniqueUser();
    
    // Act & Assert - Step by step workflow
    
    // Step 1: Navigate to registration
    await registrationPage.navigate();
    await registrationPage.expectToBeOnRegistrationPage();
    
    // Step 2: Fill registration form
    await registrationPage.fillRegistrationForm(userData);
    await registrationPage.submitForm();
    
    // Step 3: Verify email verification page
    await expect(page).toHaveURL(/.*\/verify-email/);
    
    // Step 4: Simulate email verification (in real test, check email)
    await registrationPage.verifyEmail(userData.email);
    
    // Step 5: Complete profile setup
    await registrationPage.completeProfileSetup(userData.profile);
    
    // Step 6: Verify successful registration and dashboard access
    await dashboardPage.expectDashboardLoaded();
    await dashboardPage.expectUserLoggedIn(userData.firstName);
    
    // Step 7: Verify user data is correctly displayed
    await registrationPage.verifyUserProfile(userData);
  });
});
```

### 🔧 Advanced Playwright Patterns

#### API Testing mit Playwright
```typescript
// tests/api/user-api.spec.ts
import { test, expect } from '@playwright/test';

test.describe('User API E2E Tests', () => {
  test('should create user via API and verify in UI', async ({ request, page }) => {
    // Arrange
    const userData = {
      email: 'api-test@example.com',
      firstName: 'API',
      lastName: 'Test',
      password: 'password123'
    };
    
    // Act - Create user via API
    const response = await request.post('/api/users', {
      data: userData
    });
    
    // Assert API response
    expect(response.status()).toBe(201);
    const createdUser = await response.json();
    expect(createdUser.email).toBe(userData.email);
    
    // Act - Verify user in UI
    await page.goto('/admin/users');
    const userRow = page.locator(`[data-testid="user-${createdUser.id}"]`);
    
    // Assert UI shows user
    await expect(userRow).toBeVisible();
    await expect(userRow.locator('.email')).toContainText(userData.email);
  });
  
  test('should handle API errors gracefully in UI', async ({ request, page }) => {
    // Arrange - Create user with duplicate email
    const duplicateUserData = {
      email: 'existing@example.com', // Already exists
      firstName: 'Duplicate',
      lastName: 'User'
    };
    
    // Act - Try to register via UI
    await page.goto('/register');
    await page.fill('[data-testid="email"]', duplicateUserData.email);
    await page.fill('[data-testid="firstName"]', duplicateUserData.firstName);
    await page.fill('[data-testid="lastName"]', duplicateUserData.lastName);
    await page.click('[data-testid="submit-button"]');
    
    // Assert - UI shows appropriate error
    const errorMessage = page.locator('[data-testid="error-message"]');
    await expect(errorMessage).toBeVisible();
    await expect(errorMessage).toContainText('Email already exists');
  });
});
```

#### Visual Testing
```typescript
// tests/visual/homepage.spec.ts
import { test, expect } from '@playwright/test';

test.describe('Visual Regression Tests', () => {
  test('homepage should match visual baseline', async ({ page }) => {
    // Navigate to homepage
    await page.goto('/');
    
    // Wait for page to fully load
    await page.waitForLoadState('networkidle');
    
    // Hide dynamic elements (timestamps, etc.)
    await page.addStyleTag({
      content: `
        .timestamp, .dynamic-content {
          visibility: hidden !important;
        }
      `
    });
    
    // Take screenshot and compare
    await expect(page).toHaveScreenshot('homepage.png');
  });
  
  test('mobile homepage should match baseline', async ({ page }) => {
    // Set mobile viewport
    await page.setViewportSize({ width: 375, height: 667 });
    
    await page.goto('/');
    await page.waitForLoadState('networkidle');
    
    // Mobile-specific screenshot
    await expect(page).toHaveScreenshot('homepage-mobile.png');
  });
});
```

#### Performance Testing
```typescript
// tests/performance/page-performance.spec.ts
import { test, expect } from '@playwright/test';

test.describe('Performance Tests', () => {
  test('homepage should load within performance budget', async ({ page }) => {
    // Start performance monitoring
    await page.goto('/', { waitUntil: 'networkidle' });
    
    // Get performance metrics
    const performanceMetrics = await page.evaluate(() => {
      const navigation = performance.getEntriesByType('navigation')[0] as PerformanceNavigationTiming;
      return {
        domContentLoaded: navigation.domContentLoadedEventEnd - navigation.domContentLoadedEventStart,
        firstPaint: performance.getEntriesByName('first-paint')[0]?.startTime,
        firstContentfulPaint: performance.getEntriesByName('first-contentful-paint')[0]?.startTime,
        loadComplete: navigation.loadEventEnd - navigation.loadEventStart
      };
    });
    
    // Assert performance budgets
    expect(performanceMetrics.domContentLoaded).toBeLessThan(1000); // < 1s
    expect(performanceMetrics.firstContentfulPaint).toBeLessThan(1500); // < 1.5s
    expect(performanceMetrics.loadComplete).toBeLessThan(3000); // < 3s
  });
});
```

### 🔄 Test Data Management

#### Test Data Factory
```typescript
// fixtures/test-data.ts
import { faker } from '@faker-js/faker';

export interface UserData {
  email: string;
  firstName: string;
  lastName: string;
  password: string;
  profile: {
    bio: string;
    company: string;
    location: string;
  };
}

export function generateUniqueUser(): UserData {
  return {
    email: faker.internet.email(),
    firstName: faker.person.firstName(),
    lastName: faker.person.lastName(),
    password: 'Test123!@#',
    profile: {
      bio: faker.lorem.paragraph(),
      company: faker.company.name(),
      location: faker.location.city()
    }
  };
}

export function generateValidUser(overrides: Partial<UserData> = {}): UserData {
  const defaultUser = {
    email: 'test@example.com',
    firstName: 'John',
    lastName: 'Doe',
    password: 'Test123!@#',
    profile: {
      bio: 'Test user bio',
      company: 'Test Company',
      location: 'Test City'
    }
  };
  
  return { ...defaultUser, ...overrides };
}

// Test Data für verschiedene Szenarien
export const testUsers = {
  admin: generateValidUser({
    email: 'admin@example.com',
    firstName: 'Admin',
    lastName: 'User'
  }),
  
  standardUser: generateValidUser({
    email: 'user@example.com',
    firstName: 'Standard',
    lastName: 'User'
  }),
  
  premiumUser: generateValidUser({
    email: 'premium@example.com',
    firstName: 'Premium',
    lastName: 'User'
  })
};
```

#### Database State Management
```typescript
// fixtures/database-helper.ts
import { test as base } from '@playwright/test';
import { PrismaClient } from '@prisma/client';

const prisma = new PrismaClient();

export const test = base.extend({
  // Database cleanup fixture
  cleanDatabase: async ({}, use) => {
    // Cleanup before test
    await prisma.user.deleteMany();
    await prisma.order.deleteMany();
    
    await use();
    
    // Cleanup after test
    await prisma.user.deleteMany();
    await prisma.order.deleteMany();
  },
  
  // Pre-seeded data fixture
  seededData: async ({}, use) => {
    // Create test data
    const user = await prisma.user.create({
      data: {
        email: 'test@example.com',
        firstName: 'Test',
        lastName: 'User'
      }
    });
    
    await use({ user });
    
    // Cleanup
    await prisma.user.delete({ where: { id: user.id } });
  }
});
```

### 🐛 Debugging & Troubleshooting

#### Debug-Modi
```typescript
// Debug-spezifische Test-Konfiguration
test.describe.configure({ mode: 'serial' }); // Tests sequenziell ausführen

test('debug specific scenario', async ({ page }) => {
  // Slow down actions for debugging
  test.slow();
  
  // Enable debugging
  await page.pause(); // Pausiert für manuelle Inspektion
  
  // Step-by-step debugging
  await page.goto('/');
  console.log('Navigated to homepage');
  
  await page.click('[data-testid="button"]');
  console.log('Clicked button');
  
  // Screenshot bei jedem Schritt
  await page.screenshot({ path: 'debug-step-1.png' });
});
```

#### Error Handling
```typescript
// Robuste Error-Behandlung
test('resilient test with retries', async ({ page }) => {
  // Retry-Logik für flaky elements
  await test.step('Navigate and wait for element', async () => {
    await page.goto('/');
    
    // Warte auf kritisches Element mit Retry
    const criticalElement = page.locator('[data-testid="critical-element"]');
    await expect(criticalElement).toBeVisible({ timeout: 10000 });
  });
  
  // Graceful handling von langsamen Operationen
  await test.step('Perform slow operation', async () => {
    await page.click('[data-testid="slow-action"]');
    
    // Erwarte Loading-State
    await expect(page.locator('[data-testid="loading"]')).toBeVisible();
    
    // Warte auf Completion
    await expect(page.locator('[data-testid="loading"]')).toBeHidden({ timeout: 30000 });
    await expect(page.locator('[data-testid="result"]')).toBeVisible();
  });
});
```

### 🚀 CI/CD Integration

#### GitHub Actions Konfiguration
```yaml
# .github/workflows/e2e-tests.yml
name: E2E Tests

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  e2e-tests:
    runs-on: ubuntu-latest
    
    strategy:
      matrix:
        browser: [chromium, firefox, webkit]
        
    steps:
    - uses: actions/checkout@v3
    
    - uses: actions/setup-node@v3
      with:
        node-version: '18'
        cache: 'npm'
    
    - name: Install dependencies
      run: npm ci
    
    - name: Install Playwright browsers
      run: npx playwright install --with-deps ${{ matrix.browser }}
    
    - name: Start application
      run: |
        npm run build
        npm run start &
        npx wait-on http://localhost:3000
    
    - name: Run E2E tests
      run: npx playwright test --project=${{ matrix.browser }}
      env:
        CI: true
        BASE_URL: http://localhost:3000
    
    - name: Upload test results
      if: failure()
      uses: actions/upload-artifact@v3
      with:
        name: playwright-report-${{ matrix.browser }}
        path: playwright-report/
```

### 📊 E2E Testing Best Practices

#### 1. Test-Isolation & Independence
```typescript
// ✅ Jeder Test ist unabhängig
test.describe('Independent Tests', () => {
  test.beforeEach(async ({ page }) => {
    // Fresh state für jeden Test
    await page.goto('/');
    await page.evaluate(() => localStorage.clear());
    await page.evaluate(() => sessionStorage.clear());
  });
  
  test('test A should not affect test B', async ({ page }) => {
    // Test A logic
  });
  
  test('test B runs independently', async ({ page }) => {
    // Test B logic - unabhängig von Test A
  });
});
```

#### 2. Stabile Selektoren
```typescript
// ✅ Verwende data-testid für stabile Selektoren
await page.click('[data-testid="submit-button"]');

// ❌ Vermeide fragile Selektoren
await page.click('button.btn.btn-primary'); // Kann sich ändern
await page.click('text=Submit'); // Language-dependent
```

#### 3. Page Object Model Benefits
```typescript
// ✅ Page Objects abstrahieren UI-Details
const loginPage = new LoginPage(page);
await loginPage.loginWithValidCredentials();
await loginPage.expectLoginSuccess();

// ❌ Direkte UI-Manipulation in Tests
await page.fill('#email', 'test@example.com');
await page.fill('#password', 'password');
await page.click('#login-button');
```

#### 4. Assertions-Strategy
```typescript
// ✅ Klare, erwartete Zustände testen
await expect(page.locator('[data-testid="success-message"]')).toBeVisible();
await expect(page).toHaveURL(/.*\/dashboard/);
await expect(page.locator('[data-testid="user-name"]')).toContainText('John Doe');

// ✅ Negative Assertions für Error-Cases
await expect(page.locator('[data-testid="error-message"]')).toBeVisible();
await expect(page.locator('[data-testid="loading"]')).toBeHidden();
```

### 🎯 EVE Profit Calculator Specific E2E Scenarios

#### EVE SSO Authentication Flow
```typescript
// tests/eve-auth/sso-flow.spec.ts
test('EVE SSO authentication workflow', async ({ page }) => {
  // Navigate to login
  await page.goto('/login');
  
  // Click EVE SSO login
  await page.click('[data-testid="eve-sso-login"]');
  
  // Should redirect to EVE SSO
  await expect(page).toHaveURL(/.*login\.eveonline\.com.*/);
  
  // Fill EVE credentials (test environment)
  await page.fill('[name="UserName"]', process.env.EVE_TEST_USERNAME!);
  await page.fill('[name="Password"]', process.env.EVE_TEST_PASSWORD!);
  await page.click('[type="submit"]');
  
  // Grant permissions
  await page.click('[value="Authorize"]');
  
  // Should redirect back to app
  await expect(page).toHaveURL(/.*localhost.*\/dashboard/);
  
  // Verify character data is loaded
  await expect(page.locator('[data-testid="character-name"]')).toBeVisible();
});
```

#### Market Data E2E Tests
```typescript
// tests/market/profit-calculation.spec.ts
test('profit calculation with real market data', async ({ page }) => {
  // Login first
  await loginAsTestCharacter(page);
  
  // Navigate to profit calculator
  await page.goto('/profit-calculator');
  
  // Search for item
  await page.fill('[data-testid="item-search"]', 'Tritanium');
  await page.click('[data-testid="search-button"]');
  
  // Select item from results
  await page.click('[data-testid="item-34"]'); // Tritanium type_id
  
  // Verify market data is loaded
  await expect(page.locator('[data-testid="buy-price"]')).toBeVisible();
  await expect(page.locator('[data-testid="sell-price"]')).toBeVisible();
  
  // Set quantity
  await page.fill('[data-testid="quantity-input"]', '1000000');
  
  // Calculate profit
  await page.click('[data-testid="calculate-button"]');
  
  // Verify results
  await expect(page.locator('[data-testid="profit-result"]')).toBeVisible();
  await expect(page.locator('[data-testid="margin-percentage"]')).toBeVisible();
});
```

---

## 🎯 **TDD ist universell anwendbar**

**Test-Driven Development funktioniert in jeder Programmiersprache und für jeden Projekttyp. Die Prinzipien bleiben gleich:**

1. **🔴 RED:** Test schreiben (fehlschlagend)
2. **🟢 GREEN:** Minimale Implementation für Test-Pass
3. **🔄 REFACTOR:** Code verbessern ohne Tests zu brechen

**Diese Guidelines helfen dabei, wartbaren, gut getesteten Code zu schreiben - unabhängig von Technologie oder Domain.**
