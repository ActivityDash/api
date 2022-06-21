package customerror

const (
	AlreadyExistsFieldErrorCode    = "ALREADY_EXISTS"
	InvalidFieldErrorCode          = "INVALID"
	NotProvidedFieldErrorCode      = "NOT_PROVIDED"
	InsecurePasswordFieldErrorCode = "INSECURE_PASSWORD"
)

type InvalidField struct {
	Name  string
	Error string
}
