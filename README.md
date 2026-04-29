# Book Library API

A production-ready REST API for managing a book library, built with Go, Gin framework, and SQLite database. Features a clean architecture with service, repository, and handler layers, plus JWT authentication and role-based access control.

## Features

- **User Authentication** - JWT-based token authentication with secure password hashing
- **User Management** - User registration, login, and logout functionality
- **Role-Based Access Control** - Admin-only endpoints for book management operations
- **CRUD Operations** - Create, read, update, and delete books (admin restricted)
- **Advanced Search** - Search books by name and author with pagination
- **Pagination** - Configurable page size and offset for large datasets
- **SQLite Database** - Lightweight, file-based persistence
- **Clean Architecture** - Service, Repository, and Handler layers
- **Configuration Management** - YAML-based config with Viper
- **Input Validation** - Automatic request binding and validation
- **Consistent Error Handling** - Standardized JSON error responses
- **Database Migrations** - Automatic schema management with GORM
- **Middleware Support** - Authentication and authorization middleware
- **Unit Testing** - Service tests with mock repositories

## Tech Stack

| Component | Version |
|-----------|---------|
| Go | 1.25.0 |
| Gin Framework | v1.12.0 |
| GORM ORM | v1.31.1 |
| SQLite | v1.6.0 |
| Viper Config | v1.21.0 |

## Project Architecture

```
book_library_API v3/
├── cmd/
│   └── main.go                      # Application entry point
├── internal/
│   ├── config/
│   │   ├── config.yaml              # Application configuration
│   │   └── db.go                    # Database initialization
│   ├── dto/
│   │   ├── book_dto.go              # Book request/response DTOs
│   │   ├── user_dto.go              # User request/response DTOs
│   │   └── response_dto.go          # Response wrapper DTOs
│   ├── handler/
│   │   ├── book_handler.go          # Book HTTP handlers
│   │   └── user_handler.go          # User/Auth HTTP handlers
│   ├── middleware/
│   │   └── middleware.go            # Auth & authorization middleware
│   ├── models/
│   │   ├── book_model.go            # Book database model
│   │   └── user_model.go            # User database model
│   ├── repository/
│   │   ├── book_repository.go       # Book data access layer
│   │   ├── mock_book_repository.go  # Mock for book repository testing
│   │   ├── user_repository.go       # User data access layer
│   │   └── mock_user_repository.go  # Mock for user repository testing
│   ├── response/
│   │   └── response.go              # Response helper functions
│   ├── routes/
│   │   └── routes.go                # API route definitions
│   ├── service/
│   │   ├── book_service.go          # Book business logic layer
│   │   ├── book_service_test.go     # Unit tests for book service
│   │   ├── user_service.go          # User/Auth business logic
│   │   └── user_service_test.go     # Unit tests for user service
│   └── utils/
│       └── jwt_util.go              # JWT token utilities
├── database/                         # SQLite database files
├── go.mod                           # Go module dependencies
└── README.md                        # This file
```

## Installation

### Prerequisites

- Go 1.25.0 or higher
- Git

### Steps

1. Clone the repository:
```bash
git clone https://github.com/mohdareeb0x-commits/book-library-api.git
cd book_library_API\ v3
cd book-library-api
```

2. Install dependencies:
```bash
go mod download
```

3. Configure the application by creating/modifying `internal/config/config.yaml`:
```yaml
server:
  port: "8080"
```

4. Run the application:
```bash
go run cmd/main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Base URL
```
http://localhost:8080
```

### Endpoints Overview

| Method | Endpoint | Auth | Admin |
|--------|----------|------|-------|
| POST | `/user/register` | ✗ | ✗ |
| POST | `/user/login` | ✗ | ✗ |
| POST | `/user/logout` | ✓ | ✗ |
| GET | `/books` | ✓ | ✗ |
| GET | `/books/:id` | ✓ | ✗ |
| GET | `/books/search` | ✓ | ✗ |
| POST | `/admin/books` | ✓ | ✓ |
| PATCH | `/admin/books/:id` | ✓ | ✓ |
| DELETE | `/admin/books/:id` | ✓ | ✓ |

### Common Examples

**Register user:**
```bash
curl -X POST http://localhost:8080/user/register \
  -H "Content-Type: application/json" \
  -d '{"username": "john", "password": "pass123"}'
```

**Login:**
```bash
curl -X POST http://localhost:8080/user/login \
  -H "Content-Type: application/json" \
  -d '{"username": "john", "password": "pass123"}'
```

**Get all books (requires JWT token):**
```bash
curl -X GET http://localhost:8080/books \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

**Get specific book:**
```bash
curl -X GET http://localhost:8080/books/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

**Create book (admin):**
```bash
curl -X POST http://localhost:8080/admin/books \
  -H "Authorization: Bearer ADMIN_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Clean Code",
    "author": "Robert C. Martin",
    "price": 45.99,
    "units": 10,
    "date_published": "2008-08-01T00:00:00Z"
  }'
```

**Update a book (admin-only):**
```bash
curl -X PATCH http://localhost:8080/admin/books/1 \
  -H "Authorization: Bearer ADMIN_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "price": 49.99,
    "units": 15
  }'
```

**Delete a book (admin-only):**
```bash
curl -X DELETE http://localhost:8080/admin/books/1 \
  -H "Authorization: Bearer ADMIN_JWT_TOKEN"
```

**Search books:**
```bash
curl -X GET "http://localhost:8080/books/search?name=Clean&author=Martin&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

**Logout:**
```bash
curl -X POST http://localhost:8080/user/logout \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

## Authentication

- **JWT tokens** issued on login
- Include in header: `Authorization: Bearer <token>`
- **Public:** register, login
- **Protected:** all other endpoints
- **Admin-only:** create/update/delete books

## Testing

Run tests with mock repositories for isolated service testing:

```bash
# Run all tests
go test ./...

# Verbose output
go test -v ./...

# With coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Specific test
go test -run TestCreateBook ./internal/service/
```

### Available Tests

**Book Service Tests:**
- `TestCreateBook` - Book creation
- `TestListBook` - Book listing with pagination
- `TestListBookByID` - Get book by ID

**User Service Tests:**
- `TestRegisterSuccess` - User registration validation

### Test Files

- `internal/service/book_service_test.go`
- `internal/service/user_service_test.go`
- `internal/repository/mock_book_repository.go`
- `internal/repository/mock_user_repository.go`

## Build & Development

```bash
# Build for production
go build -o book-library-api cmd/main.go
./book-library-api
```

### Architecture Layers

- **Handler** - HTTP request/response
- **Service** - Business logic
- **Repository** - Data access
- **Model** - Database entities
- **DTO** - Request/response objects
- **Middleware** - Auth and logging

## License

MIT License

## Contributing

Pull requests are welcome!

## Support

Open an issue on [GitHub](https://github.com/mohdareeb0x-commits/book-library-api/issues).

## Author

[mohdareeb0x-commits](https://github.com/mohdareeb0x-commits)
