package services

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
	"placemaking-backend-go/config"
)

// SendEmail envia um e-mail para o usuário.
func SendEmail(to string, content string, subject string) error {
	from := config.LoadSettings().GmailUser
	password := config.LoadSettings().GmailPass
	smtpServer := "smtp.gmail.com"
	smtpPort := "465"

	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = subject
	e.HTML = []byte(content)

	// Configuração para conexão segura com TLS
	smtpAuth := smtp.PlainAuth("", from, password, smtpServer)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer,
	}

	// Criando conexão SMTP segura
	smtpAddr := fmt.Sprintf("%s:%s", smtpServer, smtpPort)
	conn, err := tls.Dial("tcp", smtpAddr, tlsConfig)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao servidor SMTP: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		return fmt.Errorf("erro ao criar cliente SMTP: %w", err)
	}
	defer client.Close()

	// Autenticando no servidor SMTP
	if err := client.Auth(smtpAuth); err != nil {
		return fmt.Errorf("erro ao autenticar no SMTP: %w", err)
	}

	// Enviando e-mail
	if err := e.SendWithTLS(smtpAddr, smtpAuth, tlsConfig); err != nil {
		return fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	fmt.Printf("E-mail enviado para %s com sucesso!\n", to)
	return nil
}