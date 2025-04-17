package someuser

import (
	"context"
	"fmt"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/repository"
	"github.com/mellgit/someuser/internal/service"
)

type SomeUserService struct {
	Cfg  *config.Config
	repo repository.Repository
}

func NewSomeUserService(cfg *config.Config, repo repository.Repository) service.Service {
	return &SomeUserService{
		Cfg:  cfg,
		repo: repo,
	}
}

func (s *SomeUserService) CreateUser(payload model.CreateUserRequest) (*model.SchemaSomeUser, error) {
	user, err := s.repo.CreateUser(context.Background(), payload)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return user, nil
}

func (s *SomeUserService) GetAllUsers() (*[]model.SchemaSomeUser, error) {

	allUsers, err := s.repo.GetAllUsers(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}
	return allUsers, nil
}

func (s *SomeUserService) GetUserByID(id string) (*model.SchemaSomeUser, error) {

	user, err := s.repo.GetUserByID(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	return user, nil
}
func (s *SomeUserService) DeleteUserByID(id string) error {

	if err := s.repo.DeleteUser(context.Background(), id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}

func (s *SomeUserService) UpdateUser(id string, payload model.UpdateUserRequest) (*model.SchemaSomeUser, error) {

	user, err := s.repo.UpdateUser(context.Background(), id, payload)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	return user, nil
}
