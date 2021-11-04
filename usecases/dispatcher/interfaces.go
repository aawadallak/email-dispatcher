package dispatcher

import (
	"latest/domain"
	"latest/dto"
)

type Usecases interface {
	EventDispatch()
	MultipartAttachments(obj *dto.MultiPartEmailDTO) *domain.Err
	EncondedAttachments(obj *dto.EmailDTO) *domain.Err
}
