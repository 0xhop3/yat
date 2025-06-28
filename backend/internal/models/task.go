package models

import (
	"github.com/google/uuid"
	"time"
)

type TaskPriority string

const (
	TaskPriorityLow    TaskPriority = "low"
	TaskPriorityMedium TaskPriority = "medium"
	TaskPriorityHigh   TaskPriority = "high"
)

type Task struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	TaskListID  uuid.UUID    `json:"task_list_id" db:"task_list_id"`
	Title       string       `json:"title" db:"title"`
	IsCompleted bool         `json:"is_completed" db:"is_completed"`
	Priority    TaskPriority `json:"priority" db:"priority"`
	DueDate     *time.Time   `json:"due_date" db:"due_data"`
	CompletedAt *time.Time   `json:"completed_at" db:"completed_at"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

type CreateTaskRequest struct {
	TaskListID uuid.UUID    `json:"task_list_id" validate:"required"`
	Title      string       `json:"title" validate:"required,max=100"`
	Priority   TaskPriority `json:"priority" validate:"oneof=low medium high"`
	DueDate    *time.Time   `json:"due_date"`
}

type UpdateTaskRequest struct {
	Title       string       `json:"title"`
	IsCompleted bool         `json:"is_completed"`
	DueDate     *time.Time   `json:"due_date"`
	Priority    TaskPriority `json:"priority"`
}
