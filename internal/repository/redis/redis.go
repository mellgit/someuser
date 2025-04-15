package redis

import (
	"fmt"
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
		Password: "",             // no password set
		DB:       envCfg.RedisDB, // use default DB
	})

	fmt.Println("Redis Repository Initialized")
	return &RedisRepository{client: client}, nil
}

func (r RedisRepository) CreateUser(ctx context.Context, request model.CreateUserRequest) (*model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
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
