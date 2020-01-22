package main

import (
	"log"
	"uszn-go-test/internal/app/apiserver"
	"uszn-go-test/internal/app/logger"
	"uszn-go-test/internal/app/router"
)

func main() {
	logger := logger.New(logger.NewConfig())
	if err := logger.InitLogger(); err != nil {
		log.Fatal(err)
	}

	logger.Logger.Info("Logger module initialised")

	router := router.New(logger)
	router.ConfigureRouter()

	apiserver := apiserver.New(apiserver.NewConfig(), logger, router)
	if err := apiserver.Start(); err != nil {
		logger.Logger.Fatal(err)
	}
}
