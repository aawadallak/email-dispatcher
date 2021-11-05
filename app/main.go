package app

import (
	"latest/config"
	"latest/config/email"
	"latest/infra/repository"
	"latest/usecases/consumer"
)

func Start() {

	config.Logger().Info("Starting service")

	mail := repository.NewMailRepository(email.GetInstance())

	repository := repository.NewConsumer()

	svc := consumer.NewUsecase(mail, repository)

	svc.Dispatch()
}
