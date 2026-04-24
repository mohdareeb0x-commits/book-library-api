package service

import (
	"errors"
	"strconv"

	"github.com/mohdareeb0x-commits/book-library-api/internal/dto"
	"github.com/mohdareeb0x-commits/book-library-api/internal/models"
	"github.com/mohdareeb0x-commits/book-library-api/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(input dto.CreateBookInput) (*models.Book, error) {
	book := models.Book{
		Name:          input.Name,
		Author:        input.Author,
		Price:         input.Price,
		Units:         input.Units,
		DatePublished: input.DatePublished,
	}
	return s.repo.Create(&book)
}

func (s *BookService) ListBooks(page, limit string) (*[]models.Book, *dto.Meta ,error) {
	int_limit, err := strconv.Atoi(limit)
	if err != nil {
		return nil, nil, err
	}

	int_page, err := strconv.Atoi(page)
	if err != nil {
		return nil, nil, err
	}

	offset := (int_page -  1) * int_limit

	books, err := s.repo.List(int_limit, offset)
	if err != nil {
		return nil, nil, err
	}

	meta := dto.Meta{
		Page: int_page,
		Limit: int_limit,
		Offset: offset,
	}

	return books, &meta, nil
} 

func (s *BookService) ListBooksByID(id string) (*models.Book, error) {
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return s.repo.ListByID(int_id)
}

func (s *BookService) UpdateBookByID(id string, input dto.UpdateBookInput) (*models.Book, error) {
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	book, err := s.repo.ListByID(int_id)
	if err != nil {
		return nil, err
	}

	updates := map[string]interface{}{}

	if input.Name != nil {
		updates["name"] = *input.Name
	}
	if input.Author != nil {
		updates["author"] = *input.Author
	}
	if input.Price != nil {
		updates["price"] = *input.Price
	}
	if input.Units != nil {
		updates["units"] = *input.Units
	}
	if input.DatePublished != nil {
		updates["date_published"] = *input.DatePublished
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	if err := s.repo.UpdateByID(book, updates); err != nil {
		return nil, err
	}

	return book, nil

}

func (s *BookService) DeleteBookByID(id string) (*models.Book, error) {
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// var book *models.Book
	book, err := s.repo.ListByID(int_id)
	if err != nil {
		return nil, err
	}

	if err := s.repo.DeleteByID(book); err != nil {
		return  nil, err
	}

	return book, nil
}
