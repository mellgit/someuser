package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/repository"
	"github.com/pressly/goose/v3"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(envCfg config.EnvConfig) (repository.Repository, error) {
	dsn := fmt.Sprintf(
		"host=%v port=%v dbname=%v user=%v password=%v sslmode=disable",
		envCfg.DBHost, envCfg.DBPort, envCfg.DBName, envCfg.DBUser, envCfg.DBPassword,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}
	if err = goose.Up(db, envCfg.MigrationsPath); err != nil {
		return nil, fmt.Errorf("failed to create migrations: %w", err)
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
	var tempID uuid.UUID
	err := r.db.QueryRowContext(ctx, query, request.Username, request.Email, request.Password).Scan(&tempID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	user.ID = tempID
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
	var tempID uuid.UUID
	for rows.Next() {
		var user model.SchemaSomeUser
		if err := rows.Scan(&tempID, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, fmt.Errorf("failed to get all users: %w", err)
		}
		user.ID = tempID
		users = append(users, user)
	}
	return &users, nil
}

func (r *PostgresRepository) GetUserByID(ctx context.Context, id any) (*model.SchemaSomeUser, error) {

	uuID, err := uuid.Parse(id.(string))
	if err != nil {
		return nil, fmt.Errorf("failed to parse user id: %w", err)
	}
	query := `SELECT * FROM someusers WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, uuID)
	var user model.SchemaSomeUser
	var tempID uuid.UUID
	if err := row.Scan(&tempID, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	user.ID = tempID
	return &user, nil
}

func (r *PostgresRepository) DeleteUser(ctx context.Context, id any) error {

	uuID, err := uuid.Parse(id.(string))
	if err != nil {
		return fmt.Errorf("failed to parse user id: %w", err)
	}
	query := `DELETE FROM someusers WHERE id = $1`
	_, err = r.db.ExecContext(ctx, query, uuID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (r *PostgresRepository) UpdateUser(ctx context.Context, id any, request model.UpdateUserRequest) (*model.SchemaSomeUser, error) {

	uuID, err := uuid.Parse(id.(string))
	if err != nil {
		return nil, fmt.Errorf("failed to parse user id: %w", err)
	}
	query := `
	update someusers
	set username=$1, email=$2, password=$3
	where id=$4
	returning *`

	var user model.SchemaSomeUser
	var tempID uuid.UUID
	err = r.db.QueryRowContext(ctx, query, request.Username, request.Email, request.Password, uuID.String()).Scan(&tempID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	user.ID = tempID
	return &user, nil
}
