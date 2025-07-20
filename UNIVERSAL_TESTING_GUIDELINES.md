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

---

## 🎯 **TDD ist universell anwendbar**

**Test-Driven Development funktioniert in jeder Programmiersprache und für jeden Projekttyp. Die Prinzipien bleiben gleich:**

1. **🔴 RED:** Test schreiben (fehlschlagend)
2. **🟢 GREEN:** Minimale Implementation für Test-Pass
3. **🔄 REFACTOR:** Code verbessern ohne Tests zu brechen

**Diese Guidelines helfen dabei, wartbaren, gut getesteten Code zu schreiben - unabhängig von Technologie oder Domain.**
