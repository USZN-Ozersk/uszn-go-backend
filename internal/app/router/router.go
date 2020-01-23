package router

import (
	"io"
	"net/http"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/logger"

	"github.com/gorilla/mux"
)

// Router ...
type Router struct {
	Router *mux.Router
	logger *logger.Logger
}

// New ...
func New(logger *logger.Logger) *Router {
	return &Router{
		Router: mux.NewRouter(),
		logger: logger,
	}
}

// ConfigureRouter ...
func (r *Router) ConfigureRouter() {
	r.Router.HandleFunc("/hello", r.handleHello())
	r.logger.Logger.Info("Handlers configuration complete")
}

func (r *Router) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
