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
	r.Router.HandleFunc("/api/v1/menu", r.handleGetMenu()).Methods("GET")
	r.Router.HandleFunc("/api/v1/page/{id}", r.handleGetPage()).Methods("GET")
	r.Router.HandleFunc("/api/v1/news", r.handleGetNews()).Methods("GET")
	r.Router.HandleFunc("/api/v1/news/{id}", r.handleGetOneNews()).Methods("GET")
	r.logger.Logger.Info("Handlers configuration complete")
}

func (r *Router) handleGetMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		menu, err := repos.GetMenus(r.store)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusInternalServerError, err)
			return
		}
		r.respond(w, q, http.StatusOK, menu)
	}
}

func (r *Router) handleGetPage() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(q)
		page, err := repos.GetPage(r.store, params["id"])
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusBadRequest, err)
			return
		}
		r.respond(w, q, http.StatusOK, page)
	}
}

func (r *Router) handleGetNews() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		news, err := repos.GetNews(r.store)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusBadRequest, err)
			return
		}
		r.respond(w, q, http.StatusOK, news)
	}
}

func (r *Router) handleGetOneNews() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(q)
		news, err := repos.GetOneNews(r.store, params["id"])
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusBadRequest, err)
			return
		}
		r.respond(w, q, http.StatusOK, news)
	}
}

func (r *Router) error(w http.ResponseWriter, q *http.Request, code int, err error) {
	r.respond(w, q, code, map[string]string{"error": err.Error()})
}

func (r *Router) respond(w http.ResponseWriter, q *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
