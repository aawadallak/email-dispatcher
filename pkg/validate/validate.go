package validate

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

func Struct(str interface{}) error {
	return Validate.Struct(str)
}
