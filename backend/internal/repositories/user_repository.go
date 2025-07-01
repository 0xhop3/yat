package repositories

import (
	"database/sql"

	"github.com/0xhop3/yat/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

const (
	userCreateQuery       = `INSERT INTO yat_users (id, auth0_id, username, name) VALUES ($1, $2, $3, $4) RETURNING created_at, updated_at`
	userGetByAuth0IDQuery = `SELECT * FROM yat_users WHERE auth0_id = $1`
)

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user *models.User) error {
	return u.db.QueryRow(userCreateQuery,
		user.ID,
		user.Auth0ID,
		user.Username,
		user.Name).Scan(&user.CreatedAt, &user.UpdatedAt)
}

func (u *UserRepository) GetByAuth0ID(auth0_id string) (*models.User, error) {
	user := &models.User{}
	err := u.db.Get(user, userGetByAuth0IDQuery, auth0_id)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, nil
}
