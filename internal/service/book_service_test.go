package service

import (
	"testing"
	"time"

	"github.com/mohdareeb0x-commits/book-library-api/internal/dto"
)

var s *BookService
var input dto.CreateBookInput
func TestCreateBook(t *testing.T) {
	input = dto.CreateBookInput{
		Name: "AOT",
		Author: "Isayama",
		Price: 300,
		Units: 20,
		DatePublished: time.Now(),
	}

	createdBook, err := s.CreateBook(input)
	if err != nil {
		t.Fatalf("Book addition test failed")
	}

	_, _, err = s.SearchBook(createdBook.Name, createdBook.Author, "1", "1")
	if err != nil {
		t.Fatalf("Book addition test failed")
	}
}

func TestListBooks(t *testing.T) {
	
}