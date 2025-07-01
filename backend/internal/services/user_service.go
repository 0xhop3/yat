package services

import (
	"fmt"

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
	existingUser, err := u.userRepo.GetByAuth0ID(request.Auth0ID)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, fmt.Errorf("%s already exists", existingUser.Auth0ID)
	}

	user := &models.User{
		ID:       uuid.New(),
		Auth0ID:  request.Auth0ID,
		Username: request.Username,
		Name:     request.Name,
	}

	err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetByAuth0ID(auth0ID string) (*models.User, error) {
	user, err := u.userRepo.GetByAuth0ID(auth0ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
