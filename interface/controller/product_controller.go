package controller

import "github.com/gin-gonic/gin"

type ProductController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	GetWithPagination(c *gin.Context)
}
