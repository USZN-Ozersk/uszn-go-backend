package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/auth"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/logger"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/repos"
	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/store"

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
	r.Router.Use(r.setHeader)
	r.Router.HandleFunc("/api/v1/menu", r.handleGetMenu()).Methods("GET")
	r.Router.HandleFunc("/api/v1/page/{id}", r.handleGetPage()).Methods("GET")
	r.Router.HandleFunc("/api/v1/news/{type}/{count}", r.handleGetNews()).Methods("GET")
	r.Router.HandleFunc("/api/v1/auth", r.handleAuth()).Methods("POST")

	private := r.Router.PathPrefix("/api/v1/private").Subrouter()
	private.Use(r.authenticateUser)
	private.HandleFunc("/pages", r.handleGetAllPages()).Methods("GET")
	private.HandleFunc("/menu", r.handleGetMenu()).Methods("GET")
	r.logger.Logger.Info("Handlers configuration complete")
}

func (r *Router) setHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, q *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		next.ServeHTTP(w, q)
	})
}

func (r *Router) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, q *http.Request) {
		if q.Header["Token"] != nil {
			if auth.UserAuthenticate(r.store, q.Header["Token"][0]) {
				next.ServeHTTP(w, q)
				return
			}
		}
		var errUnauthorized = errors.New("Unauthorised")
		r.logger.Logger.Error("Incorrect token")
		r.error(w, q, http.StatusUnauthorized, errUnauthorized)
	})
}

func (r *Router) handleGetMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		menu, err := repos.GetMenus(r.store)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusInternalServerError, err)
			return
		}
		r.respond(w, q, http.StatusOK, menu)
	}
}

func (r *Router) handleGetAllPages() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		pages, err := repos.GetAllPages(r.store)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusBadRequest, err)
			return
		}
		r.respond(w, q, http.StatusOK, pages)
	}
}

func (r *Router) handleGetPage() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
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

func (r *Router) handleAuth() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, q *http.Request) {
		req := &request{}
		if err := json.NewDecoder(q.Body).Decode(req); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusBadRequest, err)
			return
		}

		token, err := auth.UserAuthorize(r.store, req.Login, req.Password)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusUnauthorized, err)
			return
		}
		result := map[string]string{"jwt": token}
		r.respond(w, q, http.StatusOK, result)
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
