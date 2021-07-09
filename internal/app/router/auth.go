package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/auth"
)

func (r *Router) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, q *http.Request) {
		if q.Header["Token"] != nil {
			if auth.UserAuthenticate(r.store, q.Header["Token"][0]) {
				next.ServeHTTP(w, q)
				return
			}
		} else {
			r.logger.Logger.Error("Token not accepted")
		}
		var errUnauthorized = errors.New("Unauthorised")
		r.logger.Logger.Error("Incorrect token")
		r.error(w, http.StatusUnauthorized, errUnauthorized)
	})
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
			r.error(w, http.StatusBadRequest, err)
			return
		}

		token, err := auth.UserAuthorize(r.store, req.Login, req.Password)
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusUnauthorized, err)
			return
		}
		result := map[string]string{"jwt": token}
		r.respond(w, http.StatusOK, result)
	}
}
