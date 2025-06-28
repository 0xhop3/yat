package models

import (
	"github.com/google/uuid"
	"time"
)

type TaskList struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateTaskListRequest struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"required,max=100"`
	Description *string   `json:"description" validate:"max=200"`
}

type UpdateTaskListRequest struct {
	Name        string  `json:"name" validate:"max=100"`
	Description *string `json:"description" validate:"max=200"`
}
