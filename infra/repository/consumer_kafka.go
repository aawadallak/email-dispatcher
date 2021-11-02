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

	dmn := domain.NewMessage(dto.From, dto.To, dto.Subject, dto.Template)

	return &dmn, nil
}
