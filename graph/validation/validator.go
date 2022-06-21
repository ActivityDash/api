package validation

import "github.com/go-playground/validator/v10"

func CustomValidator() *validator.Validate {
	instance := validator.New()
	instance.RegisterValidation("password", PasswordValidation)
	return instance
}
