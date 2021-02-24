package repository

import "test/database/model"

type ProductCardRepository interface {
	GetByUserID(userID uint) ([]model.ProductCard, error)
	GetByUserIDAndProductID(userID, productID uint) (*model.ProductCard, error)
	Create(productCart *model.ProductCard) error
	Update(productCart *model.ProductCard) error
	Delete(ID uint) error
}
