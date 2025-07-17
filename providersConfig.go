package main

import "net/smtp"

type Provider int

const (
	WillBeEnd Provider = iota
)

func AllProviders() []Provider {
	return []Provider{WillBeEnd}
}

func (p Provider) Config() MailProviderConfig {
	switch p {
	case WillBeEnd:
		host := "smtp.gmail.com"
		hostPort := "587"
		fromMail := "willbeend1@gmail.com"
		googleAppPass := "zziwdfxyxyemcrcp"
		auth := smtp.PlainAuth("", fromMail, googleAppPass, host)
		return MailProviderConfig{
			Provider: WillBeEnd,
			Host:     host,
			HostPort: hostPort,
			FromMail: fromMail,
			Auth:     &auth,
		}
	}
	panic("Unknown case")
}
