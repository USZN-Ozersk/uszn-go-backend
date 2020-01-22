package main

import (
	"flag"
	"log"
	"uszn-go-test/internal/app/apiserver"
	"uszn-go-test/internal/app/config"
	"uszn-go-test/internal/app/logger"
	"uszn-go-test/internal/app/router"
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

	router := router.New(logger)
	router.ConfigureRouter()

	apiserver := apiserver.New(config, logger, router)
	if err := apiserver.Start(); err != nil {
		logger.Logger.Fatal(err)
	}
}
