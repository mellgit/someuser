package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/repository"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(dsn string) (repository.Repository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) CreateUser(ctx context.Context, request model.CreateUserRequest) error {
	// Реализация CREATE
	return nil
}

func (r *PostgresRepository) GetAllUsers(ctx context.Context) error {
	// Реализация READ
	return nil
}

func (r *PostgresRepository) GetUserByID(ctx context.Context, id uuid.UUID) error {
	// Реализация UPDATE
	return nil
}

func (r *PostgresRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	// Реализация DELETE
	return nil
}

func (r *PostgresRepository) UpdateUser(ctx context.Context, request model.UpdateUserRequest) error {
	// Реализация DELETE
	return nil
}
