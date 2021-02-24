package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Authenticate(context *gin.Context)
	Logout(context *gin.Context)
	Register(context *gin.Context)
	GetInfo(context *gin.Context)
}
