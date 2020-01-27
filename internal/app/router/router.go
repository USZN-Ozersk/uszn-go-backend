package router

import (
	"encoding/json"
	"net/http"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/repos"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/store"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/logger"

	"github.com/gorilla/mux"
)

// Router ...
type Router struct {
	Router *mux.Router
	logger *logger.Logger
	store  *store.Store
}

// New ...
func New(logger *logger.Logger, store *store.Store) *Router {
	return &Router{
		Router: mux.NewRouter(),
		logger: logger,
		store:  store,
	}
}

// ConfigureRouter ...
func (r *Router) ConfigureRouter() {
	r.Router.HandleFunc("/api/v1/menu", r.handleGetMenu())
	r.logger.Logger.Info("Handlers configuration complete")
}

func (rt *Router) handleGetMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		menu, err := repos.GetMenus(rt.store)
		if err != nil {
			rt.logger.Logger.Error(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(menu)
	}
}
