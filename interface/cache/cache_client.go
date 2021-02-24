package cache

import "errors"

var notFoundError = "key not found"

type CacheClient interface {
	GetString(key string) (string, error)
	Get(key string, object interface{}) error
	SetString(key, value string) error
	Set(key string, object interface{}) error
	Delete(key string) error
}

func IsKeyNotFound(err error) bool {
	if err == nil {
		return false
	}

	return err.Error() == notFoundError
}

func CreateKeyNotFoundError() error {
	return errors.New(notFoundError)
}
