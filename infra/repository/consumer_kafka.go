package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"latest/config"
	domain "latest/domain/message"
	"latest/dto"
	"strings"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer() *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: strings.Split(config.GetConfig().Brokers, ","),
			GroupID: fmt.Sprintf("consumer-%s", config.GetConfig().Topic),
			Topic:   config.GetConfig().Topic}),
	}
}

func (k *Consumer) Consumer() (*domain.Message, error) {

	ctx := context.Background()

	var dto dto.EmailDTO

	message, err := k.reader.ReadMessage(ctx)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(message.Value, &dto)

	if err != nil {
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
