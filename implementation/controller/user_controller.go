package controller

import (
	"strings"
	"test/custom_errors"
	"test/interface/auth"
	"test/interface/service"
	"test/model"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func (c *userController) Authenticate(context *gin.Context) {
	signinRequest := model.SigninRequest{}
	err := context.ShouldBindJSON(&signinRequest)
	if err != nil {
		context.JSON(422, model.NewErrorResponse(err.Error()))
		return
	}

	result, err := c.userService.Authenticate(signinRequest.Username, signinRequest.Password)
	if err != nil {
		if _, ok := err.(*custom_errors.UnauthorizedError); ok {
			context.JSON(401, model.NewErrorResponse(err.Error()))
		} else {
			context.JSON(500, model.NewErrorResponse(err.Error()))
		}

		return
	}

	response := model.SigninResponse{
		Token:     result.Token,
		TokenType: result.TokenType,
	}
	context.JSON(200, &response)
}

func (c *userController) Logout(context *gin.Context) {
	reqToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		context.JSON(400, model.NewErrorResponse("No bearer token found"))
		return
	}

	token := splitToken[1]
	err := c.userService.Logout(token)
	if err != nil {
		context.JSON(500, err.Error())
		return
	}

	context.Status(200)
}

func (c *userController) GetInfo(context *gin.Context) {
	claim, ok := context.Get("Claim")
	if !ok {
		context.JSON(401, model.NewErrorResponse("No claims found"))
		return
	}

	parsedClaim, ok := claim.(*auth.JwtClaim)
	if !ok {
		context.JSON(500, model.NewErrorResponse("Cant parse user claims"))
		return
	}

	user, err := c.userService.Get(parsedClaim.UserID)
	if err != nil {
		context.JSON(500, err.Error())
		return
	}

	response := model.UserDto{
		Username: user.Username,
		ID:       user.ID,
	}

	context.JSON(200, &response)
}

func (c *userController) Register(context *gin.Context) {
	registerRequest := model.RegisterUserRequest{}
	err := context.ShouldBindJSON(&registerRequest)
	if err != nil {
		context.JSON(422, model.NewErrorResponse(err.Error()))
		return
	}

	user, err := c.userService.Register(registerRequest.Username, registerRequest.Password)
	if err != nil {
		if _, ok := err.(*custom_errors.EntityExistsError); ok {
			context.JSON(409, model.NewErrorResponse(err.Error()))
		} else {
			context.JSON(500, model.NewErrorResponse(err.Error()))
		}

		return
	}

	context.JSON(200, model.UserDto{ID: user.ID, Username: user.Username})
}

func NewUserController(userService service.UserService) *userController {
	if userService == nil {
		panic("userService argument must not be nil")
	}

	return &userController{userService: userService}
}
