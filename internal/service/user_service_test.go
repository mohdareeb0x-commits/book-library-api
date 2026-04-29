package service

import (
	"errors"
	"testing"

	"github.com/mohdareeb0x-commits/book-library-api/internal/dto"
	"github.com/mohdareeb0x-commits/book-library-api/internal/models"
	"github.com/mohdareeb0x-commits/book-library-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestRegisterSuccess(t *testing.T) {
	mockRepo := &repository.MockUserRepository{
		GetByUserNameFunc: func(username string) (*models.User, error) {
			return nil, gorm.ErrRecordNotFound
		},
		CreateFunc: func(user *models.User) (*models.User, error) {
			user.ID = 1
			return user, nil
		},
	}
	service := NewAuthService(mockRepo)

	input := dto.RegisterInput{
		UserName: "areeb",
		Password: "pass123",
	}

	_, err := service.CreateUser(input)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestRegisterFail(t *testing.T) {
	mockRepo := &repository.MockUserRepository{
		GetByUserNameFunc: func(username string) (*models.User, error) {
			return &models.User{UserName: username}, nil
		},
		CreateFunc: func(user *models.User) (*models.User, error) {
			return nil, nil
		},
	}
	service := NewAuthService(mockRepo)

	input := dto.RegisterInput{
		UserName: "areeb",
		Password: "pass123",
	}

	_, err := service.CreateUser(input)
	if err == nil {
		t.Fatal("expected error, got non")
	}
}

func TestLoginSuccess(t *testing.T) {
	input := dto.RegisterInput{
		UserName: "areeb",
		Password: "pass123",
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	
	mockRepo := &repository.MockUserRepository{
		GetByUserNameFunc: func(username string) (*models.User, error) {
			return &models.User{ID: 1, UserName: username, Password: string(hashedPassword), Role: "user"}, nil
		},
	}
	
	service := NewAuthService(mockRepo)

	_, _, err := service.Login(input) 
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestLoginFail(t *testing.T) {
	input := dto.RegisterInput{
		UserName: "areeb",
		Password: "pass123",
	}

	mockRepo := &repository.MockUserRepository{
		GetByUserNameFunc: func(username string) (*models.User, error) {
			return nil, errors.New("test error")
		},
	}
	
	service := NewAuthService(mockRepo)

	_, _, err := service.Login(input) 
	if err == nil {
		t.Fatal("expected error, got none")
	}
}