package controller

import (
	"github.com/gin-gonic/gin"
)

type ProductCardController interface {
	GetByUserID(c *gin.Context)
	DeleteItemFromCard(c *gin.Context)
	ChangeItemQuantityInCard(c *gin.Context)
}
