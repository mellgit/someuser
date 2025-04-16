package factory

import (
	"fmt"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/repository/cache"
	"github.com/mellgit/someuser/internal/repository/mongo"
	"github.com/mellgit/someuser/internal/repository/redis"

	"github.com/mellgit/someuser/internal/repository"
	"github.com/mellgit/someuser/internal/repository/postgres"
)

func NewRepository(envCfg config.EnvConfig) (repository.Repository, error) {
	switch envCfg.DBType {
	case "postgres":
		return postgres.NewPostgresRepository(envCfg)
	case "mongodb":
		return mongo.NewMongoRepository(envCfg)
	case "redis":
		return redis.NewRedisRepository(envCfg)
	case "cache":
		return cache.NewCacheRepository()
	default:
		return nil, fmt.Errorf("unsupported database type: %s", envCfg.DBType)
	}
}
