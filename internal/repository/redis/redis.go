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

	fmt.Println("Redis Repository Initialized")
	return &RedisRepository{client: client}, nil
}

func (r RedisRepository) CreateUser(ctx context.Context, request model.CreateUserRequest) (*model.SchemaSomeUser, error) {

	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshal json: %w", err)
	}
	id, _ := uuid.NewV7()
	temp := r.client.Set(ctx, id.String(), string(data), 0).Val()
	fmt.Println(temp)

	val, err := r.client.Get(ctx, id.String()).Result()
	if err != nil {
		return nil, fmt.Errorf("get value: %w", err)
	}

	var user model.SchemaSomeUser
	if err = json.Unmarshal([]byte(val), &user); err != nil {
		return nil, fmt.Errorf("unmarshal json: %w", err)
	}
	user.ID = id.String()
	return &user, nil
}

func (r RedisRepository) GetAllUsers(ctx context.Context) (*[]model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisRepository) GetUserByID(ctx context.Context, id any) (*model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisRepository) DeleteUser(ctx context.Context, id any) error {
	//TODO implement me
	panic("implement me")
}

func (r RedisRepository) UpdateUser(ctx context.Context, id any, request model.UpdateUserRequest) (*model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
}
