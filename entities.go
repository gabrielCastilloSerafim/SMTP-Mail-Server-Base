package main

type SendMailRequestBody struct {
	Provider  Provider `json:"provider"`
	Addresses []string `json:"addresses"`
	Subject   string   `json:"subject"`
	Body      string   `json:"body"`
}
