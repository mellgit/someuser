package repository

import (
	"context"
	"github.com/mellgit/someuser/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, request model.CreateUserRequest) (*model.SchemaSomeUser, error)
	GetAllUsers(ctx context.Context) (*[]model.SchemaSomeUser, error)
	GetUserByID(ctx context.Context, id string) (*model.SchemaSomeUser, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, id string, request model.UpdateUserRequest) (*model.SchemaSomeUser, error)
}
