package router

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/logger"
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
	r.Router.HandleFunc("/api/v1/feedback", r.handleMail()).Methods("POST")

	private := r.Router.PathPrefix("/api/v1/private").Subrouter()
	private.Use(r.authenticateUser)
	private.HandleFunc("/pages", r.handleGetAllPages()).Methods("GET", "OPTIONS")
	private.HandleFunc("/menu", r.handleInsertMenu()).Methods("POST", "OPTIONS")
	private.HandleFunc("/menu", r.handleDeleteMenu()).Methods("DELETE", "OPTIONS")
	private.HandleFunc("/menu", r.handleUpdateMenu()).Methods("PUT", "OPTIONS")
	private.HandleFunc("/news", r.handleInsertNews()).Methods("POST", "OPTIONS")
	private.HandleFunc("/news", r.handleDeleteNews()).Methods("DELETE", "OPTIONS")
	private.HandleFunc("/news", r.handleUpdateNews()).Methods("PUT", "OPTIONS")
	private.HandleFunc("/page", r.handleInsertPage()).Methods("POST", "OPTIONS")
	private.HandleFunc("/page", r.handleDeletePage()).Methods("DELETE", "OPTIONS")
	private.HandleFunc("/page", r.handleUpdatePage()).Methods("PUT", "OPTIONS")
	private.HandleFunc("/upload", r.handleUpload()).Methods("POST", "OPTIONS")
	r.logger.Logger.Info("Handlers configuration complete")
}

func (r *Router) setHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, q *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, q)
	})
}

func (r *Router) handleUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		if err := q.ParseMultipartForm(32 << 20); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusBadRequest, err)
			return
		}
		file, handler, err := q.FormFile("file")
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusBadRequest, err)
			return
		}

		defer file.Close()
		current := time.Now()
		path := "/upload/" + current.Format("2006") + "/" + current.Format("01") + "/" + current.Format("02") + "/"
		r.createDirIfNotExist(path)
		f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusBadRequest, err)
			return
		}
		defer f.Close()
		if _, err := io.Copy(f, file); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusInternalServerError, err)
			return
		}
		result := map[string]string{"url": f.Name()}
		r.respond(w, http.StatusOK, result)
	}
}

func (r *Router) createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			r.logger.Logger.Error(err)
		}
	}
}

func (r *Router) error(w http.ResponseWriter, code int, err error) {
	r.respond(w, code, map[string]string{"error": err.Error()})
}

func (r *Router) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			r.logger.Logger.Error(err)
		}
	}
}
