package message

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
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

func (a *Attachment) EncodeToBase64(path string) (string, error) {
	// Open file on disk.
	f, err := os.Open(fmt.Sprintf("./tmp/%s", path))

	if err != nil {
		return "", err
	}

	// Read entire file into byte slice.
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)

	if err != nil {
		return "", err
	}

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded, nil
}

func (a *Attachment) DecodeBase64(content string) ([]byte, error) {

	decode, err := base64.StdEncoding.DecodeString(content)

	if err != nil {
		return nil, err
	}

	return decode, nil
}
