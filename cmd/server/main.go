package main

import (
	"crypto/tls"
	"log"
	"strings"

	"github.com/LIYINGZHEN/gosmtp/pkg/config"
	"github.com/LIYINGZHEN/gosmtp/pkg/smtp"
)

func main() {
	config := config.New()

	smtpServer := &smtp.SmtpServer{
		Host:     config.Host,
		Port:     config.Port,
		Email:    config.Sender,
		Password: config.Password,
		TlsConfig: &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         config.Host,
		},
	}

	mail := &smtp.Mail{
		Sender:  config.Sender,
		To:      strings.Split(config.To, ","),
		Subject: config.Subject,
		Body:    config.Body,
	}

	serivce := smtp.New(smtpServer, mail)
	err := serivce.Send()
	if err != nil {
		log.Printf("[Error] Unable to send emails: %v", err)
	}
}
