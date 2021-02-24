package core

import (
	"test/interface/controller"
	"test/interface/service"
)

type Core struct {
	UserController        controller.UserController
	ProductController     controller.ProductController
	ProductCardController controller.ProductCardController
	UserService           service.UserService
}

func NewCore(userController controller.UserController, productController controller.ProductController, productCardController controller.ProductCardController, userService service.UserService) *Core {
	if userController == nil {
		panic("userConroller argument must not be nil")
	}

	if productController == nil {
		panic("productController argument must not be nil")
	}

	if productCardController == nil {
		panic("productCardController argument must not be nil")
	}

	if userService == nil {
		panic("userService argument must not be nil")
	}

	return &Core{
		UserController:        userController,
		ProductController:     productController,
		ProductCardController: productCardController,
		UserService:           userService,
	}
}
