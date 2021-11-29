package email

import (
	"fmt"
	"net/smtp"
)

type (
	// Email contains all the details about the email to be sent
	Email struct {
		Subject    string // Unique id for the current Email
		Body       string // The content of the Email
		Sender     string
		Recipients []string // The recipients of this particular Email
	}
	Config struct {
		Sender string
		Host   string
		Port   string
		Secret string
	}
	// SendService is the principal interface.
	// it sends a Email if there is a problem returns an error
	SendService interface {
		Send(s *Config, email *Email) error
	}
)

func Send(s *Config, email *Email) error {

	auth := smtp.PlainAuth("", s.Sender, s.Secret, s.Host)

	to := fmt.Sprintf("To:%s\r\n", email.Recipients[0:])

	subject := fmt.Sprintf("Subject:%s\r\n", email.Subject)

	message := []byte(to + subject + "\r\n" + email.Body)

	// Send actual email
	err := smtp.SendMail(s.Host+":"+s.Port, auth, email.Sender, email.Recipients, message)

	if err != nil {
		return err
	}
	return nil
}
