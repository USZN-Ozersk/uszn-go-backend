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
			defer file.Close()
		}

		m.SenderName = q.FormValue("name")
		m.SenderAddress = q.FormValue("address")
		m.SenderEmail = q.FormValue("email")
		m.SenderPhone = q.FormValue("phone")
		m.QuestionSubject = q.FormValue("subj")
		m.QuestionText = q.FormValue("text")
		m.ReplySendMethod = q.FormValue("sendm")

		err = m.SendMail()
		if err != nil {
			r.logger.Logger.Error(err)
			r.error(w, http.StatusInternalServerError, err)
		} else {
			r.respond(w, http.StatusOK, nil)
		}
	}
}
