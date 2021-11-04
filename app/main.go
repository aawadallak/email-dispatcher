package app

import (
	"latest/config"
	"latest/domain/mail"
	"latest/infra/repository"
	"latest/usecases/dispatcher"
)

func Start() {

	config.Logger().Info("Starting service")

	setup := mail.NewMail(
		config.GetConfig().MailSmtp,
		config.GetConfig().MailUser,
		config.GetConfig().MailPassword,
		config.GetConfig().MailPort,
	)

	mail := repository.NewMailRepository(setup)

	consumer := repository.NewConsumer()

	svc := dispatcher.NewConsumerDispatcher(mail, consumer)

	svc.EventDispatch()
}
