package service

import (
	"context"
	"fmt"
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

func (s *SomeUser) CreateUser(payload model.CreateUserRequest) (*model.SchemaSomeUser, error) {
	user, err := s.repo.CreateUser(context.Background(), payload)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return user, nil
}

func (s *SomeUser) GetAllUsers() (*[]model.SchemaSomeUser, error) {

	allUsers, err := s.repo.GetAllUsers(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}
	return allUsers, nil
}

func (s *SomeUser) GetUserByID(id uuid.UUID) (*model.SchemaSomeUser, error) {

	user, err := s.repo.GetUserByID(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	return user, nil
}
func (s *SomeUser) DeleteUserByID(id uuid.UUID) error {

	if err := s.repo.DeleteUser(context.Background(), id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}

func (s *SomeUser) UpdateUser(payload model.UpdateUserRequest) error {
	return nil
}
