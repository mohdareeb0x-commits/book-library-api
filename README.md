# Book Library API

A simple and efficient REST API for managing a book library, built with Go, Gin framework, and SQLite database.

## Features

- **Create Books** - Add new books to the library
- **Read Books** - Retrieve all books or search by ID
- **Update Books** - Modify book details (PUT/PATCH)
- **Delete Books** - Remove books from the library
- **SQLite Database** - Lightweight, file-based database
- **Structured Responses** - Consistent JSON response format with error handling
- **Form Binding** - Automatic form data binding and validation

## Tech Stack

- **Language**: Go 1.25.0
- **Web Framework**: [Gin Gonic](https://github.com/gin-gonic/gin) v1.12.0
- **ORM**: [GORM](https://gorm.io/) v1.31.1
- **Database**: SQLite v1.6.0

## Project Structure

```
book_library_API/
├── cmd/
│   └── main.go              # Application entry point
├── handler/
│   ├── db.go                # Database operations and API handlers
│   └── response.go          # Response helper functions
├── models/
│   └── models.go            # Data models for Books and Response
├── routes/
│   └── routes.go            # API route definitions
├── go.mod                   # Go module dependencies
└── README.md                # This file
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

The API server will start on `http://localhost:8080`

## API Endpoints

### Get All Books
```
GET /books
```
Returns a list of all books in the library.

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "The Go Programming Language",
      "author": "Alan Donovan and Brian Kernighan",
      "date_published": "2015-10-26",
      "units": 5,
      "price": 45
    }
  ]
}
```

### Get Book by ID
```
GET /books/:id
```
Retrieve a specific book by its ID.

**Parameters:**
- `id` (integer, required) - Book ID

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "The Go Programming Language",
    "author": "Alan Donovan and Brian Kernighan",
    "date_published": "2015-10-26",
    "units": 5,
    "price": 45
  }
}
```

### Create Book
```
POST /books
```
Add a new book to the library.

**Request Body (Form Data):**
- `book` (string, required) - Book name
- `author` (string, required) - Author name
- `date_published` (string, required) - Publication date
- `units` (integer, default = 0) - Number of units available
- `price` (integer, default = 0) - Book price

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "The Go Programming Language",
    "author": "Alan Donovan and Brian Kernighan",
    "date_published": "2015-10-26",
    "units": 5,
    "price": 45
  }
}
```

### Update Book
```
PUT /books/:id
PATCH /books/:id
```
Update an existing book's details.

**Parameters:**
- `id` (integer, required) - Book ID

**Request Body (Form Data):**
- `book` (string) - Book name
- `author` (string) - Author name
- `date_published` (string) - Publication date
- `units` (integer) - Number of units available
- `price` (integer) - Book price

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "The Go Programming Language",
    "author": "Alan Donovan and Brian Kernighan",
    "date_published": "2015-10-26",
    "units": 10,
    "price": 50
  }
}
```

### Delete Book
```
DELETE /books/:id
```
Remove a book from the library.

**Parameters:**
- `id` (integer, required) - Book ID

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "The Go Programming Language",
    "author": "Alan Donovan and Brian Kernighan",
    "date_published": "2015-10-26",
    "units": 5,
    "price": 45
  }
}
```

## Data Models

### Books
```go
type Books struct {
	ID            int    `form:"id"`
	Name          string `form:"book" binding:"required"`
	Author        string `form:"author" binding:"required"`
	DatePublished string `form:"date_published" binding:"required"`
	Units         int    `form:"units" binding:"required"`
	Price         int    `form:"price" binding:"required"`
}
```

### Response
```go
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
```

## Error Handling

The API returns consistent error responses with appropriate HTTP status codes:

- **400 Bad Request** - Form binding error
- **404 Not Found** - Book not found by ID
- **500 Internal Server Error** - Database or processing errors

Example error response:
```json
{
  "success": false,
  "error": {
    "code": "NO_BOOK_AVAILABLE",
    "message": "no book available with id: 999"
  }
}
```

## Database

The application uses SQLite with automatic schema migration. On first run, the `library.db` file will be created with the Books table.


## Author

Created by [mohdareeb0x-commits](https://github.com/mohdareeb0x-commits)
