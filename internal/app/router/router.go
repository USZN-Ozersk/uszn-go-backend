package router

import (
	"fmt"
	"io"
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

func (r *Router) handleGetMenu() http.HandlerFunc {
	menuData, err := repos.GetMenus(r.store)
	if err != nil {
		r.logger.Logger.Error(err)
	}
	fmt.Println(menuData)

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
