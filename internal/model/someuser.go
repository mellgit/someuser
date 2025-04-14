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
	ID       any    `db:"id" bson:"_id" json:"id,omitempty"`
	Username string `db:"username" bson:"username" json:"username"`
	Email    string `db:"email" bson:"email" json:"email"`
	Password string `db:"password" bson:"password" json:"password"`
}
