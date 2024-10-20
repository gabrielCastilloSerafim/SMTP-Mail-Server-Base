package main

import (
	"net/http"
	"net/mail"
)

type serverFunctionWithMailSender func(w http.ResponseWriter, req *http.Request, mailSender *MailSender)
type serverMuxDefaultFunction func(w http.ResponseWriter, req *http.Request)

func handlerWithMailServer(handlerWithMailSender serverFunctionWithMailSender, mailSender *MailSender) serverMuxDefaultFunction {
	return func(w http.ResponseWriter, req *http.Request) {
		handlerWithMailSender(w, req, mailSender)
	}
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
