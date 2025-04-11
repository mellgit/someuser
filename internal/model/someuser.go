package model

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
	// TODO uuid for postgres and primitive.ObjectID for mongodb
	ID       any    `bson:"_id" db:"id" json:"_id,omitempty"`
	Username string `bson:"username" db:"username" json:"username"`
	Email    string `bson:"email" db:"email" json:"email"`
	Password string `bson:"password" db:"password" json:"password"`
}
