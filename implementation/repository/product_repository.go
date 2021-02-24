package repository

import (
	"errors"
	"test/database/model"
	"test/interface/repository"

	"github.com/jinzhu/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func (repo *productRepository) Create(product *model.Product) error {
	if product == nil {
		return errors.New("product argument must not be nil")
	}

	if product.ID != 0 {
		return errors.New("product.ID must be zero")
	}

	return repo.db.Create(product).Error
}

func (repo *productRepository) Update(product *model.Product) error {
	if product == nil {
		return errors.New("product argument must not be nil")
	}

	if product.ID == 0 {
		return errors.New("product.ID must not be zero")
	}

	return repo.db.Save(product).Error
}

func (repo *productRepository) Delete(id uint) error {
	return repo.db.Delete(&model.Product{}, id).Error
}

func (repo *productRepository) Get(id uint) (*model.Product, error) {
	product := &model.Product{}
	err := repo.db.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return product, nil
}

func (repo *productRepository) GetAll() ([]model.Product, error) {
	products := make([]model.Product, 0)
	err := repo.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, err
}

func (repo *productRepository) GetByPagination(pageNumber uint, pageSize uint) (*repository.ProductPaginationResult, error) {
	if pageNumber == 0 {
		return nil, errors.New("pageNumber must be positive integer")
	}

	if pageSize == 0 {
		return nil, errors.New("pageSize must be positive integer")
	}

	var totalCount uint
	err := repo.db.Model(&model.Product{}).Count(&totalCount).Error
	if err != nil {
		return nil, err
	}

	offset := (pageNumber - 1) * pageSize

	products := make([]model.Product, 0)
	err = repo.db.Limit(pageSize).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}

	result := &repository.ProductPaginationResult{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		TotalCount: totalCount,
		Products:   products,
	}

	return result, nil
}

func NewProductRepository(db *gorm.DB) *productRepository {
	if db == nil {
		panic("db argument must not be nil")
	}

	return &productRepository{db: db}
}
