package hasher

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

type sha256Hasher struct{}

func (hasher *sha256Hasher) Hash(input string) (string, error) {
	if input == "" {
		return "", errors.New("input must be not empty")
	}

	shaHasher := sha256.New()
	shaHasher.Write([]byte(input))
	sha := base64.URLEncoding.EncodeToString(shaHasher.Sum(nil))
	return sha, nil
}

func NewSha256Hasher() *sha256Hasher {
	return &sha256Hasher{}
}
