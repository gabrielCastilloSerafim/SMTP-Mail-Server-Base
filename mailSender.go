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

func (ms *MailSender) SendMail(provider Provider, toAddresses []string, subject string, body string) error {
	var providerConfigMatch *MailProviderConfig
	for _, providerConfig := range *ms.MailProviderConfigs {
		if providerConfig.Provider == provider {
			providerConfigMatch = &providerConfig
			break
		}
	}
	if providerConfigMatch == nil {
		panic("No provider config found")
	}
	var msg []byte
	if subject != "" {
		msg = []byte("Subject: " + subject + "\r\n" + "\r\n" + body)
	} else {
		msg = []byte(body)
	}
	return smtp.SendMail(providerConfigMatch.Host+":"+providerConfigMatch.HostPort, *providerConfigMatch.Auth, providerConfigMatch.FromMail, toAddresses, msg)
}
