package cache

import (
	"fmt"
	"github.com/google/uuid"
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

	id, _ := uuid.NewV7()
	c.cache[id.String()] = request
	user := &model.SchemaSomeUser{
		ID:       id.String(),
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	return user, nil
}

func (c CacheRepository) GetAllUsers(ctx context.Context) (*[]model.SchemaSomeUser, error) {

	var users []model.SchemaSomeUser
	for k, v := range c.cache {
		users = append(users, model.SchemaSomeUser{
			ID:       k,
			Username: v.Username,
			Email:    v.Email,
			Password: v.Password,
		})
	}
	return &users, nil
}

func (c CacheRepository) GetUserByID(ctx context.Context, id any) (*model.SchemaSomeUser, error) {

	userCache, ok := c.cache[id.(string)]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	user := &model.SchemaSomeUser{
		ID:       id.(string),
		Username: userCache.Username,
		Email:    userCache.Email,
		Password: userCache.Password,
	}
	return user, nil
}

func (c CacheRepository) DeleteUser(ctx context.Context, id any) error {

	_, ok := c.cache[id.(string)]
	if !ok {
		return fmt.Errorf("user not found")
	}
	delete(c.cache, id.(string))
	return nil
}

func (c CacheRepository) UpdateUser(ctx context.Context, id any, request model.UpdateUserRequest) (*model.SchemaSomeUser, error) {

	_, ok := c.cache[id.(string)]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	convertRequest := model.CreateUserRequest{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	c.cache[id.(string)] = convertRequest

	user := &model.SchemaSomeUser{
		ID:       id.(string),
		Username: convertRequest.Username,
		Email:    convertRequest.Email,
		Password: convertRequest.Password,
	}
	return user, nil
}
