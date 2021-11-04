package dto

import err "latest/domain/error"

type ErrorDTO struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Module  string `json:"module,omitempty"`
}

func (e *ErrorDTO) EntityToDTO(err *err.Err) ErrorDTO {
	return ErrorDTO{
		Code:    err.Code(),
		Message: err.Message(),
		Module:  err.Module(),
	}
}
