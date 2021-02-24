package custom_errors

type EntityExistsError struct {
	message string
}

func (e *EntityExistsError) Error() string {
	return e.message
}

func NewEntityExistsError(message string) *EntityExistsError {
	return &EntityExistsError{message: message}
}
