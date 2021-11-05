package app

import (
	"latest/config/email"
	"latest/infra/repository"
	"latest/pkg/kafka"
	"latest/pkg/logger"
	"latest/usecases/consumer"
)

func Start() {

	logger.Instance().Info("Starting service")

	mail := repository.NewMailRepository(email.GetInstance())

	event := kafka.NewConsumer()

	repository := repository.NewConsumer(event.Reader())

	svc := consumer.NewUsecase(mail, repository)

	svc.Dispatch()
}
