package mongo

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"golang.org/x/net/context"
	"time"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(envCfg config.EnvConfig) (repository.Repository, error) {

	uri := fmt.Sprintf("mongodb://%s:%v", envCfg.MongoHost, envCfg.MongoPort)
	client, _ := mongo.Connect(options.Client().ApplyURI(uri))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_ = client.Ping(ctx, readpref.Primary())

	collection := client.Database(envCfg.MongoDB).Collection(envCfg.MongoCollection)

	return &MongoRepository{
		collection: collection,
	}, nil

}

func (m MongoRepository) CreateUser(ctx context.Context, request model.CreateUserRequest) (*model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
}

func (m MongoRepository) GetAllUsers(ctx context.Context) (*[]model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
}

func (m MongoRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
}

func (m MongoRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (m MongoRepository) UpdateUser(ctx context.Context, id uuid.UUID, request model.UpdateUserRequest) (*model.SchemaSomeUser, error) {
	//TODO implement me
	panic("implement me")
}
