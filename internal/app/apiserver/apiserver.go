package apiserver

import (
	"net/http"
	"uszn-go-test/internal/app/logger"
	"uszn-go-test/internal/app/router"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logger.Logger
	router *router.Router
}

// New ...
func New(config *Config, logger *logger.Logger, router *router.Router) *APIServer {
	return &APIServer{
		config: config,
		logger: logger,
		router: router,
	}
}

// Start ...
func (s *APIServer) Start() error {
	s.logger.Logger.Info("Starting API server at port " + s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router.Router)
}