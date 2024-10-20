package main

import (
	"net/smtp"
	"os"
)

type MailSender struct {
	Auth     *smtp.Auth
	Host     string
	HostPort string
	FromMail string
}

func newMailSender() *MailSender {
	host := os.Getenv("HOST")
	hostPort := os.Getenv("HOST_PORT")
	fromMail := os.Getenv("FROM_MAIL")
	googleAppPass := os.Getenv("GOOGLE_APP_PASSWORD")
	if host == "" || hostPort == "" || fromMail == "" || googleAppPass == "" {
		panic("Failed to load env data")
	}
	auth := smtp.PlainAuth("", fromMail, googleAppPass, host)
	return &MailSender{
		Auth:     &auth,
		Host:     host,
		HostPort: hostPort,
		FromMail: fromMail,
	}
}

func (ms *MailSender) sendMail(toMailAddresses []string, subject string, body string) error {
	var msg []byte
	if subject != "" {
		msg = []byte("Subject: " + subject + "\r\n" + "\r\n" + body)
	} else {
		msg = []byte(body)
	}
	return smtp.SendMail(ms.Host+":"+ms.HostPort, *ms.Auth, ms.FromMail, toMailAddresses, msg)
}
