package repository

import "test/database/model"

type ProductRepository interface {
	Create(*model.Product) error
	Update(product *model.Product) error
	Delete(id uint) error
	Get(id uint) (*model.Product, error)

	GetAll() ([]model.Product, error)
	GetByPagination(pageNumber uint, pageSize uint) (*ProductPaginationResult, error)
}

type ProductPaginationResult struct {
	PageSize   uint
	PageNumber uint
	TotalCount uint
	Products   []model.Product
}
