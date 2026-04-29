package repository

import (
	"github.com/mohdareeb0x-commits/book-library-api/internal/models"
)

type MockUserRepository struct {
	GetByUserNameFunc func(username string) (*models.User, error)
	CreateFunc func(user *models.User) (*models.User, error) 
}

func (m *MockUserRepository) GetByUserName(username string) (*models.User, error) {
	return m.GetByUserNameFunc(username)
}

func (m *MockUserRepository) Create(user *models.User) (*models.User, error) {
	return m.CreateFunc(user)
}
