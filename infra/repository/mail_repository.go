package repository

import (
	"io"
	"latest/domain"
	"latest/domain/mail"
	"latest/domain/message"

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

func (m MailRepository) SendMessage(e *message.Message) *domain.Err {
	y := gomail.NewMessage()
	y.SetHeader("From", e.From())
	y.SetHeader("To", e.To()...)
	y.SetHeader("Cc", e.CC()...)
	y.SetHeader("Subject", e.Subject())
	y.SetBody("text/html", e.Body())

	for _, attach := range e.Attachments() {
		y.Attach(attach.Name(), gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write([]byte(attach.Content()))
			return err
		}))
	}

	if err := m.Dialer.DialAndSend(y); err != nil {
		return domain.NewError(500, err.Error())
	}
	return nil
}

func (m MailRepository) SendMessageBase64(e *message.Message) *domain.Err {
	y := gomail.NewMessage()
	y.SetHeader("From", e.From())
	y.SetHeader("To", e.To()...)
	y.SetHeader("Cc", e.CC()...)
	y.SetHeader("Subject", e.Subject())
	y.SetBody("text/html", e.Body())

	for _, attach := range e.Attachments() {
		y.Attach(attach.Name(), gomail.SetCopyFunc(func(w io.Writer) error {
			f, err := attach.DecodeBase64()
			if err != nil {
				return err
			}
			_, err = w.Write(f)
			return err
		}))
	}

	if err := m.Dialer.DialAndSend(y); err != nil {
		return domain.NewError(500, err.Error())
	}
	return nil
}
