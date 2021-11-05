package dto

import (
	"latest/domain"
	"latest/domain/message"
)

type EmailDTO struct {
	From       string           `json:"from" validate:"required,email"`
	To         []string         `json:"to" validate:"required,dive,email"`
	Subject    string           `json:"subject" validate:"required"`
	Cc         []string         `json:"cc,omitempty"`
	Body       string           `json:"body,omitempty"`
	Attachment []*AttachmentDTO `json:"attachments,omitempty" validate:"dive"`
}

func (m *EmailDTO) Convert2Entity() (*message.Message, *domain.Err) {

	var files []*message.Attachment

	for _, v := range m.Attachment {

		file := message.NewAttachment(v.Name, v.Content)
		files = append(files, &file)
	}

	domain := new(message.MessageBuilder).
		WithFrom(m.From).
		WithTo(m.To).
		WithCC(m.Cc).
		WithSubject(m.Subject).
		WithBody(m.Body).
		WithAttachments(files).
		Create()

	return domain, nil
}
