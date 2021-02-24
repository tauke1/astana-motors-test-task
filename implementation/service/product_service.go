package service

import (
	"errors"
	"fmt"
	"test/custom_errors"
	"test/database/model"
	"test/interface/repository"
)

type productService struct {
	productRepository repository.ProductRepository
}

func (s *productService) Create(product *model.Product) error {
	if product == nil {
		return errors.New("product argument must not be nil")
	}

	return s.productRepository.Create(product)
}

func (s *productService) Update(ID uint, product *model.Product) error {
	if product == nil {
		return errors.New("product argument must not be nil")
	}

	product.ID = ID
	_, err := s.Get(ID)
	if err != nil {
		return err
	}

	return s.productRepository.Update(product)
}

func (s *productService) Delete(ID uint) error {
	_, err := s.Get(ID)
	if err != nil {
		return err
	}

	return s.productRepository.Delete(ID)
}

func (s *productService) Get(id uint) (*model.Product, error) {
	productFromDb, err := s.productRepository.Get(id)
	if err != nil {
		return nil, err
	} else if productFromDb == nil {
		return nil, custom_errors.NewEntityNotFoundError(fmt.Sprintf("product %d not found", id))
	}

	return productFromDb, nil
}

func (s *productService) GetAll() ([]model.Product, error) {
	products, err := s.productRepository.GetAll()
	return products, err
}

func (s *productService) GetByPagination(pageNumber uint, pageSize uint) (*repository.ProductPaginationResult, error) {
	if pageNumber == 0 {
		return nil, errors.New("pageNumber argument must be positive integer")
	}

	if pageSize == 0 {
		return nil, errors.New("pageSize argument must be positive integer")
	}

	products, err := s.productRepository.GetByPagination(pageNumber, pageSize)
	return products, err
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	if productRepository == nil {
		panic("productRepository argument must not be nil")
	}

	return &productService{productRepository: productRepository}
}
