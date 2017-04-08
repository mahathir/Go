package utilities

import (
	"errors"
	"log"

	"github.com/olebedev/config"
)

type (
	// AppConfig . Application configuration.
	AppConfig struct {
		MongoDbConn string
	}
)

// NewAppConfig . Return initialized app config
func NewAppConfig() *AppConfig {
	config := &AppConfig{}
	var err = errors.New("")
	cfg := readConfigFile()
	config.MongoDbConn, err = cfg.String("mongoconn")
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func readConfigFile() *config.Config {
	cfg, err := config.ParseJsonFile("appconfig.json")
	if err != nil {
		log.Fatal(err)
	}
	runconfig, err := cfg.String("config")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err = cfg.Get(runconfig)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
