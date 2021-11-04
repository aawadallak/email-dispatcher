package dispatcher

import (
	"latest/config"
	"latest/domain"
	"latest/domain/consumer"
	"latest/domain/mail"
	"latest/dto"
)

type Dispatcher struct {
	mail     mail.Repository
	consumer consumer.Repository
}

func NewConsumerDispatcher(mail mail.Repository, consumer consumer.Repository) Usecases {
	return &Dispatcher{
		mail:     mail,
		consumer: consumer,
	}
}

func NewEmailDispatcher(mail mail.Repository) Usecases {
	return &Dispatcher{
		mail: mail,
	}
}

func (d *Dispatcher) EventDispatch() {
	for {

		msg, err := d.consumer.Consumer()

		if err != nil {
			config.Logger().Error(err.Error())
			continue
		}

		sendErr := d.mail.SendMessage(msg)

		if sendErr != nil {
			config.Logger().Warn(sendErr.Message())
			continue
		}

	}
}

func (d *Dispatcher) MultipartAttachments(obj *dto.MultiPartEmailDTO) *domain.Err {

	dmn, err := obj.Convert2Entity()

	if err != nil {
		return err
	}

	go func() error {
		err = d.mail.SendMessage(dmn)
		if err != nil {
			config.Logger().Warnf("Cant send email %s", err.Message())
		}
		return nil
	}()

	return nil
}

func (d *Dispatcher) EncondedAttachments(obj *dto.EmailDTO) *domain.Err {

	dmn, err := obj.Convert2Entity()

	if err != nil {
		return err
	}

	go func() error {
		err = d.mail.SendMessageBase64(dmn)
		if err != nil {
			config.Logger().Warnf("Cant send email %s", err.Message())
		}
		return nil
	}()

	return nil
}
