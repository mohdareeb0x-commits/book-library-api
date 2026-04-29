package service

import (
	"testing"
	"time"

	"github.com/mohdareeb0x-commits/book-library-api/internal/dto"
	"github.com/mohdareeb0x-commits/book-library-api/internal/models"
	"github.com/mohdareeb0x-commits/book-library-api/internal/repository"
)

func TestCreateBook(t *testing.T) {
	mockRepo := &repository.MockBookRepository{
		CreateFunc: func(book *models.Book) (*models.Book, error) {
			book.ID = 1
			return book, nil
		},
	}
	service := NewBookService(mockRepo)

	input := dto.CreateBookInput{
		Name:          "AOT",
		Author:        "Isayama",
		Price:         300,
		Units:         25,
		DatePublished: time.Now(),
	}

	_, err := service.CreateBook(input)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestListBook(t *testing.T) {
	mockRepo := &repository.MockBookRepository{
		ListFunc: func(limit, offset int) (*[]models.Book, error) {
			return &[]models.Book{{ID: 1,},}, nil
		},
	}
	service := NewBookService(mockRepo)

	_, _, err := service.ListBooks("1", "1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestListBookByID(t *testing.T) {
	mockRepo := &repository.MockBookRepository{
		ListByIDFunc: func(id int) (*models.Book, error) {
			return &models.Book{ID: uint(id)}, nil
		},
	}
	service := NewBookService(mockRepo)

	_, err := service.ListBooksByID("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestUpdateBookByID(t *testing.T) {
	mockRepo := &repository.MockBookRepository{
		ListByIDFunc: func(id int) (*models.Book, error) {
			return &models.Book{ID: uint(id)}, nil
		},
		UpdateByIDFunc: func(book *models.Book, updates map[string]interface{}) error {
			return nil
		},
	}
	service := NewBookService(mockRepo)
	
	name := "Attack On Titan"

	input := dto.UpdateBookInput{
		Name: &name,
	}

	_, err := service.UpdateBookByID("1", input)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestDeleteBookByID(t *testing.T) {
	mockRepo := &repository.MockBookRepository{
		ListByIDFunc: func(id int) (*models.Book, error) {
			return &models.Book{ID: uint(id)}, nil
		},
		DeleteByIDFunc: func(book *models.Book) error {
			return nil
		},
	}
	service := NewBookService(mockRepo)

	_, err := service.DeleteBookByID("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestSearch(t *testing.T) {
	mockRepo := &repository.MockBookRepository{
		SearchFunc: func(name, author string, limit, offset int) (*[]models.Book, error) {
			return &[]models.Book{{Name: name}}, nil
		},
	}
	service := NewBookService(mockRepo)

	_, _, err := service.SearchBook("Aot", "Isayama", "1", "1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestSearchByName(t *testing.T) {
	mockRepo := &repository.MockBookRepository{
		SearchByNameFunc: func(name string) (*[]models.Book, error) {
			return &[]models.Book{{Name: name}}, nil
		},
	}
	service := NewBookService(mockRepo)

	_, _, err := service.SearchBook("Aot", "", "1", "1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestSearchByAuthor(t *testing.T) {
	mockRepo := &repository.MockBookRepository{
		SearchByAuthorFunc: func(author string) (*[]models.Book, error) {
			return &[]models.Book{{Author: author}}, nil
		},
	}
	service := NewBookService(mockRepo)

	_, _, err := service.SearchBook("", "Isayama", "1", "1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
