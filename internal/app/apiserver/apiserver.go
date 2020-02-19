package apiserver

import (
	"net/http"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/config"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/logger"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/router"
)

// APIServer ...
type APIServer struct {
	config *config.Config
	logger *logger.Logger
	router *router.Router
}

// New ...
func New(config *config.Config, logger *logger.Logger, router *router.Router) *APIServer {
	return &APIServer{
		config: config,
		logger: logger,
		router: router,
	}
}

// Start ...
func (s *APIServer) Start() error {
	s.logger.Logger.Info("Starting API server at port " + s.config.BindAddr)

	return http.ListenAndServeTLS(s.config.BindAddr, "/etc/letsencrypt/live/usznozersk.ru/fullchain.pem", "/etc/letsencrypt/live/usznozersk.ru/privkey.pem", s.router.Router)
}
