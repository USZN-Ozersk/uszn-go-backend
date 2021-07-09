package mymail

import (
	"bytes"
	"crypto/tls"
	"io"
	"mime/multipart"

	"github.com/USZN-Ozersk/uszn-go-backend/internal/app/config"

	mail "github.com/xhit/go-simple-mail/v2"
)

type MailMessage struct {
	config          *config.Config
	File            multipart.File
	FileName        string
	SenderName      string
	SenderAddress   string
	SenderEmail     string
	SenderPhone     string
	QuestionText    string
	QuestionSubject string
}

func (m *MailMessage) SendMail() error {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, m.File); err != nil {
		return err
	}

	messageBody := ""
	messageBody += "Фамилия Имя Отчество гражданина: " + m.SenderName + ".\n\r"
	messageBody += "Тема обращения гражданина: " + m.QuestionSubject + ".\n\r"
	messageBody += "Адрес регистрации гражданина: " + m.SenderAddress + ".\n\r"
	messageBody += "Адрес электронной почты гражданина: " + m.SenderEmail + ".\n\r"
	messageBody += "Телефон гражданина: " + m.SenderPhone + ".\n\r"
	messageBody += "Текст обращения гражданина: " + m.QuestionText + ".\n\r"

	server := mail.NewSMTPClient()
	server.Host = m.config.MailServer
	server.Port = m.config.MailPort
	server.Username = m.config.MailUser
	server.Password = m.config.MailPassword
	server.Encryption = mail.EncryptionSSLTLS
	server.Authentication = mail.AuthLogin
	server.KeepAlive = false
	server.TLSConfig = &tls.Config{ServerName: m.config.MailServer}
	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom("feedback@usznozersk.ru").AddTo("uszn.it@ozerskadm.ru").SetSubject("Обращение гражданина:" + m.SenderName)
	email.SetBody(mail.TextPlain, messageBody)
	if m.File != nil {
		email.Attach(&mail.File{Data: buf.Bytes(), Name: m.FileName})
		if email.Error != nil {
			return email.Error
		}
	}

	if err = email.Send(smtpClient); err != nil {
		return err
	}

	return nil
}
