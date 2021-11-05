package mail

import (
	"latest/domain"
	"latest/domain/message"
)

type Repository interface {
	SendMessage(e *message.Message) *domain.Err
	SendMessageBase64(e *message.Message) *domain.Err
}
