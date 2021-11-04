package dispatcher

import (
	"latest/domain"
	"latest/dto"
)

type Usecases interface {
	EventDispatch()
	DisptachEmail(obj *dto.MultiPartEmailDTO) *domain.Err
}
