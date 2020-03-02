package router

import (
	"encoding/json"
	"net/http"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/repos"
)

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

func (r *Router) handleInsertMenu() http.HandlerFunc {
	type request struct {
		MenuItem   string `json:"menuitem"`
		MenuParent int    `json:"menuparent"`
	}

	return func(w http.ResponseWriter, q *http.Request) {
		req := &request{}
		if err := json.NewDecoder(q.Body).Decode(req); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusBadRequest, err)
			return
		}

		err := repos.InsertMenu(r.store, req.MenuItem, req.MenuParent)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusInternalServerError, err)
			return
		}
		result := map[string]string{"result": "ok"}

		r.respond(w, q, http.StatusOK, result)
	}
}

func (r *Router) handleDeleteMenu() http.HandlerFunc {
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

		err := repos.DeleteMenu(r.store, req.ID)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusInternalServerError, err)
			return
		}
		result := map[string]string{"result": "ok"}

		r.respond(w, q, http.StatusOK, result)
	}
}

func (r *Router) handleUpdateMenu() http.HandlerFunc {
	type request struct {
		ID         int    `json:"id"`
		MenuItem   string `json:"menuitem"`
		MenuParent int    `json:"menuparent"`
	}

	return func(w http.ResponseWriter, q *http.Request) {
		req := &request{}
		if err := json.NewDecoder(q.Body).Decode(req); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusBadRequest, err)
			return
		}

		err := repos.UpdateMenu(r.store, req.ID, req.MenuItem, req.MenuParent)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, q, http.StatusInternalServerError, err)
			return
		}
		result := map[string]string{"result": "ok"}

		r.respond(w, q, http.StatusOK, result)
	}
}
