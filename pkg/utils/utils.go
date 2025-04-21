package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New("validation failed: " + err.Field() + " " + err.Tag())

		}
	}
	return nil
}
