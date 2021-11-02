package message

type Message struct {
	from       string
	to         string
	cc         string
	subject    string
	body       string
	attachment *[]Attachment
}

func NewMessage(from, to, subject, body string) Message {
	return Message{
		from:    from,
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (m *Message) From() string {
	return m.from
}

func (m *Message) To() string {
	return m.to
}

func (m *Message) CC() string {
	return m.cc
}

func (m *Message) Subject() string {
	return m.subject
}

func (m *Message) Body() string {
	return m.body
}

func (m *Message) Attachments() *[]Attachment {
	return m.attachment
}
