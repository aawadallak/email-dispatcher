package message

type Message struct {
	from       string
	to         []string
	cc         []string
	subject    string
	body       string
	attachment []*Attachment
}

func (m *Message) From() string {
	return m.from
}

func (m *Message) To() []string {
	return m.to
}

func (m *Message) CC() []string {
	return m.cc
}

func (m *Message) Subject() string {
	return m.subject
}

func (m *Message) Body() string {
	return m.body
}

func (m *Message) Attachments() []*Attachment {
	return m.attachment
}

type MessageBuilder struct {
	message Message
}

func (b *MessageBuilder) WithFrom(from string) *MessageBuilder {
	b.message.from = from
	return b
}

func (b *MessageBuilder) WithTo(to []string) *MessageBuilder {
	b.message.to = to
	return b
}

func (b *MessageBuilder) WithCC(cc []string) *MessageBuilder {
	b.message.cc = cc
	return b
}

func (b *MessageBuilder) WithSubject(subject string) *MessageBuilder {
	b.message.subject = subject
	return b
}

func (b *MessageBuilder) WithBody(body string) *MessageBuilder {
	b.message.body = body
	return b
}

func (b *MessageBuilder) WithAttachments(attachments []*Attachment) *MessageBuilder {
	b.message.attachment = attachments
	return b
}

func (b *MessageBuilder) Create() *Message {
	return &b.message
}
