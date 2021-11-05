package dto

import "latest/domain"

type ErrorDTO struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *ErrorDTO) EntityToDTO(err *domain.Err) ErrorDTO {
	return ErrorDTO{
		Code:    err.Code(),
		Message: err.Message(),
	}
}
