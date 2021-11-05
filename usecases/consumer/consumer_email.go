package consumer

import (
	"latest/config"
	"latest/domain/consumer"
	"latest/domain/mail"
)

type Consumer struct {
	mail     mail.Repository
	consumer consumer.Repository
}

func NewUsecase(mail mail.Repository, consumer consumer.Repository) ConsumerUsecase {
	return &Consumer{
		mail:     mail,
		consumer: consumer,
	}
}

func (d *Consumer) Dispatch() {
	for {

		msg, err := d.consumer.Consumer()

		if err != nil {
			config.Logger().Error(err.Error())
			continue
		}

		sendErr := d.mail.SendMessageBase64(msg)

		if sendErr != nil {
			config.Logger().Warn(sendErr.Message())
			continue
		}

	}
}
