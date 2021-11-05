package kafka

import (
	"fmt"
	"latest/config"
	"strings"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Brokers []string
	GroupID string
	Topic   string
}

func NewConsumer() *Consumer {
	return &Consumer{
		Brokers: strings.Split(config.GetConfig().Brokers, ","),
	}
}

func (c *Consumer) Reader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		GroupID: fmt.Sprintf("consumer-%s", config.GetConfig().Topic),
		Topic:   config.GetConfig().Topic})
}

func (c *Consumer) Writer() *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Topic: config.GetConfig().Topic,
	})
}
