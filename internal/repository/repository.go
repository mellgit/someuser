package repository

import (
	"context"
	"github.com/mellgit/someuser/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, request model.CreateUserRequest) (*model.SchemaSomeUser, error)
	GetAllUsers(ctx context.Context) (*[]model.SchemaSomeUser, error)
	GetUserByID(ctx context.Context, id any) (*model.SchemaSomeUser, error)
	DeleteUser(ctx context.Context, id any) error
	UpdateUser(ctx context.Context, id any, request model.UpdateUserRequest) (*model.SchemaSomeUser, error)
}
