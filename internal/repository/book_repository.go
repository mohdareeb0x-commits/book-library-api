package repository

import (
	"fmt"
	"log"

	"github.com/mohdareeb0x-commits/book-library-api/internal/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *BookRepository) Create(book *models.Book) (*models.Book, error) {
	if err := r.db.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) List(limit, offset int) (*[]models.Book, error) {
	var books *[]models.Book
	if err := r.db.Limit(limit).Offset(offset).Find(&books).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	if len(*books) == 0 {
		return nil, fmt.Errorf("No book error")
	}
	return books, nil
}

func (r *BookRepository) ListByID(id int) (*models.Book, error) {
	var book *models.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) UpdateByID(book *models.Book, updates map[string]interface{}) error {
	return r.db.Model(book).Updates(updates).Error
}

func (r *BookRepository) DeleteByID(book *models.Book) error {
	return r.db.Delete(&book).Error
}

func (r *BookRepository) SearchByAuthor(author string) (*[]models.Book, error) {
	var book *[]models.Book
	if err := r.db.Where("author = ?", author).Find(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) SearchByName(name string) (*[]models.Book, error) {
	var book *[]models.Book
	if err := r.db.Where("name = ?", name).Find(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) Search(name, author string, limit, offset int) (*[]models.Book, error) {
	var book *[]models.Book
	if err := r.db.Limit(limit).Offset(offset).Where("name = ? AND author = ?", name, author).Find(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}
