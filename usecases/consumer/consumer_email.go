package consumer

import (
	"latest/domain/consumer"
	"latest/domain/mail"
	"latest/pkg/logger"
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
			logger.Instance().Error(err.Error())
			continue
		}

		sendErr := d.mail.SendMessageBase64(msg)

		if sendErr != nil {
			continue
		}

	}
}
