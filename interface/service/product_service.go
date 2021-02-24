package service

import (
	"test/database/model"
	"test/interface/repository"
)

type ProductService interface {
	Create(product *model.Product) error
	Update(ID uint, product *model.Product) error
	Delete(ID uint) error
	Get(ID uint) (*model.Product, error)

	GetAll() ([]model.Product, error)
	GetByPagination(pageNumber uint, pageSize uint) (*repository.ProductPaginationResult, error)
}
