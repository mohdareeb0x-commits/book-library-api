package repository

import "github.com/mohdareeb0x-commits/book-library-api/internal/models"

type MockBookRepository struct {
	CreateFunc func (book *models.Book) (*models.Book, error)
	ListFunc func (limit, offset int) (*[]models.Book, error)
	ListByIDFunc func (id int) (*models.Book, error)
	UpdateByIDFunc func (book *models.Book, updates map[string]interface{}) error
	DeleteByIDFunc func (book *models.Book) error
	SearchByAuthorFunc func (author string) (*[]models.Book, error) 
	SearchByNameFunc func (name string) (*[]models.Book, error)
	SearchFunc func (name, author string, limit, offset int) (*[]models.Book, error)
}

func (m *MockBookRepository) Create(book *models.Book) (*models.Book, error) {
	return m.CreateFunc(book)
}

func (m *MockBookRepository) List(limit, offset int) (*[]models.Book, error) {
	return m.ListFunc(limit, offset)
}

func (m *MockBookRepository) ListByID(id int) (*models.Book, error) {
	return m.ListByIDFunc(id)
}

func (m *MockBookRepository) UpdateByID(book *models.Book, updates map[string]interface{}) error {
	return m.UpdateByIDFunc(book, updates)
}

func (m *MockBookRepository) DeleteByID(book *models.Book) error {
	return m.DeleteByIDFunc(book)
}

func (m *MockBookRepository) SearchByAuthor(author string) (*[]models.Book, error) {
	return m.SearchByAuthorFunc(author)
}

func (m *MockBookRepository) SearchByName(name string) (*[]models.Book, error) {
	return m.SearchByNameFunc(name)
}

func (m *MockBookRepository) Search(name, author string, limit, offset int) (*[]models.Book, error) {
	return m.SearchFunc(name, author, limit, offset)
}