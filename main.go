package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Config and start server
	mailSender := NewMailSender()
	router := http.NewServeMux()
	router.HandleFunc("POST /v1/sendMail", handlerWithMailServer(sendMailHandler, mailSender))
	handlerWithCors := cors.AllowAll().Handler(router)
	log.Printf("🌍 Server listening on port: %v 🌍", "3001")
	panic(http.ListenAndServe(":3001", handlerWithCors))
}
