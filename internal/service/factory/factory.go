package factory

import (
	"fmt"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/repository"
	"github.com/mellgit/someuser/internal/service"
	"github.com/mellgit/someuser/internal/service/someuser"
)

func NewService(cfg *config.Config, repo repository.Repository) (service.Service, error) {
	switch cfg.Service {
	case "someuser":
		return someuser.NewSomeUserService(cfg, repo), nil
	default:
		return nil, fmt.Errorf("unknown service %v", cfg.Service)
	}
}
