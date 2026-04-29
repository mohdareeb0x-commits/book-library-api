package repository

import (
	"github.com/mohdareeb0x-commits/book-library-api/internal/models"

	"gorm.io/gorm"
)



type UserRepositoryInterface interface {
	GetByUserName(username string) (*models.User, error) 
	Create(user *models.User) (*models.User, error) 
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByUserName(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("user_name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
