package router

import (
	"encoding/json"
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
				r.error(w, http.StatusBadRequest, err)
				return
			}
			r.respond(w, http.StatusOK, news)
		}

		if params["type"] == "first" {
			news, err := repos.GetFirstNews(r.store, params["count"])
			if err != nil {
				r.logger.Logger.Error(err)
				r.error(w, http.StatusBadRequest, err)
				return
			}
			r.respond(w, http.StatusOK, news)
		}

		if params["type"] == "count" {
			count, err := repos.CountNews(r.store)
			if err != nil {
				r.logger.Logger.Error(err)
				r.error(w, http.StatusBadRequest, err)
				return
			}
			result := map[string]int{"count": count}
			r.respond(w, http.StatusOK, result)
		}

		if params["type"] == "page" {
			news, err := repos.GetPageOfNews(r.store, params["count"])
			if err != nil {
				r.logger.Logger.Error(err)
				r.error(w, http.StatusBadRequest, err)
				return
			}
			r.respond(w, http.StatusOK, news)
		}
	}
}
func (r *Router) handleInsertNews() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
		Text string `json:"text"`
		Img  string `json:"img"`
	}

	return func(w http.ResponseWriter, q *http.Request) {
		req := &request{}
		if err := json.NewDecoder(q.Body).Decode(req); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusBadRequest, err)
			return
		}

		err := repos.InsertNews(r.store, req.Name, req.Text, req.Img)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusInternalServerError, err)
			return
		}
		result := map[string]string{"result": "ok"}

		r.respond(w, http.StatusOK, result)
	}
}

func (r *Router) handleDeleteNews() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, q *http.Request) {
		req := &request{}
		if err := json.NewDecoder(q.Body).Decode(req); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusBadRequest, err)
			return
		}

		err := repos.DeleteNews(r.store, req.ID)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusInternalServerError, err)
			return
		}
		result := map[string]string{"result": "ok"}

		r.respond(w, http.StatusOK, result)
	}
}

func (r *Router) handleUpdateNews() http.HandlerFunc {
	type request struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Text string `json:"text"`
		Img  string `json:"img"`
	}

	return func(w http.ResponseWriter, q *http.Request) {
		req := &request{}
		if err := json.NewDecoder(q.Body).Decode(req); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusBadRequest, err)
			return
		}

		err := repos.UpdateNews(r.store, req.ID, req.Name, req.Text, req.Img)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusInternalServerError, err)
			return
		}
		result := map[string]string{"result": "ok"}

		r.respond(w, http.StatusOK, result)
	}
}
