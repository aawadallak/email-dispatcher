package email

import (
	"latest/config"
	"latest/domain/mail"
)

var email *mail.Mail

func Setup() {
	email = mail.NewMail(
		config.GetConfig().MailSmtp,
		config.GetConfig().MailUser,
		config.GetConfig().MailPassword,
		config.GetConfig().MailPort,
	)
}

func GetInstance() *mail.Mail {
	return email
}
