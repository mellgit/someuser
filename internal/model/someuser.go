package model

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SchemaSomeUser struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"password"`
}
