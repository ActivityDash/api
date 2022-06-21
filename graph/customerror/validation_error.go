package customerror

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	ValidationErrorCode = "VALIDATION_ERROR"
	UniquenessErrorCode = "UNIQUENESS_ERROR"
)

type ValidationError struct {
	Code        string
	Description string
	Fields      []InvalidField
}

func (e *ValidationError) Error() string {
	if e.Description != "" {
		return e.Description
	}

	return "validation error"
}

func (e *ValidationError) SetUsingValidatorError(err validator.ValidationErrors) {
	var errorMessages []string
	var fields []InvalidField

	for _, fieldErr := range err {
		fieldName := strings.ToLower(fieldErr.Field())

		var errorCode string
		switch fieldErr.Tag() {
		case "required":
			errorCode = NotProvidedFieldErrorCode
			errorMessages = append(errorMessages, fmt.Sprintf("%s must be provided", fieldName))
		case "password":
			errorCode = InsecurePasswordFieldErrorCode
			errorMessages = append(errorMessages, fmt.Sprintf("%s is insecure", fieldName))
		default:
			errorCode = InvalidFieldErrorCode
			errorMessages = append(errorMessages, fmt.Sprintf("%s is invalid", fieldName))
		}

		fields = append(fields, InvalidField{Name: fieldName, Error: errorCode})
	}

	e.Code = ValidationErrorCode
	e.Description = strings.Join(errorMessages, ". ")
	e.Fields = fields
}

func (e *ValidationError) SetGqlErrorFields(err *gqlerror.Error) {
	// We iterate over field structs to select
	// specific field attributes to return.
	fields := []map[string]string{}
	for _, field := range e.Fields {
		fields = append(fields, map[string]string{
			"error": field.Error,
			"name":  field.Name,
		})
	}

	err.Message = e.Error()
	err.Extensions = map[string]interface{}{
		"code":   e.Code,
		"fields": fields,
	}
}
