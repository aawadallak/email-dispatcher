package err

type Err struct {
	code    int
	message string
	module  string
}

func (e *Err) Code() int {
	return e.code
}

func (e *Err) Message() string {
	return e.message
}
func (e *Err) Module() string {
	return e.module
}

func NewError(code int, message, module string) Err {
	return Err{
		code:    code,
		message: message,
		module:  module,
	}
}
