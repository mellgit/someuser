package cmd

import (
	"flag"
	"fmt"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/pkg/logger"

	log "github.com/sirupsen/logrus"
	"time"
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

	cnt := 0
	for {
		fmt.Printf("count is %d\n", cnt)
		cnt++
		time.Sleep(time.Second)
	}
}
