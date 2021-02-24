package controller

import (
	"test/custom_errors"
	"test/interface/auth"
	"test/interface/service"
	"test/model"

	"github.com/gin-gonic/gin"
)

type productCardController struct {
	productCardService service.ProductCardService
}

func (c *productCardController) GetByUserID(ctx *gin.Context) {
	claim, ok := ctx.Get("Claim")
	if !ok {
		ctx.JSON(401, model.NewErrorResponse("No claims found"))
		return
	}

	parsedClaim, ok := claim.(*auth.JwtClaim)
	if !ok {
		ctx.JSON(500, model.NewErrorResponse("Cant parse user claims"))
		return
	}

	cardItems, err := c.productCardService.GetByUserID(parsedClaim.UserID)
	if err != nil {
		if _, ok := err.(*custom_errors.EntityNotFoundError); ok {
			ctx.JSON(404, model.NewErrorResponse(err.Error()))
			return
		}

		ctx.JSON(500, model.NewErrorResponse(err.Error()))
		return
	}

	cardItemsDto := make([]model.ProductCardDto, 0)
	for _, cardItem := range cardItems {
		if cardItem.Product == nil {
			ctx.JSON(500, model.NewErrorResponse("product card does not contain non-nil product"))
			return
		}

		productDto := model.ProductDto{
			ID:          cardItem.Product.ID,
			Name:        cardItem.Product.Name,
			Description: cardItem.Product.Description,
			Quantity:    cardItem.Product.Quantity,
			Price:       cardItem.Product.Price,
		}

		cardItemDto := model.ProductCardDto{
			ProductID: cardItem.ProductID,
			Product:   productDto,
			Quantity:  cardItem.Quantity,
		}

		cardItemsDto = append(cardItemsDto, cardItemDto)
	}

	ctx.JSON(200, cardItemsDto)
}

func (c *productCardController) DeleteItemFromCard(ctx *gin.Context) {
	claim, ok := ctx.Get("Claim")
	if !ok {
		ctx.JSON(401, model.NewErrorResponse("No claims found"))
		return
	}

	parsedClaim, ok := claim.(*auth.JwtClaim)
	if !ok {
		ctx.JSON(500, model.NewErrorResponse("Cant parse user claims"))
		return
	}

	request := model.DeleteProductCardItemRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(422, model.NewErrorResponse(err.Error()))
		return
	}

	err = c.productCardService.DeleteItemFromCard(parsedClaim.UserID, request.ProductID)
	if err != nil {
		if _, ok := err.(*custom_errors.EntityNotFoundError); ok {
			ctx.JSON(404, model.NewErrorResponse(err.Error()))
			return
		} else if _, ok := err.(*custom_errors.BadRequestError); ok {
			ctx.JSON(400, model.NewErrorResponse(err.Error()))
			return
		}

		ctx.JSON(500, model.NewErrorResponse(err.Error()))
		return
	}

	ctx.Status(200)

}

func (c *productCardController) ChangeItemQuantityInCard(ctx *gin.Context) {
	claim, ok := ctx.Get("Claim")
	if !ok {
		ctx.JSON(401, model.NewErrorResponse("No claims found"))
		return
	}

	parsedClaim, ok := claim.(*auth.JwtClaim)
	if !ok {
		ctx.JSON(500, model.NewErrorResponse("Cant parse user claims"))
		return
	}

	request := model.ChangeProductCardItemQuantityRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(422, model.NewErrorResponse(err.Error()))
		return
	}

	cardItem, err := c.productCardService.ChangeItemCountInCard(parsedClaim.UserID, request.ProductID, request.Quantity)
	if err != nil {
		if _, ok := err.(*custom_errors.EntityNotFoundError); ok {
			ctx.JSON(404, model.NewErrorResponse(err.Error()))
			return
		} else if _, ok := err.(*custom_errors.BadRequestError); ok {
			ctx.JSON(400, model.NewErrorResponse(err.Error()))
			return
		}

		ctx.JSON(500, model.NewErrorResponse(err.Error()))
		return
	}

	if cardItem.Product == nil {
		ctx.JSON(500, model.NewErrorResponse("product card does not contain non-nil product"))
		return
	}

	productDto := model.ProductDto{
		ID:          cardItem.Product.ID,
		Name:        cardItem.Product.Name,
		Description: cardItem.Product.Description,
		Quantity:    cardItem.Product.Quantity,
		Price:       cardItem.Product.Price,
	}

	cardItemDto := model.ProductCardDto{
		ProductID: cardItem.ProductID,
		Product:   productDto,
		Quantity:  cardItem.Quantity,
	}

	ctx.JSON(200, &cardItemDto)
}

func NewProductCardController(productCardService service.ProductCardService) *productCardController {
	if productCardService == nil {
		panic("productCardService argument must not be nil")
	}

	return &productCardController{productCardService: productCardService}
}
