package custom_errors

type BadRequestError struct {
	message string
}

func (e *BadRequestError) Error() string {
	return e.message
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{message: message}
}
