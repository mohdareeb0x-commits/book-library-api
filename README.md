# Book Library API

A simple and efficient REST API for managing a book library, built with Go, Gin framework, and SQLite database.

## Features

-  **Create Books** - Add new books to the library
-  **Read Books** - Retrieve all books or search by ID
-  **Update Books** - Modify book details
-  **Delete Books** - Remove books from the library
-  **SQLite Database** - Lightweight, file-based database
-  **Structured Responses** - Consistent JSON response format

## Tech Stack

- **Language**: Go 1.25.0
- **Web Framework**: [Gin Gonic](https://github.com/gin-gonic/gin) v1.12.0
- **ORM**: [GORM](https://gorm.io/) v1.31.1
- **Database**: SQLite v1.6.0

## Project Structure

```
book_library_API/
├── cmd/
│   └── main.go           # Application entry point
├── handler/
│   ├── db.go             # Database operations and API handlers
│   └── response.go        # Response helper functions
├── models/
│   └── models.go         # Data models for Books and Response
├── routes/
│   └── routes.go         # API route definitions
├── go.mod               # Go module dependencies
└── README.md            # This file
```

## Installation

### Prerequisites

- Go 1.25.0 or higher
- Git

### Steps

1. Clone the repository:
```bash
git clone https://github.com/mohdareeb0x-commits/book-library-api.git
cd book-library-api
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/main.go
```

The API will start on `http://localhost:8080`

## API Endpoints

### 1. List All Books
```
GET /books
```
Returns all books in the library.

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "book": "The Go Programming Language",
      "author": "Alan Donovan",
      "date_published": "2015-10-26",
      "units": 5,
      "price": 45
    }
  ]
}
```

### 2. Get Book by ID
```
GET /books/:id
```
Retrieves a specific book by its ID.

**Parameters:**
- `id` (URL parameter) - The book ID

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "book": "The Go Programming Language",
    "author": "Alan Donovan",
    "date_published": "2015-10-26",
    "units": 5,
    "price": 45
  }
}
```

### 3. Create a New Book
```
POST /books?book=<name>&author=<author>&published=<date>&units=<units>&price=<price>
```
Adds a new book to the library.

**Query Parameters:**
- `book` (required) - Book title
- `author` (required) - Author name
- `published` (required) - Publication date (YYYY-MM-DD format)
- `units` (optional, default: 1) - Number of units available
- `price` (optional, default: 0) - Book price

**Example:**
```bash
curl "http://localhost:8080/books?book=Golang&author=John&published=2020-01-01&units=10&price=50"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 2,
    "book": "Golang",
    "author": "John",
    "date_published": "2020-01-01",
    "units": 10,
    "price": 50
  }
}
```

### 4. Update a Book
```
PUT /books/:id?book=<name>&author=<author>&...
```
Updates one or more fields of an existing book.

**Parameters:**
- `id` (URL parameter) - The book ID
- `book` (optional) - New book title
- `author` (optional) - New author name
- `published` (optional) - New publication date
- `units` (optional) - New number of units
- `price` (optional) - New price

**Example:**
```bash
curl -X PUT "http://localhost:8080/books/1?author=Alan%20A.%20Donovan&price=50"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "book": "The Go Programming Language",
    "author": "Alan A. Donovan",
    "date_published": "2015-10-26",
    "units": 5,
    "price": 50
  }
}
```

### 5. Delete a Book
```
DELETE /books/:id
```
Removes a book from the library.

**Parameters:**
- `id` (URL parameter) - The book ID

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "book": "The Go Programming Language",
    "author": "Alan Donovan",
    "date_published": "2015-10-26",
    "units": 5,
    "price": 45
  }
}
```

## Response Format

All API responses follow a consistent structure:

**Success Response:**
```json
{
  "success": true,
  "data": { /* response data */ }
}
```

**Error Response:**
```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Error description"
  }
}
```

## Database

The application uses SQLite with the database file stored as `library.db` in the root directory. The database schema is automatically created on first run using GORM's `AutoMigrate`.

### Book Model

| Field         | Type   | Description                  |
|---------------|--------|------------------------------|
| ID            | int    | Primary key (auto-generated) |
| Name          | string | Book title                   |
| Author        | string | Author name                  |
| DatePublished | string | Publication date             |
| Units         | int    | Number of copies available   |
| Price         | int    | Book price                   |

## Error Codes

| Code                  | HTTP Status | Description                           |
|-----------------------|-------------|---------------------------------------|
| REQUIRED_QUERY_EMPTY  | 400         | Required query parameters are missing |
| NO_BOOK_AVAILABLE     | 404         | Book with the specified ID not found  |
| NO_BOOKS_AVAILABLE    | 200         | Database is empty                     |
| INTERNAL_SERVER_ERROR | 500         | Server error (invalid parameters)     |

## Example Usage

### Using curl

```bash
# Get all books
curl http://localhost:8080/books

# Create a new book
curl "http://localhost:8080/books?book=Clean%20Code&author=Robert%20C.%20Martin&published=2008-08-01&units=3&price=40"

# Get a specific book
curl http://localhost:8080/books/1

# Update a book
curl -X PUT "http://localhost:8080/books/1?price=35"

# Delete a book
curl -X DELETE http://localhost:8080/books/1
```

## Author

Created by [mohdareeb0x-commits](https://github.com/mohdareeb0x-commits)
