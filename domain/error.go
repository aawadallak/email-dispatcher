package domain

type Err struct {
	code    int
	message string
}

func (e *Err) Code() int {
	return e.code
}

func (e *Err) Message() string {
	return e.message
}

func NewError(code int, message string) *Err {
	return &Err{
		code:    code,
		message: message,
	}
}

func (e *Err) ToJson() interface{} {
	return struct {
		Code    int
		Message string
	}{
		Code:    e.Code(),
		Message: e.Message(),
	}
}
