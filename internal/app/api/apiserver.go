package api

import (
	"net/http"
	"uszn-go-backend/internal/app/conf"
	"uszn-go-backend/internal/app/logg"
	"uszn-go-backend/internal/app/route"
)

// APIServer ...
type APIServer struct {
	config *conf.Config
	logger *logg.Logger
	router *route.Router
}

// New ...
func New(config *conf.Config, logger *logg.Logger, router *route.Router) *APIServer {
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
