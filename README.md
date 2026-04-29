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

## Tech Stack

| Component | Version | Description |
|-----------|---------|-------------|
| **Language** | Go 1.25.0 | Programming language |
| **Framework** | Gin v1.12.0 | HTTP web framework |
| **ORM** | GORM v1.31.1 | Object-relational mapping |
| **Database** | SQLite v1.6.0 | SQL database engine |
| **Config** | Viper v1.21.0 | Configuration management |

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
│   │   └── user_repository.go       # User data access layer
│   ├── response/
│   │   └── response.go              # Response helper functions
│   ├── routes/
│   │   └── routes.go                # API route definitions
│   ├── service/
│   │   ├── book_service.go          # Book business logic layer
│   │   └── user_service.go          # User/Auth business logic
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

The API server will start on `http://localhost:8080`

## Quick Start

### Using cURL

**Register a new user:**
```bash
curl -X POST http://localhost:8080/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "password": "SecurePassword123!"
  }'
```

**Login to get JWT token:**
```bash
curl -X POST http://localhost:8080/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "password": "SecurePassword123!"
  }'
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

**Create a book (admin-only):**
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

## API Endpoints Reference

### Base URL
```
http://localhost:8080
```

### Endpoint Summary

| Method | Endpoint | Description | Auth Required | Admin Only |
|--------|----------|-------------|----------------|-----------|
| POST | `/user/register` | Create new user account | ✗ | ✗ |
| POST | `/user/login` | Authenticate user and get token | ✗ | ✗ |
| POST | `/user/logout` | Logout user session | ✓ | ✗ |
| GET | `/books` | List all books (paginated) | ✓ | ✗ |
| GET | `/books/:id` | Get book by ID | ✓ | ✗ |
| GET | `/books/search` | Search books | ✓ | ✗ |
| POST | `/admin/books` | Create new book | ✓ | ✓ |
| PATCH | `/admin/books/:id` | Update book | ✓ | ✓ |
| DELETE | `/admin/books/:id` | Delete book | ✓ | ✓ |

### 1. List All Books
```http
GET /books
```

**Authentication:** Required (Bearer token)

**Query Parameters:**
| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `page` | integer | 1 | Page number for pagination |
| `limit` | integer | 10 | Number of books per page |

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "The Go Programming Language",
      "author": "Alan Donovan and Brian Kernighan",
      "date_published": "2015-10-26T00:00:00Z",
      "units": 5,
      "price": 45.99,
      "created_at": "2024-04-20T10:30:00Z",
      "updated_at": "2024-04-20T10:30:00Z"
    }
  ],
  "meta": {
    "page": 1,
    "limit": 10
  }
}
```

### 2. Get Book by ID
```http
GET /books/:id
```

**Authentication:** Required (Bearer token)

**Parameters:**
| Name | Type | Description |
|------|------|-------------|
| `id` | integer | Book ID (path parameter) |

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "The Go Programming Language",
    "author": "Alan Donovan and Brian Kernighan",
    "date_published": "2015-10-26T00:00:00Z",
    "units": 5,
    "price": 45.99,
    "created_at": "2024-04-20T10:30:00Z",
    "updated_at": "2024-04-20T10:30:00Z"
  }
}
```

**Response (404 Not Found):**
```json
{
  "success": false,
  "error": {
    "code": "NOT_FOUND",
    "message": "unable to get books by id: 999"
  }
}
```

### 3. Create Book
```http
POST /admin/books
```

**Authentication:** Required (Bearer token with admin role)

**Request Headers:**
```
Authorization: Bearer <admin_token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Clean Code",
  "author": "Robert C. Martin",
  "price": 45.99,
  "units": 10,
  "date_published": "2008-08-01T00:00:00Z"
}
```

**Field Validation:**
| Field | Type | Required | Constraints |
|-------|------|----------|-------------|
| `name` | string | ✓ Yes | Non-empty |
| `author` | string | ✓ Yes | Non-empty |
| `price` | float64 | ✗ No | ≥ 0 |
| `units` | integer | ✗ No | ≥ 0 |
| `date_published` | ISO 8601 | ✗ No | Valid datetime |

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Clean Code",
    "author": "Robert C. Martin",
    "date_published": "2008-08-01T00:00:00Z",
    "units": 10,
    "price": 45.99,
    "created_at": "2024-04-20T10:30:00Z",
    "updated_at": "2024-04-20T10:30:00Z"
  }
}
```

**Response (400 Bad Request):**
```json
{
  "success": false,
  "error": {
    "code": "JSON_BINDING_ERROR",
    "message": "unable to get request body"
  }
}
```

**Response (403 Forbidden):**
```json
{
  "success": false,
  "error": {
    "code": "FORBIDDEN",
    "message": "admin access required"
  }
}
```

### 4. Update Book
```http
PATCH /admin/books/:id
```

**Authentication:** Required (Bearer token with admin role)

**Parameters:**
| Name | Type | Description |
|------|------|-------------|
| `id` | integer | Book ID (path parameter) |

**Request Headers:**
```
Authorization: Bearer <admin_token>
Content-Type: application/json
```

**Request Body (all fields optional):**
```json
{
  "name": "Clean Code: A Handbook of Agile Software Craftsmanship",
  "author": "Robert C. Martin",
  "price": 49.99,
  "units": 15,
  "date_published": "2008-08-01T00:00:00Z"
}
```

**Field Validation:**
| Field | Type | Constraints |
|-------|------|-------------|
| `name` | string | Non-empty if provided |
| `author` | string | Non-empty if provided |
| `price` | float64 | ≥ 0 if provided |
| `units` | integer | ≥ 0 if provided |
| `date_published` | ISO 8601 | Valid datetime if provided |

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Clean Code: A Handbook of Agile Software Craftsmanship",
    "author": "Robert C. Martin",
    "date_published": "2008-08-01T00:00:00Z",
    "units": 15,
    "price": 49.99,
    "created_at": "2024-04-20T10:30:00Z",
    "updated_at": "2024-04-20T10:35:00Z"
  }
}
```

### 5. Delete Book
```http
DELETE /admin/books/:id
```

**Authentication:** Required (Bearer token with admin role)

**Parameters:**
| Name | Type | Description |
|------|------|-------------|
| `id` | integer | Book ID (path parameter) |

**Request Headers:**
```
Authorization: Bearer <admin_token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Clean Code",
    "author": "Robert C. Martin",
    "date_published": "2008-08-01T00:00:00Z",
    "units": 10,
    "price": 45.99,
    "created_at": "2024-04-20T10:30:00Z",
    "updated_at": "2024-04-20T10:35:00Z"
  }
}
```

**Response (500 Internal Server Error):**
```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_SERVER_ERROR",
    "message": "unable to delete book by id: 1"
  }
}
```

### 6. Search Books
```http
GET /books/search?name=query&author=query&page=1&limit=10
```

**Authentication:** Required (Bearer token)

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `name` | string | Search by book name (partial match) |
| `author` | string | Search by author name (partial match) |
| `page` | integer | Page number (default: 1) |
| `limit` | integer | Results per page (default: 10) |

**Usage Examples:**
```bash
# Search by book name
GET /books/search?name=Clean

# Search by author
GET /books/search?author=Martin

# Search by both name and author
GET /books/search?name=Code&author=Martin

# Search with pagination
GET /books/search?name=Go&page=2&limit=5
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Clean Code",
      "author": "Robert C. Martin",
      "date_published": "2008-08-01T00:00:00Z",
      "units": 10,
      "price": 45.99,
      "created_at": "2024-04-20T10:30:00Z",
      "updated_at": "2024-04-20T10:30:00Z"
    }
  ],
  "meta": {
    "page": 1,
    "limit": 10
  }
}
```

## User Authentication Endpoints

### 7. User Registration
```http
POST /user/register
```

**Request Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "username": "john_doe",
  "password": "SecurePassword123!"
}
```

**Field Validation:**
| Field | Type | Required | Constraints |
|-------|------|----------|-------------|
| `username` | string | ✓ Yes | Non-empty, unique |
| `password` | string | ✓ Yes | Non-empty |

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "username": "john_doe",
    "role": "user",
    "created_at": "2024-04-20T10:30:00Z"
  }
}
```

**Response (400 Bad Request):**
```json
{
  "success": false,
  "error": {
    "code": "JSON_BINDING_ERROR",
    "message": "unable to get request body"
  }
}
```

### 8. User Login
```http
POST /user/login
```

**Request Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "username": "john_doe",
  "password": "SecurePassword123!"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "john_doe",
      "role": "user"
    }
  }
}
```

**Response (401 Unauthorized):**
```json
{
  "success": false,
  "error": {
    "code": "UNAUTHORIZED",
    "message": "invalid username or password"
  }
}
```

### 9. User Logout
```http
POST /user/logout
```

**Request Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "message": "logged out successfully"
  }
}
```

**Response (401 Unauthorized):**
```json
{
  "success": false,
  "error": {
    "code": "UNAUTHORIZED",
    "message": "missing or invalid token"
  }
}
```

## Authentication & Authorization

### JWT Tokens
- Tokens are issued upon successful login
- Include token in the `Authorization` header for protected endpoints
- Token format: `Bearer <token>`

### User Roles
- **user** - Default role with read-only access to books
- **admin** - Full access including book creation, update, and deletion

### Protected Endpoints
All book endpoints (list, get, search) require valid JWT token in the Authorization header.

**Admin-Only Endpoints:**
- `POST /admin/books` - Create book
- `PATCH /admin/books/:id` - Update book
- `DELETE /admin/books/:id` - Delete book

### Authorization Example
```bash
# Login to get token
LOGIN_RESPONSE=$(curl -X POST http://localhost:8080/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin_user",
    "password": "admin_password"
  }')

TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.data.token')

# Use token in subsequent requests
curl -X POST http://localhost:8080/admin/books \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Design Patterns",
    "author": "Gang of Four",
    "price": 49.99,
    "units": 20,
    "date_published": "1994-01-01T00:00:00Z"
  }'
```

## Data Models

### Book Model
```go
type Book struct {
    ID            uint      `gorm:"primaryKey"`
    Name          string    `gorm:"not null"`
    Author        string    `gorm:"not null"`
    Price         float64
    Units         int
    DatePublished time.Time
    CreatedAt     time.Time
    UpdatedAt     time.Time
}
```

### User Model
```go
type User struct {
	ID        uint   `gorm:"primaryKey"`
	UserName  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
```

### CreateBookInput DTO
```go
type CreateBookInput struct {
	Name          string    `json:"name" binding:"required"`
	Author        string    `json:"author" binding:"required"`
	Price         float64   `json:"price" binding:"gte=0"`
	Units         int       `json:"units" binding:"gte=0"`
	DatePublished time.Time `json:"date_published"`
}
```

### UpdateBookInput DTO
```go
type UpdateBookInput struct {
	Name          *string    `json:"name"`
	Author        *string    `json:"author"`
	Price         *float64   `json:"price" binding:"omitempty,gte=0"`
	Units         *int       `json:"units" binding:"omitempty,gte=0"`
	DatePublished *time.Time `json:"date_published"`
}
```

### CreateUserInput DTO
```go
type CreateUserInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
```

### LoginUserInput DTO
```go
type LoginUserInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
```

### Response Wrapper
```go
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Page   int `json:"page,omitempty"`
	Limit  int `json:"limit,omitempty"`
}
```

## Error Handling

The API returns errors with appropriate HTTP status codes and consistent JSON format.

### HTTP Status Codes

| Code | Meaning | Scenario |
|------|---------|----------|
| 200 | OK | Successful request |
| 400 | Bad Request | Invalid JSON or validation error |
| 404 | Not Found | Book ID does not exist |
| 500 | Internal Server Error | Database or server error |

### Error Response Structure
```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable error message"
  }
}
```

### Common Error Codes

| Code | Status | Meaning |
|------|--------|---------|
| JSON_BINDING_ERROR | 400 | Invalid JSON in request body |
| NOT_FOUND | 404 | Book or user not found |
| UNAUTHORIZED | 401 | Invalid credentials or missing token |
| FORBIDDEN | 403 | Insufficient permissions (non-admin user) |
| INTERNAL_SERVER_ERROR | 500 | Server error during processing |

## Configuration

The application uses YAML-based configuration via Viper. Configuration file: `internal/config/config.yaml`

**Available Settings:**
```yaml
server:
  port: "8080"
```

**Environment Variables:**
You can override YAML settings with environment variables.

## Architecture

This project follows **Clean Architecture** principles:

- **Handler Layer** (`handler/`) - HTTP request processing and response formatting
- **Service Layer** (`service/`) - Business logic and data processing
- **Repository Layer** (`repository/`) - Data access and database operations
- **Model Layer** (`models/`) - Database schemas and entities
- **DTO Layer** (`dto/`) - Request/response data transfer objects
- **Middleware Layer** (`middleware/`) - Authentication and authorization
- **Config Layer** (`config/`) - Application configuration and initialization

### Middleware

- **AuthMiddleware** - Validates JWT tokens on protected routes
- **AdminOnly** - Restricts access to admin-only endpoints

## Development

### Project Setup
```bash
# Clone repository
git clone https://github.com/mohdareeb0x-commits/book-library-api.git
cd book-library-api

# Install dependencies
go mod download

# Run application
go run cmd/main.go
```

### Directory Organization
- `cmd/` - Application entry point
- `internal/` - Unexported packages (internal to the application)
- `database/` - SQLite database file location

### Build for Production
```bash
go build -o book-library-api cmd/main.go
./book-library-api
```

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

For issues and questions, please open an issue on the [GitHub repository](https://github.com/mohdareeb0x-commits/book-library-api/issues).

## Author

Created by [mohdareeb0x-commits](https://github.com/mohdareeb0x-commits)
