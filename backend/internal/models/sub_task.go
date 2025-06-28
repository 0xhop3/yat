package models

import (
	"github.com/google/uuid"
	"time"
)

type SubTask struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	ParentTaskID uuid.UUID  `json:"parent_task_id" db:"parent_task_id"`
	Title        string     `json:"title" db:"title"`
	IsCompleted  bool       `json:"is_completed" db:"is_completed"`
	CompletedAt  *time.Time `json:"completed_at" db:"completed_at"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateSubTaskRequest struct {
	Title        string    `json:"title" validate:"required,max=100"`
	ParentTaskID uuid.UUID `json:"parent_task_id" validate:"required"`
}

type UpdateSubTaskRequest struct {
	Title string `json:"title" validate:"required,max=100"`
}
