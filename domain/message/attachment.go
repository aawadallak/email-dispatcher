package message

import (
	"encoding/base64"
)

type Attachment struct {
	name    string
	content string
}

func NewAttachment(name string, content string) Attachment {
	return Attachment{
		name:    name,
		content: content,
	}
}

func (a *Attachment) Name() string {
	return a.name
}

func (a *Attachment) Content() string {
	return a.content
}

func (a *Attachment) DecodeBase64() ([]byte, error) {

	decode, err := base64.StdEncoding.DecodeString(a.content)

	if err != nil {
		return nil, err
	}

	return decode, nil
}
