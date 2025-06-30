package services

import (
	"errors"
	"github.com/0xhop3/yat/backend/internal/models"
	"github.com/0xhop3/yat/backend/internal/repositories"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) CreateUser(request *models.CreateUserRequest) (*models.User, error) {
	existingUser, err := u.userRepo.GetByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("User with this email already exists")
	}

	user := &models.User{
		ID:       uuid.New(),
		Auth0ID:  request.Auth0ID,
		Username: request.Username,
	}

	err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
