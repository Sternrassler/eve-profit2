# Universal Clean Code Reference - EVE Profit Calculator 2.0

> **Hinweis:** Primäre Coding Rules sind in `.github/copilot-instructions.md` definiert.  
> Diese Datei enthält erweiterte Clean Code Prinzipien und detaillierte Implementierungsrichtlinien.

## 📚 Clean Code Prinzipien (Robert C. Martin)

Diese erweiterte Referenz ergänzt die Basis-Regeln aus `.github/copilot-instructions.md` mit detaillierten Implementierungsrichtlinien für Go und TypeScript.

## 🎯 Erweiterte Clean Code Prinzipien

> **Basis-Regeln:** Siehe `.github/copilot-instructions.md`  
> **Meaningful Names:** `calculateNetProfit()` statt `calc()`  
> **Single Responsibility:** Eine Funktion = eine Aufgabe  
> **Dependency Injection:** Interfaces für Testbarkeit  

### 1. Advanced Function Design

### 1. Advanced Function Design

**EVE-spezifische Implementierung:**

```go
// ✅ EVE Domain-spezifische Namen mit Business Logic
func CalculateTradingProfitWithCharacterSkills(
    buyPrice, sellPrice float64, 
    quantity int, 
    characterSkills models.CharacterSkills,
) (*models.TradingProfit, error) {
    if buyPrice <= 0 || sellPrice <= 0 {
        return nil, errors.New("invalid prices")
    }
    
    baseProfitPerUnit := sellPrice - buyPrice
    skillMultiplier := characterSkills.GetTradingMultiplier()
    adjustedProfit := baseProfitPerUnit * skillMultiplier
    
    return &models.TradingProfit{
        NetProfit: adjustedProfit * float64(quantity),
        Margin:    (adjustedProfit / buyPrice) * 100,
        Volume:    quantity,
    }, nil
}
```

**Erweiterte Regeln:**

```go
// ❌ Schlecht: Zu viele Verantwortlichkeiten
func ProcessUserOrder(data []byte) error {
    // Validierung
    if len(data) == 0 { return errors.New("no data") }
    
    // Parsing
    var order Order
    json.Unmarshal(data, &order)
    
    // Business Logic
    order.Total = order.Price * order.Quantity
    
    // Persistierung
    database.Save(order)
    
    // Logging
    log.Printf("Processed order %d", order.ID)
    
    return nil
}

// ✅ Gut: Single Responsibility
func ValidateOrderData(data []byte) error {
    if len(data) == 0 {
        return errors.New("order data cannot be empty")
    }
    return nil
}

func ParseOrder(data []byte) (*Order, error) {
    var order Order
    if err := json.Unmarshal(data, &order); err != nil {
        return nil, fmt.Errorf("failed to parse order: %w", err)
    }
    return &order, nil
}

func CalculateOrderTotal(order *Order) {
    order.Total = order.Price * order.Quantity
}

func SaveOrder(order *Order) error {
    return database.Save(order)
}

func LogOrderProcessed(orderID int) {
    log.Printf("Successfully processed order %d", orderID)
}
```

```typescript
// ❌ Schlecht: Mehrere Verantwortlichkeiten
async function processUser(userData: any): Promise<void> {
    // Validierung, Transformation, API Call, Caching, Logging - alles in einer Funktion
}

// ✅ Gut: Getrennte Verantwortlichkeiten
function validateUserData(userData: unknown): UserData {
    // Nur Validierung
}

function transformUserData(userData: UserData): ApiUserData {
    // Nur Transformation
}

async function saveUser(userData: ApiUserData): Promise<User> {
    // Nur API Call
}

function cacheUser(user: User): void {
    // Nur Caching
}

function logUserCreated(userId: string): void {
    // Nur Logging
}
```

**Regeln:**
- Funktionen sollten klein sein (< 20 Zeilen)
- Funktionen sollten eine Sache tun
- Ein Abstraktionslevel pro Funktion
- Verwende deskriptive Namen

### 3. Comments (Kommentare)

```go
// ❌ Schlecht: Kommentare erklären WAS
func CalculateTax(amount float64) float64 {
    // Multiply amount by 0.1 to get 10% tax
    return amount * 0.1
}

// ✅ Gut: Kommentare erklären WARUM (wenn überhaupt nötig)
func CalculateVAT(amount float64) float64 {
    // European VAT rate varies by country, but 19% is common standard
    // This rate should be configurable in production systems
    const STANDARD_VAT_RATE = 0.19
    return amount * STANDARD_VAT_RATE
}

// ✅ Besser: Selbstdokumentierender Code
const STANDARD_VAT_RATE = 0.19

func CalculateVAT(amount float64) float64 {
    return amount * STANDARD_VAT_RATE
}
```

```typescript
// ❌ Schlecht: Redundante Kommentare
function getUserById(id: string): User {
    // Get user by ID from database
    return database.users.find(user => user.id === id);
}

// ✅ Gut: Kommentare für komplexe Business Logic
function calculateCompoundInterest(principal: number, rate: number, time: number): number {
    // Using compound interest formula: A = P(1 + r/n)^(nt)
    // Assumes annual compounding (n=1) for simplicity
    return principal * Math.pow(1 + rate, time);
}

// ✅ Besser: Selbstdokumentierender Code
const ANNUAL_COMPOUNDING_PERIODS = 1;

function calculateCompoundInterest(
    principal: number, 
    annualInterestRate: number, 
    yearsInvested: number
): number {
    return principal * Math.pow(1 + annualInterestRate, yearsInvested);
}
```

**Regeln:**
- Kommentare sind ein Versagen des Codes, sich selbst zu erklären
- Erkläre WARUM, nicht WAS
- Halte Kommentare aktuell oder lösche sie

### 4. Formatting (Formatierung)

```go
// ✅ Konsistente Formatierung
type UserService struct {
    repository UserRepositoryInterface
    cache      CacheInterface
    logger     LoggerInterface
}

func NewUserService(
    repository UserRepositoryInterface,
    cache CacheInterface,
    logger LoggerInterface,
) *UserService {
    return &UserService{
        repository: repository,
        cache:      cache,
        logger:     logger,
    }
}

func (s *UserService) GetActiveUsers(
    department string,
    limit int,
) ([]User, error) {
    if limit <= 0 {
        return nil, errors.New("limit must be positive")
    }
    
    users := make([]User, 0, limit)
    
    allUsers, err := s.repository.GetUsersByDepartment(department)
    if err != nil {
        s.logger.Error("Failed to get users", 
            "department", department, 
            "error", err)
        return nil, err
    }
    
    for _, user := range allUsers {
        if user.IsActive && len(users) < limit {
            users = append(users, user)
        }
    }
    
    return users, nil
}
```

```typescript
// ✅ Konsistente TypeScript Formatierung
interface UserServiceDependencies {
    repository: UserRepository;
    cache: CacheService;
    logger: Logger;
}

class UserService {
    constructor(
        private readonly repository: UserRepository,
        private readonly cache: CacheService,
        private readonly logger: Logger
    ) {}

    async getActiveUsers(
        department: string,
        limit: number
    ): Promise<User[]> {
        if (limit <= 0) {
            throw new Error('Limit must be positive');
        }

        const cacheKey = `active-users-${department}-${limit}`;
        const cachedUsers = await this.cache.get<User[]>(cacheKey);
        
        if (cachedUsers) {
            return cachedUsers;
        }

        try {
            const allUsers = await this.repository.getUsersByDepartment(department);
            const activeUsers = allUsers
                .filter(user => user.isActive)
                .slice(0, limit);

            await this.cache.set(cacheKey, activeUsers, { ttl: 300 });
            return activeUsers;

        } catch (error) {
            this.logger.error('Failed to get active users', {
                department,
                limit,
                error: error.message
            });
            throw error;
        }
    }
}
```

## 🏗️ SOLID Prinzipien

### S - Single Responsibility Principle (SRP)

```go
// ✅ Jede Struktur hat eine klare Verantwortlichkeit
type UserDataFetcher struct {
    apiClient APIClientInterface
}

type UserValidator struct {
    rules ValidationRules
}

type UserNotifier struct {
    emailService EmailServiceInterface
    smsService   SMSServiceInterface
}

type UserProcessor struct {
    fetcher   UserDataFetcher
    validator UserValidator
    notifier  UserNotifier
}
```

```typescript
// ✅ Getrennte Verantwortlichkeiten
class PaymentProcessor {
    // Nur Payment Processing
}

class OrderCalculator {
    // Nur Berechnungen
}

class InventoryManager {
    // Nur Inventory Management
}

class NotificationService {
    // Nur Benachrichtigungen
}
```

### O - Open/Closed Principle (OCP)

```go
// ✅ Erweiterbar ohne Modifikation
type PriceCalculator interface {
    CalculatePrice(product Product) float64
}

type StandardPriceCalculator struct{}
func (c *StandardPriceCalculator) CalculatePrice(product Product) float64 {
    return product.BasePrice
}

type DiscountPriceCalculator struct {
    discountRate float64
}
func (c *DiscountPriceCalculator) CalculatePrice(product Product) float64 {
    return product.BasePrice * (1 - c.discountRate)
}

type PremiumPriceCalculator struct {
    premiumRate float64
}
func (c *PremiumPriceCalculator) CalculatePrice(product Product) float64 {
    return product.BasePrice * (1 + c.premiumRate)
}
```

```typescript
// ✅ Erweiterbar durch Interfaces
interface PaymentProcessor {
    processPayment(amount: number): Promise<PaymentResult>;
}

class CreditCardProcessor implements PaymentProcessor {
    async processPayment(amount: number): Promise<PaymentResult> {
        // Credit card logic
    }
}

class PayPalProcessor implements PaymentProcessor {
    async processPayment(amount: number): Promise<PaymentResult> {
        // PayPal logic
    }
}

class CryptoProcessor implements PaymentProcessor {
    async processPayment(amount: number): Promise<PaymentResult> {
        // Cryptocurrency logic
    }
}
```

### L - Liskov Substitution Principle (LSP)

```go
// ✅ Subtypes müssen ersetzbar sein
type Storage interface {
    Save(data interface{}) error
    Load(id string) (interface{}, error)
}

type DatabaseStorage struct{}
func (d *DatabaseStorage) Save(data interface{}) error { /* DB implementation */ }
func (d *DatabaseStorage) Load(id string) (interface{}, error) { /* DB implementation */ }

type FileStorage struct{}
func (f *FileStorage) Save(data interface{}) error { /* File implementation */ }
func (f *FileStorage) Load(id string) (interface{}, error) { /* File implementation */ }

type CacheStorage struct{}
func (c *CacheStorage) Save(data interface{}) error { /* Cache implementation */ }
func (c *CacheStorage) Load(id string) (interface{}, error) { /* Cache implementation */ }

// Alle sind vollständig austauschbar
func ProcessData(storage Storage, data interface{}) error {
    return storage.Save(data) // Funktioniert mit allen Implementierungen
}
```

### I - Interface Segregation Principle (ISP)

```go
// ❌ Schlecht: Zu große Interface
type UserManager interface {
    GetUser(id string) (*User, error)
    SaveUser(user *User) error
    SendEmail(email string) error
    LogActivity(activity string) error
    ValidateUser(user *User) error
}

// ✅ Gut: Getrennte, spezifische Interfaces
type UserRepository interface {
    GetUser(id string) (*User, error)
    SaveUser(user *User) error
}

type EmailSender interface {
    SendEmail(email string) error
}

type ActivityLogger interface {
    LogActivity(activity string) error
}

type UserValidator interface {
    ValidateUser(user *User) error
}
```

```typescript
// ❌ Schlecht: Fettes Interface
interface UserService {
    getUser(id: string): Promise<User>;
    saveUser(user: User): Promise<void>;
    sendEmail(email: string): Promise<void>;
    logActivity(activity: string): void;
    validateUser(user: User): boolean;
    generateReport(): Promise<Report>;
    backupData(): Promise<void>;
}

// ✅ Gut: Getrennte Interfaces
interface UserRepository {
    getUser(id: string): Promise<User>;
    saveUser(user: User): Promise<void>;
}

interface EmailService {
    sendEmail(email: string): Promise<void>;
}

interface ActivityLogger {
    logActivity(activity: string): void;
}

interface UserValidator {
    validateUser(user: User): boolean;
}

interface ReportGenerator {
    generateReport(): Promise<Report>;
}

interface DataBackupService {
    backupData(): Promise<void>;
}
```

### D - Dependency Inversion Principle (DIP)

```go
// ✅ Abhängigkeiten über Interfaces
type OrderService struct {
    repository    OrderRepositoryInterface    // Interface, nicht konkrete Implementierung
    calculator    PriceCalculatorInterface   // Interface, nicht konkrete Implementierung
    notifier      NotificationInterface      // Interface, nicht konkrete Implementierung
    logger        LoggerInterface            // Interface, nicht konkrete Implementierung
}

func NewOrderService(
    repository OrderRepositoryInterface,
    calculator PriceCalculatorInterface,
    notifier NotificationInterface,
    logger LoggerInterface,
) *OrderService {
    return &OrderService{
        repository: repository,
        calculator: calculator,
        notifier:   notifier,
        logger:     logger,
    }
}
```

```typescript
// ✅ Dependency Injection mit Interfaces
class OrderService {
    constructor(
        private readonly repository: OrderRepository,      // Interface
        private readonly calculator: PriceCalculator,      // Interface
        private readonly notifier: NotificationService,    // Interface
        private readonly logger: Logger                    // Interface
    ) {}

    async processOrder(order: Order): Promise<void> {
        const calculatedOrder = await this.calculator.calculateTotal(order);
        await this.repository.save(calculatedOrder);
        await this.notifier.sendConfirmation(order.customerId);
        this.logger.info('Order processed', { orderId: order.id });
    }
}
```

## 📋 Universal Clean Code Checkliste

### Funktionen:
- [ ] Funktion hat eine einzige Verantwortlichkeit
- [ ] Funktionsname beschreibt genau was sie tut
- [ ] Maximal 20 Zeilen
- [ ] Maximal 3-4 Parameter
- [ ] Keine Side Effects

### Variablen:
- [ ] Aussagekräftige Namen
- [ ] Keine Abkürzungen oder Codierungen
- [ ] Konstanten für Magic Numbers
- [ ] Scope so klein wie möglich

### Klassen/Structs:
- [ ] Single Responsibility
- [ ] Kleine, cohäsive Klassen
- [ ] Dependency Injection verwendet
- [ ] Interface-basierte Abhängigkeiten

### Allgemein:
- [ ] Keine Code-Duplikation (DRY)
- [ ] Konsistente Formatierung
- [ ] Selbstdokumentierender Code
- [ ] Alle Tests grün

## 🚀 Universal Clean Code Patterns

### Service Layer Pattern:
```go
// ✅ Clean Service Implementation
type BusinessService struct {
    dataProvider   DataProviderInterface
    calculator     CalculatorInterface
    repository     RepositoryInterface
    logger         LoggerInterface
}

func (s *BusinessService) ProcessBusinessLogic(
    ctx context.Context,
    criteria BusinessCriteria,
) ([]Result, error) {
    data, err := s.dataProvider.GetData(criteria.Filter)
    if err != nil {
        return nil, fmt.Errorf("failed to get data: %w", err)
    }
    
    var results []Result
    for _, item := range data {
        result, err := s.processItem(ctx, item, criteria)
        if err != nil {
            s.logger.Warn("Failed to process item", 
                "item", item.ID, 
                "error", err)
            continue
        }
        
        if s.calculator.MeetsCriteria(result, criteria) {
            results = append(results, result)
        }
    }
    
    return s.sortResults(results), nil
}
```

```typescript
// ✅ Clean TypeScript Service
class BusinessService {
    constructor(
        private readonly dataProvider: DataProvider,
        private readonly calculator: Calculator,
        private readonly repository: Repository,
        private readonly logger: Logger
    ) {}

    async processBusinessLogic(
        criteria: BusinessCriteria
    ): Promise<Result[]> {
        try {
            const data = await this.dataProvider.getData(criteria.filter);
            
            const results = await Promise.all(
                data.map(item => this.processItem(item, criteria))
            );

            const validResults = results
                .filter(result => result !== null)
                .filter(result => this.calculator.meetsCriteria(result, criteria));

            return this.sortResults(validResults);

        } catch (error) {
            this.logger.error('Failed to process business logic', {
                criteria,
                error: error.message
            });
            throw error;
        }
    }

    private async processItem(
        item: DataItem, 
        criteria: BusinessCriteria
    ): Promise<Result | null> {
        try {
            const calculatedResult = await this.calculator.calculate(item);
            await this.repository.save(calculatedResult);
            return calculatedResult;
        } catch (error) {
            this.logger.warn('Failed to process item', {
                itemId: item.id,
                error: error.message
            });
            return null;
        }
    }

    private sortResults(results: Result[]): Result[] {
        return results.sort((a, b) => b.priority - a.priority);
    }
}
```

---

**💡 Clean Code ist universell: Es geht um Kommunikation, Verständlichkeit und Wartbarkeit - unabhängig von Programmiersprache oder Projekt!**
