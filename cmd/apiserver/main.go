package main

import (
	"flag"
	"log"
	"uszn-go-backend/internal/app/apiserver"
	"uszn-go-backend/internal/app/config"
	"uszn-go-backend/internal/app/logger"
	"uszn-go-backend/internal/app/router"
	"uszn-go-backend/internal/app/store"
)

var config_path string

func init() {
	flag.StringVar(&config_path, "config-path", "config/config.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := config.New()
	if err := config.InitConfig(config_path); err != nil {
		log.Fatal(err)
	}

	logger := logger.New(config)
	if err := logger.InitLogger(); err != nil {
		log.Fatal(err)
	}

	logger.Logger.Info("Logger module initialised")

	store := store.New(config, logger)
	if err := store.Open(); err != nil {
		logger.Logger.Fatal(err)
	}

	defer store.Close()

	router := router.New(logger)
	router.ConfigureRouter()

	apiserver := apiserver.New(config, logger, router)
	if err := apiserver.Start(); err != nil {
		logger.Logger.Fatal(err)
	}
}
