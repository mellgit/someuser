package mongo

import (
	"errors"
	"fmt"
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
	//user.ID = resID.(bson.ObjectID).Hex()

	return &user, nil
}

func (m MongoRepository) GetAllUsers(ctx context.Context) (*[]model.SchemaSomeUser, error) {

	res, err := m.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to find all users: %w", err)
	}
	var users []model.SchemaSomeUser
	err = res.All(ctx, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to find all users: %w", err)
	}
	return &users, nil

}

func (m MongoRepository) GetUserByID(ctx context.Context, id string) (*model.SchemaSomeUser, error) {

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %w", err)
	}

	filter := bson.M{"_id": objectID}
	var user model.SchemaSomeUser
	if err := m.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}

	return &user, nil
}

func (m MongoRepository) DeleteUser(ctx context.Context, id string) error {

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id format: %w", err)
	}
	filter := bson.M{"_id": objectID}
	_, err = m.collection.DeleteOne(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("user not found: %w", err)
		}
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (m MongoRepository) UpdateUser(ctx context.Context, id string, request model.UpdateUserRequest) (*model.SchemaSomeUser, error) {

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %w", err)
	}
	var user model.SchemaSomeUser
	_, err = m.collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": request})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	filter := bson.M{"_id": objectID}
	err = m.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find inserted document: %w", err)
	}

	return &user, nil

}
