package router

import (
	"net/http"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/repos"

	"github.com/gorilla/mux"
)

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
