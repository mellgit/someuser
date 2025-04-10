package cmd

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/mellgit/someuser/docs"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/repository/factory"
	"github.com/mellgit/someuser/internal/service"
	"github.com/mellgit/someuser/internal/transport/http/handler"
	"github.com/mellgit/someuser/pkg/logger"
	log "github.com/sirupsen/logrus"
)

// Up
// @title Some Users
// @version 1.0
// @host localhost:3000
// @BasePath /api/v1
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

	repo, err := factory.NewRepository(*envCfg)
	if err != nil {
		log.WithFields(log.Fields{
			"action": "repository.NewRepository",
		}).Fatal(err)
	}
	app := fiber.New()
	{
		someUserService := service.NewSomeUser(cfg, repo)
		someUserHandler := handler.NewSomeUser(cfg, someUserService, log.WithFields(log.Fields{"service": "SomeUser"}))
		someUserHandler.Register(app)

		app.Get("/swagger/*", swagger.HandlerDefault)
	}
	log.WithFields(log.Fields{
		"action": "app.Listen",
	}).Fatal(app.Listen(fmt.Sprintf(":%v", envCfg.APIPort)))

}
