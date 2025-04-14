package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateUserRequest struct {
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type UpdateUserRequest struct {
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type SchemaSomeUser struct {
	// TODO uuid for postgres and bson.ObjectID for mongodb
	ID       bson.ObjectID `bson:"_id" json:"_id,omitempty"`
	Username string        `bson:"username" json:"username"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
}
