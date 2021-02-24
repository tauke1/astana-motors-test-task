package repository

import (
	"errors"
	"test/database/model"

	"github.com/jinzhu/gorm"
)

type productCardRepository struct {
	db *gorm.DB
}

func (repo *productCardRepository) GetByUserID(userID uint) ([]model.ProductCard, error) {
	cardItems := make([]model.ProductCard, 0)
	err := repo.db.Preload("Product").Where("user_id = ?", userID).Find(&cardItems).Error
	if err != nil {
		return nil, err
	}

	return cardItems, nil
}

func (repo *productCardRepository) GetByUserIDAndProductID(userID, productID uint) (*model.ProductCard, error) {
	cardItem := model.ProductCard{}
	err := repo.db.Where("user_id = ? and product_id = ?", userID, productID).First(&cardItem).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &cardItem, nil
}

func (repo *productCardRepository) Create(productCard *model.ProductCard) error {
	if productCard == nil {
		return errors.New("productCard argument must not be nil")
	}

	if productCard.ID != 0 {
		return errors.New("productCard.ID must be zero")
	}

	return repo.db.Create(productCard).Error
}

func (repo *productCardRepository) Update(productCard *model.ProductCard) error {
	if productCard == nil {
		return errors.New("productCard argument must not be nil")
	}

	if productCard.ID == 0 {
		return errors.New("productCard.ID must not be zero")
	}

	return repo.db.Save(productCard).Error
}

func (repo *productCardRepository) Delete(ID uint) error {
	return repo.db.Delete(&model.ProductCard{}, ID).Error
}

func NewProductCardRepository(db *gorm.DB) *productCardRepository {
	if db == nil {
		panic("db argument must not be nil")
	}

	return &productCardRepository{db: db}
}
