package main

import (
	"flag"
	"log"
	"uszn-go-backend/internal/app/api"
	"uszn-go-backend/internal/app/conf"
	"uszn-go-backend/internal/app/logg"
	"uszn-go-backend/internal/app/route"
)

var config_path string

func init() {
	flag.StringVar(&config_path, "config-path", "config/config.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := conf.New()
	if err := config.InitConfig(config_path); err != nil {
		log.Fatal(err)
	}

	logger := logg.New(config)
	if err := logger.InitLogger(); err != nil {
		log.Fatal(err)
	}

	logger.Logger.Info("Logger module initialised")

	router := route.New(logger)
	router.ConfigureRouter()

	apiserver := api.New(config, logger, router)
	if err := apiserver.Start(); err != nil {
		logger.Logger.Fatal(err)
	}
}
