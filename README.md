# Reika - Transaction Extractor & Meeting Management

A Go application built with Domain-Driven Design (DDD) principles for extracting transaction data from documents (images and PDFs) using Google Gemini AI and managing meetings.

## 🏗️ Architecture

This project follows **Domain-Driven Design (DDD)** and **Clean Architecture** principles, organized into distinct layers:

```
reika/
├── domain/                 # Domain Layer (Business Logic Core)
│   ├── transaction/
│   │   ├── entity.go      # Transaction entity with business rules
│   │   ├── repository.go  # Repository interface (port)
│   │   └── service.go     # Domain services
│   └── errors/
│       └── errors.go      # Domain error types
│
├── application/           # Application Layer (Use Cases)
│   ├── usecase/
│   │   └── extract_transactions.go
│   └── dto/
│       └── transaction_dto.go
│
├── infrastructure/        # Infrastructure Layer (External Services)
│   ├── gemini/
│   │   └── client.go     # Gemini API implementation
│   └── file/
│       └── processor.go  # File processing logic
│
├── interfaces/           # Interface/Presentation Layer
│   └── http/
│       ├── handler/
│       │   └── transaction_handler.go
│       ├── middleware/
│       │   ├── cors.go
│       │   ├── logger.go
│       │   └── recovery.go
│       └── router/
│           └── router.go
│
├── config/              # Configuration & DI
│   ├── config.go       # Configuration management
│   └── container.go    # Dependency injection container
│
└── main.go             # Application entry point
```

## 📋 Layer Responsibilities

### 1. **Domain Layer** (Core Business Logic)

- Contains business entities and value objects
- Defines repository interfaces (ports)
- Implements domain services with business rules
- **No dependencies** on other layers
- Pure business logic

### 2. **Application Layer** (Use Cases)

- Orchestrates business flows
- Implements use cases (application services)
- Defines DTOs for data transfer
- Coordinates between domain and infrastructure

### 3. **Infrastructure Layer** (External Implementations)

- Implements repository interfaces
- Handles external service integrations (Gemini API)
- Manages file processing and I/O operations
- Database access (if applicable)

### 4. **Interface/Presentation Layer** (HTTP/API)

- HTTP handlers and routing
- Request/Response transformation
- Middleware (CORS, logging, recovery)
- API documentation

### 5. **Config Layer** (Configuration & DI)

- Configuration management
- Dependency injection container
- Wiring up all components

## 🚀 Getting Started

### Prerequisites

- Go 1.25 or higher
- Google Gemini API key

### Installation

1. Clone the repository

```bash
git clone <repository-url>
cd sandbox
```

2. Install dependencies

```bash
go mod download
```

3. Configure environment variables

```bash
cp .env.example .env
# Edit .env and add your GEMINI_API_KEY
```

4. Run the application

```bash
go run main.go
```

The server will start on port 5002 (or the port specified in your .env file).

## 📡 API Endpoints

### Upload and Extract Transactions

```
POST /api/upload
Content-Type: multipart/form-data

Parameters:
- file: One or more image/PDF files

Response:
[
  {
    "name": "John Doe",
    "type": "accommodation",
    "subtype": "hotel",
    "amount": 1000000,
    "total_night": 2,
    "subtotal": 2000000
  }
]
```

### Upload and Extract (Detailed Response)

```
POST /api/upload/detailed
Content-Type: multipart/form-data

Parameters:
- file: One or more image/PDF files

Response:
{
  "transactions": [...],
  "count": 5
}
```

### Health Check

```
GET /api/health

Response:
{
  "status": "healthy"
}
```

## 🎯 Design Patterns & Best Practices

### 1. **Dependency Inversion Principle**

- Domain layer defines interfaces
- Infrastructure layer implements them
- Dependencies point inward (towards domain)

### 2. **Repository Pattern**

- `ExtractorRepository` interface in domain layer
- `GeminiClient` implements it in infrastructure layer

### 3. **Use Case Pattern**

- Each business operation has a dedicated use case
- Clear separation of concerns
- Easy to test and maintain

### 4. **Dependency Injection**

- Constructor injection for all dependencies
- Centralized container for wiring components
- No global state

### 5. **Error Handling**

- Domain-specific error types
- Proper error wrapping and context
- Global error handler in HTTP layer

### 6. **Validation**

- Business rule validation in entities
- Input validation in use cases
- HTTP validation in handlers

## 🔧 Configuration

Environment variables:

| Variable             | Description           | Default               |
| -------------------- | --------------------- | --------------------- |
| `PORT`               | Server port           | 5002                  |
| `GEMINI_API_KEY`     | Google Gemini API key | Required              |
| `CORS_ALLOW_ORIGINS` | Allowed CORS origins  | http://localhost:3000 |

## 🧪 Testing Strategy

```
domain/          → Unit tests (pure business logic)
application/     → Integration tests (use cases)
infrastructure/  → Integration tests (external services)
interfaces/      → E2E tests (API endpoints)
```

## 📚 DDD Concepts Applied

1. **Entities**: `Transaction` with identity and business rules
2. **Value Objects**: `TransactionType` immutable type
3. **Repositories**: `ExtractorRepository` for data access abstraction
4. **Services**: `TransactionService` for domain logic
5. **Use Cases**: Application-specific business flows
6. **DTOs**: Data transfer between layers

## 🔒 Security Best Practices

- API key stored in environment variables
- CORS configuration for web security
- Request validation at multiple layers
- Error messages don't expose sensitive data
- Recovery middleware to prevent crashes

## 📈 Scalability Considerations

- **Horizontal Scaling**: Stateless design allows multiple instances
- **Service Isolation**: Easy to extract services to microservices
- **Cache Layer**: Can be added at infrastructure level
- **Queue Processing**: Can add async processing for large files

## 🛠️ Extending the Application

### Adding a New Use Case

1. Create use case in `application/usecase/`
2. Define DTOs in `application/dto/`
3. Create handler in `interfaces/http/handler/`
4. Register route in `interfaces/http/router/`
5. Wire dependencies in `config/container.go`

### Adding a New Repository

1. Define interface in `domain/`
2. Implement in `infrastructure/`
3. Update container for DI

## 📝 License

This project is for educational purposes.

## 👥 Contributing

Follow the existing architecture patterns and maintain separation of concerns when contributing.
