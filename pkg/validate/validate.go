package validate

import (
	"encoding/base64"
	"latest/config"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	err := Validate.RegisterValidation("decoded", isBase64)
	if err != nil {
		config.Logger().Warn(err.Error())
	}
}

func Struct(str interface{}) error {
	return Validate.Struct(str)
}

func isBase64(fl validator.FieldLevel) bool {
	// get value
	value := fl.Field().Interface().(string)

	_, err := base64.StdEncoding.DecodeString(value)

	if err != nil {
		return err == nil
	}

	return true
}
