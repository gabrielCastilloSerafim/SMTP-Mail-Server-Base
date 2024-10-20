package main

type SendMailRequestBody struct {
	Addresses []string `json:"addresses"`
	Subject   string   `json:"subject"`
	Body      string   `json:"body"`
}
