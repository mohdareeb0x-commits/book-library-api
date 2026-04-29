package service

import (
	"errors"

	"github.com/mohdareeb0x-commits/book-library-api/internal/dto"
	"github.com/mohdareeb0x-commits/book-library-api/internal/models"
	"github.com/mohdareeb0x-commits/book-library-api/internal/repository"
	"github.com/mohdareeb0x-commits/book-library-api/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) CreateUser(input dto.RegisterInput) (*dto.UserResponse, error) {
	_, err := s.userRepo.GetByUserName(input.UserName)
	if err == nil {
		return nil, errors.New("user with that name already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		UserName: input.UserName,
		Password: string(hashedPassword),
	}

	newUser, err := s.userRepo.Create(&user)

	userResponse := &dto.UserResponse{
		ID:       newUser.ID,
		UserName: newUser.UserName,
	}

	return userResponse, err
}

func (s *AuthService) Login(input dto.RegisterInput) (string, error) {
	user, err := s.userRepo.GetByUserName(input.UserName)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return "", err
	}

	return token, nil
}
