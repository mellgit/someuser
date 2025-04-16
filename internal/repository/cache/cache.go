package cache

import (
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/repository"
	"golang.org/x/net/context"
)

type CacheRepository struct {
	cache map[string]model.CreateUserRequest
}

func NewCacheRepository() (repository.Repository, error) {
	res := make(map[string]model.CreateUserRequest)
	return &CacheRepository{
		cache: res,
	}, nil
}

func (c CacheRepository) CreateUser(ctx context.Context, request model.CreateUserRequest) (*model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
}

func (c CacheRepository) GetAllUsers(ctx context.Context) (*[]model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
}

func (c CacheRepository) GetUserByID(ctx context.Context, id any) (*model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
}

func (c CacheRepository) DeleteUser(ctx context.Context, id any) error {
	//TODO implement me
	panic("implement me")
}

func (c CacheRepository) UpdateUser(ctx context.Context, id any, request model.UpdateUserRequest) (*model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
}
