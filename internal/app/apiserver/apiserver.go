package apiserver

import (
	"net/http"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/config"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/logger"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/router"
	"github.com/gorilla/handlers"
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

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Token"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	if s.config.UseSSL != true {
		return http.ListenAndServe(s.config.BindAddr, handlers.CORS(headers, methods, origins)(s.router.Router))
	} else {
		return http.ListenAndServeTLS(s.config.BindAddr, "/etc/letsencrypt/live/usznozersk.ru/fullchain.pem", "/etc/letsencrypt/live/usznozersk.ru/privkey.pem", s.router.Router)
	}
}
