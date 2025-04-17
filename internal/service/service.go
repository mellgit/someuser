package service

import (
	"github.com/mellgit/someuser/internal/model"
)

type Service interface {
	CreateUser(payload model.CreateUserRequest) (*model.SchemaSomeUser, error)
	GetAllUsers() (*[]model.SchemaSomeUser, error)
	GetUserByID(id string) (*model.SchemaSomeUser, error)
	DeleteUserByID(id string) error
	UpdateUser(id string, payload model.UpdateUserRequest) (*model.SchemaSomeUser, error)
}
