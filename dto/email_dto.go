package dto

type EmailDTO struct {
	From       string           `json:"from"`
	To         []string         `json:"to"`
	Subject    string           `json:"subject"`
	Cc         []string         `json:"cc,omitempty"`
	Body       string           `json:"body"`
	Attachment []*AttachmentDTO `json:"attachments,omitempty"`
}
