package main

import (
	"net/smtp"
)

type MailSender struct {
	MailProviderConfigs *[]MailProviderConfig
}

type MailProviderConfig struct {
	Provider Provider
	Auth     *smtp.Auth
	Host     string
	HostPort string
	FromMail string
}

func NewMailSender() *MailSender {

	var mailProviderConfigs []MailProviderConfig
	for _, provider := range AllProviders() {
		mailProviderConfigs = append(mailProviderConfigs, provider.Config())
	}

	return &MailSender{
		MailProviderConfigs: &mailProviderConfigs,
	}
}

func (ms *MailSender) SendMail(requestData SendMailRequestData) error {
	var providerConfigMatch *MailProviderConfig
	for _, providerConfig := range *ms.MailProviderConfigs {
		if providerConfig.Provider == requestData.Provider {
			providerConfigMatch = &providerConfig
			break
		}
	}
	if providerConfigMatch == nil {
		panic("No provider config found")
	}
	var msg []byte
	if requestData.IsHTML {
		msg = []byte("Subject: " + requestData.Subject + "\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n" +
			"\r\n" + requestData.Body)
	} else {
		if requestData.Subject != "" {
			msg = []byte("Subject: " + requestData.Subject + "\r\n" + "\r\n" + requestData.Body)
		} else {
			msg = []byte(requestData.Body)
		}
	}
	return smtp.SendMail(providerConfigMatch.Host+":"+providerConfigMatch.HostPort, *providerConfigMatch.Auth, providerConfigMatch.FromMail, requestData.Addresses, msg)
}
