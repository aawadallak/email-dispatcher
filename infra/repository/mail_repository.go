package repository

import (
	"latest/domain/mail"
	"latest/domain/message"
	"log"

	"gopkg.in/gomail.v2"
)

type MailRepository struct {
	*gomail.Dialer
}

func NewMailRepository(mail *mail.Mail) MailRepository {
	return MailRepository{
		Dialer: gomail.NewDialer(mail.Smtp(), int(mail.Port()), mail.User(), mail.Password()),
	}
}

func (m MailRepository) SendMessage(e message.Message) error {
	y := gomail.NewMessage()
	y.SetHeader("From", e.From())
	y.SetHeader("To", e.To()...)
	y.SetHeader("Cc", e.CC()...)
	y.SetHeader("Subject", e.Subject())
	y.SetBody("text/html", e.Body())

	if err := m.Dialer.DialAndSend(y); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
