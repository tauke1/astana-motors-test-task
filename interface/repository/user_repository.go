package repository

import (
	"test/database/model"
)

type UserRepository interface {
	Get(id uint) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	Create(user *model.User) error
}
