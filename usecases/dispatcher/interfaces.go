package dispatcher

import (
	"latest/domain"
	"latest/dto"
)

type EmailUsecase interface {
	MultipartAttachments(obj *dto.MultiPartEmailDTO) *domain.Err
	Base64Attachments(obj *dto.EmailDTO) *domain.Err
}
