package dto

type EmailDTO struct {
	From       string           `json:"from,omitempty"`
	To         string           `json:"to,omitempty"`
	Subject    string           `json:"subject,omitempty"`
	Cc         string           `json:"cc,omitempty"`
	Template   string           `json:"template,omitempty"`
	Attachment []*AttachmentDTO `json:"attachments,omitempty"`
}
