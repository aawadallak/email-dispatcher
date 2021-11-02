package message

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
