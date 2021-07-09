package router

import (
	"net/http"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/mymail"
)

func (r *Router) handleMail() http.HandlerFunc {

	return func(w http.ResponseWriter, q *http.Request) {
		m := &mymail.MailMessage{}
		if err := q.ParseMultipartForm(32 << 20); err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusBadRequest, err)
			return
		}

		file, handler, err := q.FormFile("file")
		if err == nil {
			m.File = file
			m.FileName = handler.Filename
		}

		err = m.SendMail()
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusInternalServerError, err)
		} else {
			r.respond(w, http.StatusOK, nil)
		}
	}
}
