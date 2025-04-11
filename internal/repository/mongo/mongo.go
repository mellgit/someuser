package mongo

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
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

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", envCfg.MongoUser, envCfg.MongoPassword, envCfg.MongoHost, envCfg.MongoPort)
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("mongo connect err: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("mongo ping err: %v", err)
	}

	collection := client.Database(envCfg.MongoDB).Collection(envCfg.MongoCollection)

	return &MongoRepository{
		collection: collection,
	}, nil

}

func (m MongoRepository) CreateUser(ctx context.Context, request model.CreateUserRequest) (*model.SchemaSomeUser, error) {

	res, err := m.collection.InsertOne(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	resID := res.InsertedID
	var user model.SchemaSomeUser
	filter := bson.M{"_id": resID}
	err = m.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find inserted document: %w", err)
	}
	user.ID = resID.(bson.ObjectID).Hex()

	return &user, nil
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
