package app

import (
	"latest/config"
	"latest/domain/mail"
	"latest/infra/repository"
	"latest/usecases/dispatcher"
)

func Start() {

	setup := mail.NewMail(
		config.GetConfig().MailSmtp,
		config.GetConfig().MailUser,
		config.GetConfig().MailPassword,
		config.GetConfig().MailPort,
	)

	consumer := repository.NewConsumer()
	mail := repository.NewMailRepository(setup)

	svc := dispatcher.NewEmailDispatcher(mail, consumer)

	svc.EventDispatch()
}
