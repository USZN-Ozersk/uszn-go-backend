package router

import (
	"net/http"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/repos"

	"github.com/gorilla/mux"
)

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
