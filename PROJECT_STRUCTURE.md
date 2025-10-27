# Project Structure - Domain-Driven Design

## 📁 Complete File Structure

```
reika/
│
├── 📚 Documentation Files
│   ├── README.md                    # Overview & quick start
│   ├── ARCHITECTURE.md              # Deep dive arsitektur DDD
│   ├── SETUP.md                     # Panduan setup & troubleshooting
│   ├── MIGRATION_SUMMARY.md         # Summary migrasi
│   └── PROJECT_STRUCTURE.md         # File ini
│
├── 🎯 Domain Layer (Business Logic Core)
│   └── domain/
│       ├── transaction/
│       │   ├── entity.go           # ✅ Transaction entity dengan business rules
│       │   ├── repository.go       # ✅ Repository interface (port)
│       │   └── service.go          # ✅ Domain service
│       └── errors/
│           └── errors.go           # ✅ Domain-specific errors
│
├── 📋 Application Layer (Use Cases)
│   └── application/
│       ├── usecase/
│       │   └── extract_transactions.go  # ✅ Main use case
│       └── dto/
│           └── transaction_dto.go       # ✅ Data Transfer Objects
│
├── 🔧 Infrastructure Layer (External Services)
│   └── infrastructure/
│       ├── gemini/
│       │   └── client.go           # ✅ Gemini API implementation
│       └── file/
│           └── processor.go        # ✅ File processing
│
├── 🌐 Interface Layer (HTTP/API)
│   └── interfaces/
│       └── http/
│           ├── handler/
│           │   └── transaction_handler.go  # ✅ HTTP handlers
│           ├── middleware/
│           │   ├── cors.go         # ✅ CORS configuration
│           │   ├── logger.go       # ✅ Request logging
│           │   └── recovery.go     # ✅ Panic recovery
│           └── router/
│               └── router.go       # ✅ Route definitions
│
├── ⚙️ Configuration Layer
│   └── config/
│       ├── config.go               # ✅ Config management
│       └── container.go            # ✅ Dependency injection
│
├── 🚀 Application Entry Point
│   └── main.go                     # ✅ Clean entry point (43 lines!)
│
├── 📦 Go Modules
│   ├── go.mod                      # ✅ Dependencies
│   └── go.sum                      # ✅ Checksums
│
└── 🔒 Configuration Files
    ├── .gitignore                  # ✅ Git ignore rules
    └── .env                        # ⚠️  Create this (see SETUP.md)
```

## 📊 File Statistics

### Go Source Files: 16 files

#### Domain Layer (4 files)

- `domain/transaction/entity.go` - 98 lines
- `domain/transaction/repository.go` - 14 lines
- `domain/transaction/service.go` - 30 lines
- `domain/errors/errors.go` - 58 lines

#### Application Layer (2 files)

- `application/usecase/extract_transactions.go` - 48 lines
- `application/dto/transaction_dto.go` - 28 lines

#### Infrastructure Layer (2 files)

- `infrastructure/gemini/client.go` - 167 lines
- `infrastructure/file/processor.go` - 85 lines

#### Interface Layer (5 files)

- `interfaces/http/handler/transaction_handler.go` - 107 lines
- `interfaces/http/middleware/cors.go` - 13 lines
- `interfaces/http/middleware/logger.go` - 12 lines
- `interfaces/http/middleware/recovery.go` - 16 lines
- `interfaces/http/router/router.go` - 20 lines

#### Configuration Layer (2 files)

- `config/config.go` - 58 lines
- `config/container.go` - 46 lines

#### Entry Point (1 file)

- `main.go` - 43 lines (reduced from 170!)

### Documentation: 4 markdown files

- `README.md` - 327 lines
- `ARCHITECTURE.md` - 540 lines
- `SETUP.md` - 292 lines
- `MIGRATION_SUMMARY.md` - 400+ lines
- `PROJECT_STRUCTURE.md` - This file

**Total Documentation: ~1,600 lines**

## 🎯 Layer Dependencies

```
┌─────────────────────────────────────┐
│         main.go                     │
│    (Wires everything together)      │
└────────────────┬────────────────────┘
                 │
                 ▼
┌────────────────────────────────────────┐
│         Config Layer                   │
│  ┌──────────────────────────────────┐  │
│  │  config.go - Load configuration  │  │
│  │  container.go - Dependency       │  │
│  │                injection          │  │
│  └──────────────────────────────────┘  │
└────────────────┬───────────────────────┘
                 │
     ┌───────────┴────────────┐
     │                        │
     ▼                        ▼
┌─────────────┐        ┌──────────────┐
│  Interface  │        │Infrastructure│
│   Layer     │        │    Layer     │
│             │        │              │
│  - Handlers │───────▶│ - Gemini     │
│  - Router   │        │ - Files      │
│  - Middleware        │              │
└──────┬──────┘        └──────┬───────┘
       │                      │
       └──────────┬───────────┘
                  │
                  ▼
        ┌──────────────────┐
        │  Application     │
        │     Layer        │
        │                  │
        │  - Use Cases     │
        │  - DTOs          │
        └────────┬─────────┘
                 │
                 ▼
        ┌──────────────────┐
        │   Domain Layer   │
        │  (No dependencies)│
        │                  │
        │  - Entities      │
        │  - Interfaces    │
        │  - Services      │
        │  - Errors        │
        └──────────────────┘
```

**Dependency Flow: Semua dependencies mengalir ke dalam (inward) menuju Domain Layer**

## 🔄 Request Flow

### Example: Upload Transaction Flow

```
1. HTTP Request
   │
   ▼
2. interfaces/http/middleware/logger.go
   │ ├─ Log request
   ▼
3. interfaces/http/middleware/recovery.go
   │ ├─ Panic protection
   ▼
4. interfaces/http/middleware/cors.go
   │ ├─ CORS validation
   ▼
5. interfaces/http/router/router.go
   │ ├─ Route to handler
   ▼
6. interfaces/http/handler/transaction_handler.go
   │ ├─ Parse multipart form
   │ ├─ Validate request
   ▼
7. infrastructure/file/processor.go
   │ ├─ Process files
   │ ├─ Validate file types
   │ ├─ Check file sizes
   ▼
8. application/usecase/extract_transactions.go
   │ ├─ Convert DTO → Domain objects
   │ ├─ Execute business logic
   ▼
9. domain/transaction/service.go
   │ ├─ Apply business rules
   │ ├─ Call repository
   ▼
10. infrastructure/gemini/client.go (implements domain/transaction/repository.go)
    │ ├─ Call Gemini API
    │ ├─ Parse response
    │ ├─ Create domain entities
    ▼
11. domain/transaction/entity.go
    │ ├─ Validate business rules
    │ ├─ Create Transaction entities
    ▼
12. application/usecase/extract_transactions.go
    │ ├─ Convert Domain → DTO
    ▼
13. interfaces/http/handler/transaction_handler.go
    │ ├─ Return HTTP response
    ▼
14. HTTP Response (JSON)
```

## 🎨 Design Patterns Map

### Repository Pattern

```
Location: domain/transaction/repository.go (interface)
Implementation: infrastructure/gemini/client.go
Purpose: Abstract data access, decouple business logic from data source
```

### Use Case Pattern

```
Location: application/usecase/extract_transactions.go
Purpose: Encapsulate business operations, orchestrate flow
```

### Factory Pattern

```
Location: domain/transaction/entity.go (NewTransaction)
Purpose: Entity creation with validation
```

### Dependency Injection

```
Location: config/container.go
Purpose: Wire all dependencies, no global state
```

### Strategy Pattern

```
Location: domain/transaction/repository.go (interface)
Purpose: Swappable implementations (Gemini → OpenAI → etc)
```

## 📝 Key Files Explained

### 1. `main.go` (Entry Point)

**Purpose**: Application bootstrap

- Loads configuration
- Initializes DI container
- Sets up middleware
- Starts server

### 2. `config/container.go` (Dependency Injection)

**Purpose**: Wire all dependencies

- Creates all instances
- Injects dependencies
- Single source of truth

### 3. `domain/transaction/entity.go` (Business Entity)

**Purpose**: Core business object

- Encapsulated fields (private)
- Business validation
- Business methods

### 4. `domain/transaction/repository.go` (Port)

**Purpose**: Define data access contract

- Interface only
- No implementation
- Abstraction layer

### 5. `infrastructure/gemini/client.go` (Adapter)

**Purpose**: Implement repository interface

- External API calls
- Response parsing
- Error handling

### 6. `application/usecase/extract_transactions.go` (Use Case)

**Purpose**: Business operation orchestration

- Coordinate domain logic
- Convert DTOs ↔ Domain objects
- Handle application flow

### 7. `interfaces/http/handler/transaction_handler.go` (Handler)

**Purpose**: HTTP request handling

- Parse requests
- Validate input
- Call use cases
- Format responses

## 🔧 Configuration Files

### Required: `.env`

```env
PORT=5002
GEMINI_API_KEY=your_api_key_here
CORS_ALLOW_ORIGINS=http://localhost:3000
```

### Generated: `go.mod`

```
- Fiber v2 (HTTP framework)
- godotenv (env variables)
```

### Version Control: `.gitignore`

```
- .env (secrets)
- Binary files
- Vendor directory
- IDE files
```

## 📚 Documentation Map

| File                   | Purpose                    | Audience                 |
| ---------------------- | -------------------------- | ------------------------ |
| `README.md`            | Quick start & overview     | All developers           |
| `ARCHITECTURE.md`      | Deep dive DDD architecture | Architects & senior devs |
| `SETUP.md`             | Setup & troubleshooting    | New developers           |
| `MIGRATION_SUMMARY.md` | Migration details          | Team leads               |
| `PROJECT_STRUCTURE.md` | File organization          | All developers           |

## ✅ Quality Checklist

- [x] ✅ Linter: No errors
- [x] ✅ Build: Successful
- [x] ✅ Dependencies: Clean (go mod tidy)
- [x] ✅ Architecture: DDD compliant
- [x] ✅ SOLID: All principles applied
- [x] ✅ Error Handling: Comprehensive
- [x] ✅ Security: No hardcoded secrets
- [x] ✅ Documentation: Extensive
- [x] ✅ Backward Compatibility: Maintained

## 🚀 Quick Commands

### Development

```bash
# Run application
go run main.go

# Build binary
go build -o reika

# Clean build
go clean && go build

# Run tests (when added)
go test ./...
```

### Maintenance

```bash
# Update dependencies
go mod tidy

# Verify dependencies
go mod verify

# View dependency graph
go mod graph
```

## 🎓 Learning Path

1. **Start Here**: `README.md`
2. **Understand Setup**: `SETUP.md`
3. **Learn Architecture**: `ARCHITECTURE.md`
4. **Review Structure**: This file
5. **Study Code**:
   - Start with `domain/`
   - Then `application/`
   - Then `infrastructure/`
   - Finally `interfaces/`

## 📈 Next Steps for Developers

### Understanding the Code

1. Read `domain/transaction/entity.go` - Core business logic
2. Read `domain/transaction/repository.go` - Abstraction
3. Read `infrastructure/gemini/client.go` - Implementation
4. Read `application/usecase/extract_transactions.go` - Orchestration
5. Read `interfaces/http/handler/transaction_handler.go` - HTTP layer

### Adding Features

1. Create entity in `domain/`
2. Define repository interface (if needed)
3. Create use case in `application/`
4. Implement infrastructure in `infrastructure/`
5. Add handler in `interfaces/http/handler/`
6. Register route in `router.go`
7. Wire in `container.go`

### Testing

1. Unit test domain entities
2. Integration test use cases
3. Mock external services
4. E2E test HTTP endpoints

---

**Arsitektur ini production-ready dan siap untuk scale! 🚀**
