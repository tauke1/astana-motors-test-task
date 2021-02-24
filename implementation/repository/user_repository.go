package repository

import (
	"errors"
	"test/database/model"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) Get(id uint) (*model.User, error) {
	user := model.User{}
	err := repo.db.Find(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, err
}

func (repo *userRepository) GetByUsername(username string) (*model.User, error) {
	user := model.User{}
	err := repo.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, err
}

func (repo *userRepository) Create(user *model.User) error {
	if user == nil {
		return errors.New("user argument must be not nil")
	}

	return repo.db.Create(user).Error
}

func NewUserRepository(db *gorm.DB) *userRepository {
	if db == nil {
		panic("db argument must not be nil")
	}

	return &userRepository{
		db: db,
	}
}
