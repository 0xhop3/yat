package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Auth0ID  string    `json:"auth0_id" db:"auth0_id"`
	Username string    `json:"username" db:"username"`
	Name     string    `json:"name" db:"name"`
}

type CreateUserRequest struct {
	Auth0ID  string `json:"auth0_id" validate:"required"`
	Name     string `json:"name" validate:"max=100"`
	Username string `json:"username" validate:"required,min=3,max=50"`
}
type UpdateUserRequest struct {
	Name     string `json:"name" validate:"max=100"`
	Username string `json:"username" validate:"required,min=3,max=50"`
}
