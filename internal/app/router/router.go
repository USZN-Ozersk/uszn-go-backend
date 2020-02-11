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
	r.Router.HandleFunc("/api/v1/news/{type}/{count}", r.handleGetNews()).Methods("GET")
	r.logger.Logger.Info("Handlers configuration complete")
}

func (r *Router) handleGetMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "86400")
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "86400")
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "86400")
		params := mux.Vars(q)

		if params["type"] == "single" {
			news, err := repos.GetOneNews(r.store, params["count"])
			if err != nil {
				r.logger.Logger.Error(err)
				r.error(w, q, http.StatusBadRequest, err)
				return
			}
			r.respond(w, q, http.StatusOK, news)
		}

		if params["type"] == "first" {
			news, err := repos.GetFirstNews(r.store, params["count"])
			if err != nil {
				r.logger.Logger.Error(err)
				r.error(w, q, http.StatusBadRequest, err)
				return
			}
			r.respond(w, q, http.StatusOK, news)
		}

		if params["type"] == "count" {
			count, err := repos.CountNews(r.store)
			if err != nil {
				r.logger.Logger.Error(err)
				r.error(w, q, http.StatusBadRequest, err)
				return
			}
			result := map[string]int{"count": count}
			r.respond(w, q, http.StatusOK, result)
		}

		if params["type"] == "page" {
			news, err := repos.GetPageOfNews(r.store, params["count"])
			if err != nil {
				r.logger.Logger.Error(err)
				r.error(w, q, http.StatusBadRequest, err)
				return
			}
			r.respond(w, q, http.StatusOK, news)
		}
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
