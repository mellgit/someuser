package factory

import (
	"fmt"
	"github.com/mellgit/someuser/internal/config"

	//"github.com/mellgit/someuser/internal/repository/mongo"
	"github.com/mellgit/someuser/internal/repository"
	"github.com/mellgit/someuser/internal/repository/postgres"
)

func NewRepository(envCfg config.EnvConfig) (repository.Repository, error) {
	switch envCfg.DBType {
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%v port=%v dbname=%v user=%v password=%v sslmode=disable",
			envCfg.DBHost, envCfg.DBPort, envCfg.DBName, envCfg.DBUser, envCfg.DBPassword,
		)
		return postgres.NewPostgresRepository(dsn)
	//case "mongodb":
	//	// Разделите config на URI, dbName и collectionName
	//	return mongo.NewMongoRepository(config, "dbName", "collectionName")
	default:
		return nil, fmt.Errorf("unsupported database type: %s", envCfg.DBType)
	}
}
