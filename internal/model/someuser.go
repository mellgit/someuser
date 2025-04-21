package model

type CreateUserRequest struct {
	Username string `bson:"username" json:"username" validate:"required,min=3,max=50"`
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"gte=6,lte=50"`
}

type UpdateUserRequest struct {
	Username string `bson:"username" json:"username" validate:"required,min=3,max=50"`
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"gte=6,lte=50"`
}

type SchemaSomeUser struct {
	ID       any    `db:"id" bson:"_id" json:"id,omitempty"` // validate:"-" - no validation
	Username string `db:"username" bson:"username" json:"username"`
	Email    string `db:"email" bson:"email" json:"email"`
	Password string `db:"password" bson:"password" json:"password"`
}
