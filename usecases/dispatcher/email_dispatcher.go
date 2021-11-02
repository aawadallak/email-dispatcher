package dispatcher

import (
	"latest/config"
	"latest/domain/consumer"
	"latest/domain/mail"
)

type Dispatcher struct {
	mail     mail.Repository
	consumer consumer.Repository
}

func NewEmailDispatcher(mail mail.Repository, consumer consumer.Repository) Usecases {
	return &Dispatcher{
		mail:     mail,
		consumer: consumer,
	}
}

func (d *Dispatcher) EventDispatch() {
	for {

		msg, err := d.consumer.Consumer()

		if err != nil {
			config.Logger().Error(err.Error())
			continue
		}

		err = d.mail.SendMessage(*msg)

		if err != nil {
			config.Logger().Warn(err.Error())
			continue
		}
	}
}
