package repositories

import (
	"database/sql"
	"github.com/0xhop3/yat/backend/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TaskListRepository struct {
	db *sqlx.DB
}

const (
	taskListCreateQuery = `INSERT INTO task_lists (id, name, description, created_at, updated_id) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING created_at, updated_at`
	taskListByIDQuery   = `SELECT * FROM task_lists WHERE id = $1`
	taskListDeleteQuery = `DELETE * FROM task_lists WHERE id = $1`
	taskListUpdateQuery = `UPDATE task_lists SET name = $2, description = $3, updated_at = NOW() where id = $1 RETURNING updated_at`
	taskListListQuery   = `SELECT * FROM task_lists WHERE user_id = $1 ORDER BY created_at DESC LIMIT $1 OFFSET $2`
)

func NewTaskListRepository(db *sqlx.DB) *TaskListRepository {
	return &TaskListRepository{db: db}
}

func (t *TaskListRepository) Create(taskList *models.TaskList) error {
	return t.db.QueryRow(taskListCreateQuery,
		taskList.ID,
		taskList.Name,
		taskList.Description,
		taskList.CreatedAt,
		taskList.UpdatedAt).Scan(&taskList.CreatedAt, taskList.UpdatedAt)
}

func (t *TaskListRepository) GetByID(taskListID uuid.UUID) (*models.TaskList, error) {
	taskList := &models.TaskList{}
	err := t.db.Get(taskList, taskListByIDQuery, taskListID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return taskList, err
}

func (t *TaskListRepository) Delete(taskListID uuid.UUID) error {
	_, err := t.db.Exec(taskListDeleteQuery, taskListID)
	return err
}

func (t *TaskListRepository) Update(taskList *models.TaskList) error {
	return t.db.QueryRow(taskListUpdateQuery, taskList.ID, taskList.Name).Scan(&taskList.UpdatedAt)
}

func (t *TaskListRepository) ListTaskList(userID uuid.UUID, limit, offset int) ([]*models.TaskList, error) {
	taskLists := []*models.TaskList{}

	err := t.db.Select(&taskLists, taskListListQuery, userID, limit, offset)
	return taskLists, err
}
