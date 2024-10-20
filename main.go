package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load dot env file if needed
	_, err := os.Stat("config.env")
	if err == nil {
		godotenv.Load("config.env")
	}
	// Config and start server
	mailSender := newMailSender()
	router := http.NewServeMux()
	router.HandleFunc("POST /v1/sendMail", handlerWithMailServer(sendMailHandler, mailSender))
	handlerWithCors := cors.AllowAll().Handler(router)
	log.Printf("🌍 Server listening on port: %v 🌍", "3001")
	panic(http.ListenAndServe(":3001", handlerWithCors))
}
