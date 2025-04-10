package service

import (
	"github.com/google/uuid"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/repository"
)

type SomeUser struct {
	Cfg  *config.Config
	repo repository.Repository
}

func NewSomeUser(cfg *config.Config, repo repository.Repository) *SomeUser {
	return &SomeUser{
		Cfg:  cfg,
		repo: repo,
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

func (s *SomeUser) UpdateUser(payload model.UpdateUserRequest) error {
	return nil
}
