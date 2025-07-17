package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type SendMailRequestBody struct {
	Provider  Provider `json:"provider"`
	Addresses []string `json:"addresses"`
	Subject   string   `json:"subject"`
	Body      string   `json:"body"`
}

func sendMailHandler(w http.ResponseWriter, req *http.Request, mailSender *MailSender) {
	// Get mail to send data from reuquest body
	var sendMailData SendMailRequestBody
	err := json.NewDecoder(req.Body).Decode(&sendMailData)
	// Validate received data
	allAddressesAreValid := true
	for _, address := range sendMailData.Addresses {
		if !isValidEmail(address) {
			allAddressesAreValid = false
			break
		}
	}
	if err != nil || !allAddressesAreValid || len(sendMailData.Addresses) == 0 || sendMailData.Body == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Send mail
	err = mailSender.SendMail(sendMailData.Provider, sendMailData.Addresses, sendMailData.Subject, sendMailData.Body)
	if err != nil {
		log.Printf("💥 Failed to send with addresses: %v, and error: %s", sendMailData.Addresses, err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusOK)
}
