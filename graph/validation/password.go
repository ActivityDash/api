package validation

import (
	"github.com/go-playground/validator/v10"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func PasswordValidation(fieldLevel validator.FieldLevel) bool {
	err := passwordvalidator.Validate(fieldLevel.Field().String(), 50)
	return err == nil
}
