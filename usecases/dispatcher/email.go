package dispatcher

import (
	"latest/config"
	"latest/domain"
	"latest/domain/mail"
	"latest/dto"
)

type Dispatcher struct {
	mail mail.Repository
}

func NewUsecase(mail mail.Repository) EmailUsecase {
	return &Dispatcher{
		mail: mail,
	}
}

func (d *Dispatcher) MultipartAttachments(obj *dto.MultiPartEmailDTO) *domain.Err {

	dmn, err := obj.Convert2Entity()

	if err != nil {
		return err
	}

	go func() {
		err = d.mail.SendMessage(dmn)
		if err != nil {
			config.Logger().Warnf("Cant send email %s", err.Message())
		}
	}()

	return nil
}

func (d *Dispatcher) Base64Attachments(obj *dto.EmailDTO) *domain.Err {

	dmn, err := obj.Convert2Entity()

	if err != nil {
		return err
	}

	go func() {
		err = d.mail.SendMessageBase64(dmn)
		if err != nil {
			config.Logger().Warnf("Cant send email %s", err.Message())
		}
	}()

	return nil
}
