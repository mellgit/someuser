package factory

import (
	"fmt"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/service"
	"github.com/mellgit/someuser/internal/transport"
	"github.com/mellgit/someuser/internal/transport/http/handler"
	log "github.com/sirupsen/logrus"
)

func NewTransport(cfg *config.Config, service service.Service, logger *log.Entry) (transport.Transport, error) {
	switch cfg.Transport {
	case "http":
		return handler.NewSomeUser(cfg, service, logger), nil
	default:
		return nil, fmt.Errorf("unknown transport type: %s", cfg.Transport)
	}

}
