package redis

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/repository"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(envCfg config.EnvConfig) (repository.Repository, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", envCfg.RedisHost, envCfg.RedisPort),
		Password: "", // no password set
		DB:       envCfg.RedisDB,
	})

	return &RedisRepository{client: client}, nil
}

func (r RedisRepository) CreateUser(ctx context.Context, request model.CreateUserRequest) (*model.SchemaSomeUser, error) {

	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshal json: %w", err)
	}
	id, _ := uuid.NewV7()
	_ = r.client.Set(ctx, id.String(), string(data), 0).Val()

	user, err := r.GetUserByID(ctx, id.String())
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	return user, nil
}

func (r RedisRepository) GetAllUsers(ctx context.Context) (*[]model.SchemaSomeUser, error) {

	keys, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, fmt.Errorf("get keys: %w", err)
	}

	var users []model.SchemaSomeUser
	for _, key := range keys {
		user, err := r.GetUserByID(ctx, key)
		if err != nil {
			return nil, fmt.Errorf("get user: %w", err)
		}
		users = append(users, *user)
	}
	return &users, nil
}

func (r RedisRepository) GetUserByID(ctx context.Context, id any) (*model.SchemaSomeUser, error) {

	val, err := r.client.Get(ctx, id.(string)).Result()
	if err != nil {
		return nil, fmt.Errorf("get value: %w", err)
	}

	var user model.SchemaSomeUser
	if err = json.Unmarshal([]byte(val), &user); err != nil {
		return nil, fmt.Errorf("unmarshal json: %w", err)
	}
	user.ID = id.(string)
	return &user, nil
}

func (r RedisRepository) DeleteUser(ctx context.Context, id any) error {

	_, err := r.client.Del(ctx, id.(string)).Result()
	if err != nil {
		return fmt.Errorf("delete value: %w", err)
	}
	return nil
}

func (r RedisRepository) UpdateUser(ctx context.Context, id any, request model.UpdateUserRequest) (*model.SchemaSomeUser, error) {

	updateUser := model.SchemaSomeUser{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	data, err := json.Marshal(updateUser)
	if err != nil {
		return nil, fmt.Errorf("marshal json: %w", err)
	}
	_ = r.client.Set(ctx, id.(string), string(data), 0).Val()

	updateUser.ID = id.(string)
	return &updateUser, nil

}
