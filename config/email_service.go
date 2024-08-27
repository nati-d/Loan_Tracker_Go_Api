package config

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"loan_tracker/bootstrap"

	"github.com/jordan-wright/email"
)

// getPassword fetches the email password from environment variables
func getPassword() string {
	password, err := bootstrap.GetEnv("EMAIL_PASSWORD")
	if err != nil {
		log.Println("Failed to get email password from environment variables")
		return ""
	}

	return password
}

// SendEmail sends an email with the specified content type (text or HTML)
func SendEmail(to, subject, body string, isHTML bool) error {
	e := email.NewEmail()
	e.From = "nathnaeldes@gmail.com"
	e.To = []string{to}
	e.Subject = subject

	if isHTML {
		e.HTML = []byte(body) // Set HTML content
	} else {
		e.Text = []byte(body) // Set plain text content
	}

	password := getPassword()
	if password == "" {
		log.Println("Email password is not set in environment variables")
		return fmt.Errorf("email password is not set")
	}

	// Create TLS configuration
	tlsConfig := &tls.Config{
		ServerName: "smtp.gmail.com",
		MinVersion: tls.VersionTLS12,
	}

	// Debug logging
	log.Printf("Sending email to %s with subject: %s", to, subject)

	// Send email with TLS
	err := e.SendWithTLS("smtp.gmail.com:465",
		smtp.PlainAuth("", "nathnaeldes@gmail.com", password, "smtp.gmail.com"),
		tlsConfig,
	)

	if err != nil {
		log.Printf("Failed to send email: %v\n", err)
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}
