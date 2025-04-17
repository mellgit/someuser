package factory

import (
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/repository"
	"github.com/mellgit/someuser/internal/service"
	"github.com/mellgit/someuser/internal/service/someuser"
)

func NewService(cfg *config.Config, repo repository.Repository) service.Service {
	return someuser.NewSomeUserService(cfg, repo)
}
