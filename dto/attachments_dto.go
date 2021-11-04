package dto

type AttachmentDTO struct {
	Name    string `json:"name,omitempty" validate:"required"`
	Content string `json:"content,omitempty" validate:"required,decoded"`
}
