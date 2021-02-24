package custom_errors

import "fmt"

type KeyNotFoundError struct {
	key string
}

func (e *KeyNotFoundError) Error() string {
	return fmt.Sprintf("key %s not found in cache", e.key)
}

func NewKeyNotFoundError(key string) *KeyNotFoundError {
	return &KeyNotFoundError{key: key}
}
