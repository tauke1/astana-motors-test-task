package registry

import (
	"test/configuration"
	"test/core"
	"test/database"
	dbModel "test/database/model"
	"test/implementation/auth"
	"test/implementation/cache"
	"test/implementation/controller"
	"test/implementation/hasher"
	"test/implementation/repository"
	"test/implementation/service"
	"test/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterServices() *core.Core {
	hasher := hasher.NewSha256Hasher()
	jwtWrapper := auth.NewJwtWrapper(configuration.C.JwtSecret, configuration.C.JwtIssuer, int64(configuration.C.JwtExpirationHours))
	redisClient := cache.NewRedisClient(configuration.C.RedisPassword, configuration.C.RedisDatabase, configuration.C.RedisAddress)
	db := database.NewDB(configuration.C.DbHost, configuration.C.DbUser, configuration.C.DbPassword, configuration.C.DbName)
	db.AutoMigrate(&dbModel.User{}, &dbModel.Product{}, &dbModel.ProductCard{})
	if configuration.C.RunDBSeed {
		database.SeedData(db)
	}

	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	productCardRepository := repository.NewProductCardRepository(db)
	userService := service.NewUserService(userRepository, hasher, jwtWrapper, redisClient)
	productService := service.NewProductService(productRepository)
	productCardService := service.NewProductCardService(productService, userService, productCardRepository)
	userConroller := controller.NewUserController(userService)
	productController := controller.NewProductController(productService)
	productCardController := controller.NewProductCardController(productCardService)
	return core.NewCore(
		userConroller,
		productController,
		productCardController,
		userService,
	)
}

func RegisterServicesAndRoutes(engine *gin.Engine) {
	core := RegisterServices()
	if engine == nil {
		panic("engine argument must not be nil")
	}

	authHandler := middleware.TokenAuthMiddleware(core.UserService)
	engine.POST("api/auth/signin", core.UserController.Authenticate)
	engine.POST("api/auth/register", core.UserController.Register)
	engine.POST("api/auth/logout", authHandler, core.UserController.Logout)
	engine.GET("api/auth/info", authHandler, core.UserController.GetInfo)

	engine.GET("api/products", authHandler, core.ProductController.GetWithPagination)
	engine.GET("api/products/:ID", authHandler, core.ProductController.Get)
	engine.GET("api/products-all", authHandler, core.ProductController.GetAll)
	engine.POST("api/products", authHandler, core.ProductController.Create)
	engine.PUT("api/products/:ID", authHandler, core.ProductController.Update)
	engine.DELETE("api/products/:ID", authHandler, core.ProductController.Delete)

	engine.GET("api/card", authHandler, core.ProductCardController.GetByUserID)
	engine.POST("api/card", authHandler, core.ProductCardController.ChangeItemQuantityInCard)
	engine.DELETE("api/card", authHandler, core.ProductCardController.DeleteItemFromCard)
}
