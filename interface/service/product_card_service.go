package service

import "test/database/model"

type ProductCardService interface {
	GetByUserID(userID uint) ([]model.ProductCard, error)
	DeleteItemFromCard(userID, productID uint) error
	ChangeItemCountInCard(userID, productID, quantity uint) (*model.ProductCard, error)
}
