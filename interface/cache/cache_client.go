package cache

import (
	"time"
)

type CacheClient interface {
	GetString(key string) (string, error)
	Get(key string, object interface{}) error
	SetString(key, value string, duration time.Duration) error
	Set(key string, object interface{}, duration time.Duration) error
	Delete(key string) error
}
