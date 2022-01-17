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

	//Client struct hold auth client
	Client struct {
		Auth smtp.Auth
		Host string
		Port string
	}
)

// Initailize email-client by provinding sender config
func New(c *Config) *Client {
	return &Client{
		Auth: smtp.PlainAuth("", c.Sender, c.Secret, c.Host),
		Host: c.Host,
		Port: c.Port,
	}

}

// Send that handler send email with it's subject
func (s *Client) Send(email *Email) error {

	// auth := smtp.PlainAuth("", s.Sender, s.Secret, s.Host)

	to := fmt.Sprintf("To:%s\r\n", email.Recipients[0:])

	subject := fmt.Sprintf("Subject:%s\r\n", email.Subject)

	message := []byte(to + subject + "\r\n" + email.Body)

	// Send actual email
	err := smtp.SendMail(s.Host+":"+s.Port, s.Auth, email.Sender, email.Recipients, message)

	if err != nil {
		return err
	}
	return nil
}
