package mail

import (
	"github.com/hwangseonu/paperless.dev/internal/common"
	"gopkg.in/gomail.v2"
)

const from = "hwangseonu123@gmail.com"

type Content struct {
	To      []string
	Subject string
	Body    string
}

type Client struct {
	dialer *gomail.Dialer
}

func NewClient(config common.SMTPConfig) *Client {
	dialer := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)

	return &Client{
		dialer: dialer,
	}
}

func (c *Client) SendMail(mail Content) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", mail.To...)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Body)

	return c.dialer.DialAndSend(m)
}
