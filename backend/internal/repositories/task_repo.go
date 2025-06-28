package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	taskCreateQuery         = `INSERT INTO tasks (id, task_list_id, title, priority, created_at, upadted_at, is_complete, completed_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	taskDeleteQuery         = `SELECT * FROM tasks WHERE id = $1`
	taskGetByIDQuery        = `SELECT * FROM tasks WHERE id = $1`
	taskUpdateQuery         = `UPDATE tasks SET title = $2, updated_at = NOW(), is_complete = $4, completed_at = $5 RETURNING update_at`
	taskListByTaskListQuery = `SELECT * FROM tasks where task_list_id = $1 ORDER_BY created_at DESC LIMIT $2 OFFSET $3`
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) Create(task *models.Task) error {
	return t.db.QueryRow(taskCreateQuery, task.ID, task.TaskListID, task.Title, task.Priority, task.CreatedAt, task.UpdatedAt, task.IsCompleted, task.CompletedAt).Scan(&task.CreatedAt)
}

func (t *TaskRepository) GetByID(taskID uuid.UUID) (*models.Task, error) {
	task := &models.Task{}

	err := t.db.Get(task, taskGetByIDQuery, taskID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return task, err
}

func (t *TaskRepository) Update(task *models.Task) error {
	return t.db.QueryRow(taskUpdateQuery, task.ID, task.Title).Scan(&task.UpdatedAt)
}

func (t *TaskRepository) Delete(taskID uuid.UUID) error {
	_, err := t.db.Exec(taskDeleteQuery, taskID)
	return err
}

func (t *TaskRepository) ListByTaskList(taskListID uuid.UUID, limit, offset int) ([]*models.Task, error) {
	tasks := []*models.Task{}
	err := t.db.Select(&tasks, taskListByTaskListQuery, limit, offset)
	return tasks, err
}
