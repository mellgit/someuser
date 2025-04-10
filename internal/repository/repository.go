package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/mellgit/someuser/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, request model.CreateUserRequest) (*model.SchemaSomeUser, error)
	GetAllUsers(ctx context.Context) error
	GetUserByID(ctx context.Context, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	UpdateUser(ctx context.Context, request model.UpdateUserRequest) error
}
