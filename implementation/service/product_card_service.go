package service

import (
	"fmt"
	"test/custom_errors"
	"test/database/model"
	"test/interface/repository"
	"test/interface/service"
)

type productCardService struct {
	productService        service.ProductService
	userService           service.UserService
	productCardRepository repository.ProductCardRepository
}

func (s *productCardService) GetByUserID(userID uint) ([]model.ProductCard, error) {
	_, err := s.userService.Get(userID)
	if err != nil {
		return nil, err
	}

	return s.productCardRepository.GetByUserID(userID)
}

func (s *productCardService) DeleteItemFromCard(userID, productID uint) error {
	_, err := s.userService.Get(userID)
	if err != nil {
		return err
	}

	_, err = s.productService.Get(userID)
	if err != nil {
		return err
	}

	productCard, err := s.productCardRepository.GetByUserIDAndProductID(userID, productID)
	if err != nil {
		return err
	} else if productCard == nil {
		return custom_errors.NewEntityNotFoundError(fmt.Sprintf("product card item by product_id %d and user_id %d not found", productID, userID))
	}

	return s.productCardRepository.Delete(productCard.ID)
}

func (s *productCardService) ChangeItemCountInCard(userID, productID, quantity uint) (*model.ProductCard, error) {
	if quantity == 0 {
		return nil, custom_errors.NewBadRequestError("quantity must be positive integer")
	}

	_, err := s.userService.Get(userID)
	if err != nil {
		return nil, err
	}

	product, err := s.productService.Get(productID)
	if err != nil {
		return nil, err
	}

	if product.Quantity < quantity {
		return nil, custom_errors.NewBadRequestError("quantity of given product is less than requested quantity")
	}

	productCard, err := s.productCardRepository.GetByUserIDAndProductID(userID, productID)
	if err != nil {
		return nil, err
	} else if productCard == nil {
		productCard = &model.ProductCard{
			UserID:    userID,
			ProductID: productID,
			Quantity:  quantity,
		}

		err = s.productCardRepository.Create(productCard)

	} else {
		productCard.Quantity = quantity
		err = s.productCardRepository.Update(productCard)
	}

	if err != nil {
		return nil, err
	}

	productCard.Product = product
	return productCard, nil
}

func NewProductCardService(productService service.ProductService, userService service.UserService, productCardRepository repository.ProductCardRepository) *productCardService {
	if productService == nil {
		panic("productService argument must not be nil")
	}

	if userService == nil {
		panic("userService argument must not be nil")
	}

	if productCardRepository == nil {
		panic("productCardRepository argument must not be nil")
	}

	return &productCardService{
		productService:        productService,
		userService:           userService,
		productCardRepository: productCardRepository,
	}
}
