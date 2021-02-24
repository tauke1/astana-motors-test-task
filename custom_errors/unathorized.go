package custom_errors

type UnauthorizedError struct {
	message string
}

func (err *UnauthorizedError) Error() string {
	return err.message
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{message: message}
}
