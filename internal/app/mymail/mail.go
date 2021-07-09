package mymail

import (
	"bytes"
	"io"
	"mime/multipart"

	mail "github.com/xhit/go-simple-mail"
)

type MailMessage struct {
	File     multipart.File
	FileName string
}

func (m *MailMessage) SendMail() error {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, m.File); err != nil {
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = "smtp.beget.com"
	server.Port = 465

	email := mail.NewMSG()
	email.SetFrom("feedback@usznozersk.ru").AddTo("uszn.it@ozerskadm.ru").SetSubject("Обращение гражданина")
	email.SetBody(mail.TextPlain, "Сообщение")
	email.Attach(&mail.File{Data: buf.Bytes(), Name: m.FileName})

	return nil
}
