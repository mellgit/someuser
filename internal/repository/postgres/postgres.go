package postgres

import (
	"context"
	"database/sql"
	"fmt"
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
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) CreateUser(ctx context.Context, request model.CreateUserRequest) (*model.SchemaSomeUser, error) {

	query := `
        INSERT INTO someusers (username, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, username, email, password
    `
	var user model.SchemaSomeUser
	err := r.db.QueryRowContext(ctx, query, request.Username, request.Email, request.Password).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

func (r *PostgresRepository) GetAllUsers(ctx context.Context) (*[]model.SchemaSomeUser, error) {

	query := `SELECT * FROM someusers`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	var users []model.SchemaSomeUser
	for rows.Next() {
		var user model.SchemaSomeUser
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, fmt.Errorf("failed to get all users: %w", err)
		}
		users = append(users, user)
	}
	return &users, nil
}

func (r *PostgresRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*model.SchemaSomeUser, error) {

	query := `SELECT * FROM someusers WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	var user model.SchemaSomeUser
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (r *PostgresRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {

	query := `DELETE FROM someusers WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (r *PostgresRepository) UpdateUser(ctx context.Context, id uuid.UUID, request model.UpdateUserRequest) (*model.SchemaSomeUser, error) {

	query := `
	update someusers
	set username=$1, email=$2, password=$3
	where id=$4
	returning *`

	var user model.SchemaSomeUser
	err := r.db.QueryRowContext(ctx, query, request.Username, request.Email, request.Password, id.String()).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return &user, nil
}
