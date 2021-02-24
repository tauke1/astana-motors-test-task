package repository

import "test/database/model"

type ProductCardRepository interface {
	GetByUserID(userID uint) ([]model.ProductCard, error)
	GetByUserIDAndProductID(userID, productID uint) (*model.ProductCard, error)
	Create(productCart *model.ProductCard) error
	Update(productCart *model.ProductCard) error
	Delete(id uint) error
}
