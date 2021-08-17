package mymail

import (
	"bytes"
	"crypto/tls"
	"io"
	"mime/multipart"
	"os"
	"strconv"

	mail "github.com/xhit/go-simple-mail/v2"
)

type MailMessage struct {
	File            multipart.File
	FileName        string
	SenderName      string
	SenderAddress   string
	SenderEmail     string
	SenderPhone     string
	QuestionText    string
	QuestionSubject string
	ReplySendMethod string
}

func (m *MailMessage) SendMail() error {
	messageBody := ""
	messageBody += "Фамилия Имя Отчество гражданина: " + m.SenderName + ".\n\r"
	messageBody += "Тема обращения гражданина: " + m.QuestionSubject + ".\n\r"
	messageBody += "Адрес регистрации гражданина: " + m.SenderAddress + ".\n\r"
	messageBody += "Адрес электронной почты гражданина: " + m.SenderEmail + ".\n\r"
	messageBody += "Телефон гражданина: " + m.SenderPhone + ".\n\r"
	messageBody += "Текст обращения гражданина: " + m.QuestionText + ".\n\r"
	switch m.ReplySendMethod {
	case "email":
		messageBody += "Ответ направить по электронной почте.\n\r"
	case "letter":
		messageBody += "Ответ направить письмом на адрес регистрации.\n\r"
	default:
		messageBody += "Метод отправки сообщения не определен"
	}

	server := mail.NewSMTPClient()
	server.Host = os.Getenv("MAIL_SERVER")
	server.Port, _ = strconv.Atoi(os.Getenv("MAIL_PORT"))
	server.Username = os.Getenv("MAIL_USER")
	server.Password = os.Getenv("MAIL_PASSWORD")
	server.Encryption = mail.EncryptionSSLTLS
	server.Authentication = mail.AuthLogin
	server.KeepAlive = false
	server.TLSConfig = &tls.Config{ServerName: os.Getenv("MAIL_SERVER")}
	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom("feedback@usznozersk.ru").AddTo("ksz@ozerskadm.ru").SetSubject("Обращение гражданина:" + m.SenderName)
	email.SetBody(mail.TextPlain, messageBody)
	if m.File != nil {
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, m.File); err != nil {
			return err
		} else {
			email.Attach(&mail.File{Data: buf.Bytes(), Name: m.FileName})
		}
	}

	if email.Error != nil {
		return email.Error
	}

	if err = email.Send(smtpClient); err != nil {
		return err
	}
	return nil
}
