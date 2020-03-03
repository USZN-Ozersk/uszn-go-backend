package router

import (
	"encoding/json"
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
func (r *Router) handleInsertPage() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
		Text string `json:"text"`
		Menu int    `json:"menu"`
	}

	return func(w http.ResponseWriter, q *http.Request) {
		req := &request{}
		if err := json.NewDecoder(q.Body).Decode(req); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusBadRequest, err)
			return
		}

		err := repos.InsertPage(r.store, req.Name, req.Text, req.Menu)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusInternalServerError, err)
			return
		}
		result := map[string]string{"result": "ok"}

		r.respond(w, q, http.StatusOK, result)
	}
}

func (r *Router) handleDeletePage() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, q *http.Request) {
		req := &request{}
		if err := json.NewDecoder(q.Body).Decode(req); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusBadRequest, err)
			return
		}

		err := repos.DeletePage(r.store, req.ID)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusInternalServerError, err)
			return
		}
		result := map[string]string{"result": "ok"}

		r.respond(w, q, http.StatusOK, result)
	}
}

func (r *Router) handleUpdatePage() http.HandlerFunc {
	type request struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Text string `json:"text"`
		Menu int    `json:"menu"`
	}

	return func(w http.ResponseWriter, q *http.Request) {
		req := &request{}
		if err := json.NewDecoder(q.Body).Decode(req); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusBadRequest, err)
			return
		}

		err := repos.UpdatePage(r.store, req.ID, req.Name, req.Text, req.Menu)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusInternalServerError, err)
			return
		}
		result := map[string]string{"result": "ok"}

		r.respond(w, q, http.StatusOK, result)
	}
}
