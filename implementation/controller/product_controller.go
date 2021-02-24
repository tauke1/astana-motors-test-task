package controller

import (
	"strconv"
	"test/custom_errors"
	dbModel "test/database/model"
	"test/interface/service"
	"test/model"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService service.ProductService
}

func (c *productController) Create(ctx *gin.Context) {
	productDto := model.ProductDto{}
	err := ctx.ShouldBindJSON(&productDto)
	if err != nil {
		ctx.JSON(422, model.NewErrorResponse(err.Error()))
		return
	}

	product := dbModel.Product{
		ID:          productDto.ID,
		Name:        productDto.Name,
		Description: productDto.Description,
		Quantity:    productDto.Quantity,
		Price:       productDto.Price,
	}

	err = c.productService.Create(&product)
	if err != nil {
		ctx.JSON(500, model.NewErrorResponse(err.Error()))
		return
	}

	response := model.ProductDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Quantity:    product.Quantity,
		Price:       product.Price,
	}

	ctx.JSON(200, &response)
}

func (c *productController) Update(ctx *gin.Context) {
	productIDString := ctx.Param("ID")
	productID64, err := strconv.ParseUint(productIDString, 10, 32)
	if err != nil {
		ctx.JSON(400, model.NewErrorResponse("route parameter ID is not an unsigned integer"))
		return
	}

	productID := uint(productID64)
	productDto := model.ProductDto{}
	err = ctx.ShouldBindJSON(&productDto)
	if err != nil {
		ctx.JSON(422, model.NewErrorResponse(err.Error()))
		return
	}

	product := dbModel.Product{
		ID:          productDto.ID,
		Name:        productDto.Name,
		Description: productDto.Description,
		Quantity:    productDto.Quantity,
		Price:       productDto.Price,
	}

	err = c.productService.Update(productID, &product)
	if err != nil {
		if _, ok := err.(*custom_errors.EntityNotFoundError); ok {
			ctx.JSON(404, model.NewErrorResponse(err.Error()))
			return
		}

		ctx.JSON(500, model.NewErrorResponse(err.Error()))
		return
	}

	response := model.ProductDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Quantity:    product.Quantity,
		Price:       product.Price,
	}

	ctx.JSON(200, &response)
}

func (c *productController) Delete(ctx *gin.Context) {
	productIDString := ctx.Param("ID")
	productID64, err := strconv.ParseUint(productIDString, 10, 32)
	if err != nil {
		ctx.JSON(400, model.NewErrorResponse("route parameter ID is not an unsigned integer"))
		return
	}

	productID := uint(productID64)
	err = c.productService.Delete(productID)
	if err != nil {
		if _, ok := err.(*custom_errors.EntityNotFoundError); ok {
			ctx.JSON(404, model.NewErrorResponse(err.Error()))
			return
		}

		ctx.JSON(500, model.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(200)
}

func (c *productController) Get(ctx *gin.Context) {
	productIDString := ctx.Param("ID")
	productID64, err := strconv.ParseUint(productIDString, 10, 32)
	if err != nil {
		ctx.JSON(400, model.NewErrorResponse("route parameter ID is not an unsigned integer"))
		return
	}

	productID := uint(productID64)
	product, err := c.productService.Get(productID)
	if err != nil {
		if _, ok := err.(*custom_errors.EntityNotFoundError); ok {
			ctx.JSON(404, model.NewErrorResponse(err.Error()))
			return
		}

		ctx.JSON(500, model.NewErrorResponse(err.Error()))
		return
	}

	productDto := model.ProductDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Quantity:    product.Quantity,
		Price:       product.Price,
	}

	ctx.JSON(200, &productDto)
}

func (c *productController) GetAll(ctx *gin.Context) {
	products, err := c.productService.GetAll()
	if err != nil {
		ctx.JSON(500, model.NewErrorResponse(err.Error()))
		return
	}

	productDtos := make([]model.ProductDto, 0)
	for _, product := range products {
		productDto := model.ProductDto{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Quantity:    product.Quantity,
			Price:       product.Price,
		}

		productDtos = append(productDtos, productDto)
	}

	ctx.JSON(200, productDtos)
}

func (c *productController) GetWithPagination(ctx *gin.Context) {
	pageSizeString := ctx.Query("pageSize")
	if pageSizeString == "" {
		ctx.JSON(400, model.NewErrorResponse("pageSize is required query parameter"))
		return
	}

	pageSize64, err := strconv.ParseUint(pageSizeString, 10, 32)
	if err != nil {
		ctx.JSON(400, model.NewErrorResponse("query parameter pageSize is not an unsigned integer"))
		return
	}

	pageSize := uint(pageSize64)

	pageNumberString := ctx.Query("pageNumber")
	if pageNumberString == "" {
		ctx.JSON(400, model.NewErrorResponse("pageNumber is required query parameter"))
		return
	}

	pageNumber64, err := strconv.ParseUint(pageNumberString, 10, 32)
	if err != nil {
		ctx.JSON(400, model.NewErrorResponse("query parameter pageNumber is not an unsigned integer"))
		return
	}

	pageNumber := uint(pageNumber64)

	productsWithPagination, err := c.productService.GetByPagination(pageNumber, pageSize)
	if err != nil {
		ctx.JSON(500, model.NewErrorResponse(err.Error()))
		return
	}

	productDtos := make([]model.ProductDto, 0)
	for _, product := range productsWithPagination.Products {
		productDto := model.ProductDto{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Quantity:    product.Quantity,
			Price:       product.Price,
		}

		productDtos = append(productDtos, productDto)
	}

	response := model.ProductsByPaginationResponse{
		PageSize:   productsWithPagination.PageSize,
		PageNumber: productsWithPagination.PageNumber,
		TotalCount: productsWithPagination.TotalCount,
		Products:   productDtos,
	}

	ctx.JSON(200, response)
}

func NewProductController(productService service.ProductService) *productController {
	if productService == nil {
		panic("productService argument must not be nil")
	}

	return &productController{productService: productService}
}
