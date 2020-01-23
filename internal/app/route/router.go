package route

import (
	"io"
	"net/http"
	"uszn-go-backend/internal/app/logg"

	"github.com/gorilla/mux"
)

// Router ...
type Router struct {
	Router *mux.Router
	logger *logg.Logger
}

// New ...
func New(logger *logg.Logger) *Router {
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
