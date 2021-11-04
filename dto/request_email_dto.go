package dto

import (
	"bytes"
	"io"
	"latest/domain/message"
	"mime/multipart"
)

type MultiPartEmailDTO struct {
	From       []string                `mapstructure:"from" validate:"required,min=1,max=1"`
	To         []string                `mapstructure:"to" validate:"required"`
	Subject    []string                `mapstructure:"subject" validate:"required,min=1,max=1"`
	Cc         []string                `mapstructure:"cc"`
	Body       []string                `mapstructure:"body,omitempty" validate:"min=1,max=1"`
	Attachment []*multipart.FileHeader `json:"attachments,omitempty"`
}

func (m *MultiPartEmailDTO) Convert2Entity() (*message.Message, *ErrorDTO) {

	var files []*message.Attachment

	for _, v := range m.Attachment {
		f, _ := v.Open()
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, f); err != nil {
			return nil, &ErrorDTO{
				Code:    500,
				Message: err.Error(),
			}
		}
		file := message.NewAttachment(v.Filename, buf.String())
		files = append(files, &file)
	}

	domain := new(message.MessageBuilder).
		WithFrom(m.From[0]).
		WithTo(m.To).
		WithCC(m.Cc).
		WithSubject(m.Subject[0]).
		WithBody(m.Body[0]).
		WithAttachments(files).
		Create()

	return domain, nil
}
