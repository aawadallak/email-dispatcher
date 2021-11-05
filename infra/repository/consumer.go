package repository

import (
	"context"
	"encoding/json"
	domain "latest/domain/message"
	"latest/dto"
	"latest/pkg/validate"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(reader *kafka.Reader) *Consumer {
	return &Consumer{
		reader: reader,
	}
}

func (k *Consumer) Consumer() (*domain.Message, error) {

	ctx := context.Background()

	var dto dto.EmailDTO

	message, err := k.reader.ReadMessage(ctx)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(message.Value, &dto); err != nil {
		return nil, err
	}

	if err = validate.Struct(&dto); err != nil {
		return nil, err
	}

	var attachments []*domain.Attachment

	for _, v := range dto.Attachment {
		dmn := domain.NewAttachment(v.Name, v.Content)
		attachments = append(attachments, &dmn)
	}

	domain := new(domain.MessageBuilder).
		WithFrom(dto.From).
		WithTo(dto.To).
		WithCC(dto.Cc).
		WithSubject(dto.Subject).
		WithBody(dto.Body).
		WithAttachments(attachments).
		Create()

	return domain, nil
}
