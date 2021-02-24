package custom_errors

type EntityNotFoundError struct {
	message string
}

func (e *EntityNotFoundError) Error() string {
	return e.message
}

func NewEntityNotFoundError(message string) *EntityNotFoundError {
	return &EntityNotFoundError{message: message}
}
