package service

import (
	"github.com/google/uuid"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/model"
)

type SomeUser struct {
	Cfg *config.Config
	//repo
}

func NewSomeUser(cfg *config.Config) *SomeUser {
	return &SomeUser{
		Cfg: cfg,
	}
}

func (s *SomeUser) CreateUser(payload model.CreateUserRequest) error {

	return nil
}

func (s *SomeUser) GetAllUsers() error {
	return nil
}

func (s *SomeUser) GetUserByID(id uuid.UUID) error {
	return nil
}
func (s *SomeUser) DeleteUserByID(id uuid.UUID) error {
	return nil
}

func (s *SomeUser) UpdateUser(payload model.CreateUserRequest) error {
	return nil
}
