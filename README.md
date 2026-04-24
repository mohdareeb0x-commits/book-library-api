# Book Library API v3

A production-ready REST API for managing a book library, built with Go, Gin framework, and SQLite database. Features a clean architecture with service, repository, and handler layers.

## Features

-  **CRUD Operations** - Create, read, update, and delete books
-  **Advanced Search** - Search books by name and author with pagination
-  **Pagination** - Configurable page size and offset for large datasets
-  **SQLite Database** - Lightweight, file-based persistence
-  **Clean Architecture** - Service, Repository, and Handler layers
-  **Configuration Management** - YAML-based config with Viper
-  **Input Validation** - Automatic request binding and validation
-  **Consistent Error Handling** - Standardized JSON error responses
-  **Database Migrations** - Automatic schema management with GORM

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
│   │   ├── book_dto.go              # Request/response DTOs
│   │   └── response_dto.go          # Response wrapper DTOs
│   ├── handler/
│   │   └── book_handler.go          # HTTP request handlers
│   ├── models/
│   │   └── book_model.go            # Database models
│   ├── repository/
│   │   └── book_repository.go       # Data access layer
│   ├── response/
│   │   └── response.go              # Response helper functions
│   ├── routes/
│   │   └── routes.go                # API route definitions
│   └── service/
│       └── book_service.go          # Business logic layer
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

**Get all books (paginated):**
```bash
curl -X GET http://localhost:8080/books
```

**Get specific book:**
```bash
curl -X GET http://localhost:8080/books/1
```

**Create a book:**
```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Clean Code",
    "author": "Robert C. Martin",
    "price": 45.99,
    "units": 10,
    "date_published": "2008-08-01T00:00:00Z"
  }'
```

**Update a book:**
```bash
curl -X PATCH http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{
    "price": 49.99,
    "units": 15
  }'
```

**Delete a book:**
```bash
curl -X DELETE http://localhost:8080/books/1
```

**Search books:**
```bash
curl -X GET "http://localhost:8080/books/search?name=Clean&author=Martin&page=1&limit=10"
```

## API Endpoints Reference

### Base URL
```
http://localhost:8080
```

### 1. List All Books
```http
GET /books
```

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
POST /books
```

**Request Headers:**
```
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

### 4. Update Book
```http
PATCH /books/:id
```

**Parameters:**
| Name | Type | Description |
|------|------|-------------|
| `id` | integer | Book ID (path parameter) |

**Request Headers:**
```
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
DELETE /books/:id
```

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
| NOT_FOUND | 404 | Book not found by ID |
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
- **Config Layer** (`config/`) - Application configuration and initialization

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
