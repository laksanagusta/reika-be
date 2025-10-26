# Architecture Documentation

## Domain-Driven Design (DDD) Implementation

This document provides an in-depth look at the Domain-Driven Design architecture implemented in this project.

## 🎯 Core Principles

### 1. Separation of Concerns

Each layer has a single, well-defined responsibility:

- **Domain**: Business logic and rules
- **Application**: Use case orchestration
- **Infrastructure**: External dependencies
- **Interface**: API and presentation

### 2. Dependency Rule

Dependencies flow inward:

```
Interface → Application → Domain
Infrastructure → Domain
```

The Domain layer has **zero dependencies** on other layers.

### 3. Dependency Inversion

- High-level modules don't depend on low-level modules
- Both depend on abstractions (interfaces)
- Abstractions defined in Domain layer

## 📊 Architecture Diagram

```
┌─────────────────────────────────────────────────────────┐
│                     main.go                              │
│  (Application Entry Point & Dependency Wiring)          │
└───────────────────────┬─────────────────────────────────┘
                        │
        ┌───────────────▼───────────────┐
        │      Config Layer             │
        │  - config.go (Configuration)  │
        │  - container.go (DI)          │
        └───────────────┬───────────────┘
                        │
    ┌───────────────────┴────────────────────┐
    │                                        │
┌───▼──────────────────┐          ┌─────────▼──────────────┐
│  Interface Layer     │          │  Infrastructure Layer  │
│  (HTTP/API)          │          │  (External Services)   │
│                      │          │                        │
│  ┌────────────────┐  │          │  ┌──────────────────┐  │
│  │   Handlers     │  │          │  │  Gemini Client   │  │
│  │   Middleware   │  │          │  │  File Processor  │  │
│  │   Router       │  │          │  └──────────────────┘  │
│  └────────────────┘  │          │                        │
└──────────┬───────────┘          └────────────┬───────────┘
           │                                   │
           │         ┌────────────────────────┘
           │         │
       ┌───▼─────────▼─────────────────┐
       │    Application Layer          │
       │    (Use Cases)                │
       │                               │
       │  ┌─────────────────────────┐  │
       │  │  Extract Transactions   │  │
       │  │  Use Case               │  │
       │  └─────────────────────────┘  │
       │                               │
       │  ┌─────────────────────────┐  │
       │  │  DTOs                   │  │
       │  └─────────────────────────┘  │
       └────────────┬──────────────────┘
                    │
       ┌────────────▼──────────────────┐
       │      Domain Layer             │
       │      (Business Logic)         │
       │                               │
       │  ┌─────────────────────────┐  │
       │  │  Transaction Entity     │  │
       │  │  (Business Rules)       │  │
       │  └─────────────────────────┘  │
       │                               │
       │  ┌─────────────────────────┐  │
       │  │  Repository Interface   │  │
       │  │  (Port)                 │  │
       │  └─────────────────────────┘  │
       │                               │
       │  ┌─────────────────────────┐  │
       │  │  Domain Service         │  │
       │  └─────────────────────────┘  │
       │                               │
       │  ┌─────────────────────────┐  │
       │  │  Domain Errors          │  │
       │  └─────────────────────────┘  │
       └───────────────────────────────┘
```

## 🔄 Request Flow

### Upload Transaction Flow

```
1. HTTP Request
   │
   ▼
2. [Interface Layer] TransactionHandler
   │ - Validates HTTP request
   │ - Extracts multipart files
   ▼
3. [Infrastructure] FileProcessor
   │ - Validates file types
   │ - Processes file content
   ▼
4. [Application Layer] ExtractTransactionsUseCase
   │ - Converts DTOs to domain objects
   │ - Orchestrates business logic
   ▼
5. [Domain Layer] TransactionService
   │ - Applies business rules
   │ - Uses repository interface
   ▼
6. [Infrastructure] GeminiClient (implements ExtractorRepository)
   │ - Calls external Gemini API
   │ - Parses response
   │ - Creates domain entities
   ▼
7. [Application Layer] Use Case
   │ - Converts domain entities to DTOs
   ▼
8. [Interface Layer] Handler
   │ - Returns HTTP response
   ▼
9. HTTP Response
```

## 📦 Layer Details

### Domain Layer (`domain/`)

**Purpose**: Core business logic, entities, and rules

**Components**:

- `transaction/entity.go`: Transaction entity with business logic

  - Encapsulated fields (private)
  - Business rule validation
  - Domain methods (IsAccommodation, CalculateTotal, etc.)

- `transaction/repository.go`: Repository interface (Port)

  - Defines contract for data access
  - No implementation details

- `transaction/service.go`: Domain service

  - Complex business logic spanning multiple entities
  - Uses repository interface

- `errors/errors.go`: Domain-specific errors
  - Business error types
  - Error wrapping utilities

**Key Principles**:

- ✅ No external dependencies
- ✅ Pure business logic
- ✅ Framework-agnostic
- ✅ Testable without mocks

### Application Layer (`application/`)

**Purpose**: Use case orchestration and application workflows

**Components**:

- `usecase/extract_transactions.go`: Main use case

  - Orchestrates the extraction flow
  - Converts between DTOs and domain objects
  - Handles application-level errors

- `dto/transaction_dto.go`: Data Transfer Objects
  - Request/Response structures
  - No business logic
  - JSON serialization

**Key Principles**:

- ✅ Depends only on Domain layer
- ✅ No framework coupling
- ✅ Application-specific logic
- ✅ DTO ↔ Domain conversion

### Infrastructure Layer (`infrastructure/`)

**Purpose**: External service implementations

**Components**:

- `gemini/client.go`: Gemini API implementation

  - Implements `ExtractorRepository` interface
  - HTTP client management
  - API response parsing
  - Error handling

- `file/processor.go`: File processing
  - File validation
  - MIME type detection
  - Size limits
  - Content processing

**Key Principles**:

- ✅ Implements domain interfaces
- ✅ External dependency management
- ✅ Infrastructure concerns only
- ✅ Swappable implementations

### Interface Layer (`interfaces/`)

**Purpose**: HTTP API and request handling

**Components**:

- `http/handler/transaction_handler.go`: HTTP handlers

  - Request parsing
  - Response formatting
  - HTTP-specific logic

- `http/middleware/`: Middleware components

  - CORS configuration
  - Logging
  - Recovery (panic handling)

- `http/router/router.go`: Route definitions
  - Endpoint mapping
  - Handler registration

**Key Principles**:

- ✅ HTTP-specific concerns
- ✅ No business logic
- ✅ Delegates to use cases
- ✅ Input validation

### Config Layer (`config/`)

**Purpose**: Configuration and dependency injection

**Components**:

- `config.go`: Configuration management

  - Environment variable loading
  - Configuration validation
  - Default values

- `container.go`: Dependency injection
  - Creates all dependencies
  - Wires components together
  - Single source of truth for DI

**Key Principles**:

- ✅ Constructor injection
- ✅ No global state
- ✅ Centralized wiring
- ✅ Configuration validation

## 🎨 Design Patterns Used

### 1. Repository Pattern

```go
// Domain defines the interface
type ExtractorRepository interface {
    ExtractFromDocuments(ctx context.Context, documents []Document) ([]*Transaction, error)
}

// Infrastructure implements it
type GeminiClient struct { ... }
func (c *GeminiClient) ExtractFromDocuments(...) { ... }
```

### 2. Use Case Pattern

```go
type ExtractTransactionsUseCase struct {
    transactionService *transaction.Service
}

func (uc *ExtractTransactionsUseCase) Execute(...) { ... }
```

### 3. Dependency Injection

```go
// Constructor injection
func NewTransactionHandler(
    extractUseCase *usecase.ExtractTransactionsUseCase,
    fileProcessor *file.Processor,
) *TransactionHandler {
    return &TransactionHandler{
        extractUseCase: extractUseCase,
        fileProcessor:  fileProcessor,
    }
}
```

### 4. Factory Pattern

```go
// Entity factory with validation
func NewTransaction(name, txType, subtype string, ...) (*Transaction, error) {
    // Validation
    // Business rules
    return &Transaction{...}, nil
}
```

### 5. Strategy Pattern

- Repository interface allows different extraction strategies
- Can swap Gemini with another AI service

## 🧪 Testing Strategy

### Unit Tests

- **Domain Layer**: Pure business logic tests
  ```go
  func TestTransaction_CalculateTotal(t *testing.T) { ... }
  ```

### Integration Tests

- **Application Layer**: Use case tests with mock repositories

  ```go
  func TestExtractTransactionsUseCase_Execute(t *testing.T) { ... }
  ```

- **Infrastructure Layer**: External service tests
  ```go
  func TestGeminiClient_ExtractFromDocuments(t *testing.T) { ... }
  ```

### E2E Tests

- **Interface Layer**: HTTP endpoint tests
  ```go
  func TestTransactionHandler_UploadAndExtract(t *testing.T) { ... }
  ```

## 🔧 Extension Points

### Adding New Features

1. **New Entity**:

   - Create in `domain/`
   - Add repository interface if needed
   - Implement in `infrastructure/`

2. **New Use Case**:

   - Create in `application/usecase/`
   - Define DTOs in `application/dto/`
   - Wire in `config/container.go`

3. **New Endpoint**:
   - Add handler method
   - Register route
   - Use existing or new use case

### Switching External Services

Example: Replace Gemini with OpenAI

1. Create `infrastructure/openai/client.go`
2. Implement `ExtractorRepository` interface
3. Update `config/container.go` to use new client
4. No changes needed in domain or application layers!

## 📈 Benefits of This Architecture

### 1. **Maintainability**

- Clear separation of concerns
- Each layer has single responsibility
- Easy to locate and fix bugs

### 2. **Testability**

- Domain logic easily testable
- Mock external dependencies
- Integration tests at boundaries

### 3. **Flexibility**

- Swap implementations easily
- Add features without breaking existing code
- Framework independence

### 4. **Scalability**

- Clear boundaries for microservices
- Can extract layers to separate services
- Independent deployment of layers

### 5. **Team Collaboration**

- Teams can work on different layers
- Clear contracts (interfaces)
- Reduced merge conflicts

## 🚦 Best Practices Implemented

1. ✅ **SOLID Principles**

   - Single Responsibility
   - Open/Closed
   - Liskov Substitution
   - Interface Segregation
   - Dependency Inversion

2. ✅ **Clean Code**

   - Meaningful names
   - Small functions
   - No comments needed (self-documenting)

3. ✅ **Error Handling**

   - Domain-specific errors
   - Proper error wrapping
   - Contextual error messages

4. ✅ **Security**

   - No hardcoded secrets
   - Input validation at boundaries
   - CORS configuration

5. ✅ **Performance**
   - Context propagation
   - Timeout handling
   - Resource cleanup (defer)

## 📚 Further Reading

- [Domain-Driven Design by Eric Evans](https://www.domainlanguage.com/ddd/)
- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Implementing DDD in Go](https://threedots.tech/post/ddd-lite-in-go-introduction/)
- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)
