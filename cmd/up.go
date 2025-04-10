package cmd

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/service"
	"github.com/mellgit/someuser/internal/transport/http/handler"
	"github.com/mellgit/someuser/pkg/logger"
	log "github.com/sirupsen/logrus"
)

func Up() {

	cfgPath := flag.String("config", "config.yml", "config file path")
	flag.Parse()

	cfg, err := config.LoadConfig(*cfgPath)
	if err != nil {
		log.WithFields(log.Fields{
			"action": "config.LoadConfig",
		}).Fatal(err)
	}
	envCfg, err := config.LoadEnvConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"action": "config.LoadEnvConfig",
		}).Fatal(err)
	}

	if err = logger.SetUpLogger(*cfg); err != nil {
		log.WithFields(log.Fields{
			"action": "logger.SetUpLogger",
		}).Fatal(err)
	}

	log.Debugf("config: %+v", cfg)
	log.Debugf("env: %+v", envCfg)

	app := fiber.New()
	{
		someUserService := service.NewSomeUser(cfg)
		someUserHandler := handler.NewSomeUser(cfg, someUserService, log.WithFields(log.Fields{"service": "SomeUser"}))
		someUserHandler.Register(app)
	}
	log.WithFields(log.Fields{
		"action": "app.Listen",
	}).Fatal(app.Listen(fmt.Sprintf(":%v", envCfg.APIPort)))

}
