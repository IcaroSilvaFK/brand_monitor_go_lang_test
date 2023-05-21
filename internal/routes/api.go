package routes

import (
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/http/controllers"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	use_cases "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/use-cases"
	"github.com/gin-gonic/gin"
)

func NewApiRoutes(r *gin.Engine) {

	db := database.NewDatabaseConnection()
	userRepository := repositories.NewUserRepository(db)
	createUserUseCase := use_cases.NewCreateUserUseCase(userRepository)
	findByEmailUserUseCase := use_cases.NewFindUserByEmailUseCase(userRepository)
	deleteUserUseCase := use_cases.NewDeleteUserUseCase(userRepository)

	userController := controllers.NewUserController(
		createUserUseCase,
		findByEmailUserUseCase,
		deleteUserUseCase,
	)

	userGroup := r.Group("/users")

	userGroup.POST("", userController.CreateUser)
	userGroup.POST("/signIn", userController.FindUserByEmail)
	userGroup.DELETE("/:id", userController.DeleteUser)

	productRepository := repositories.NewProductRepository(db)
	createProductUseCase := use_cases.NewCreateProductUseCase(productRepository)
	listAllProductsUseCase := use_cases.NewListAllProductsUseCase(productRepository)
	findByIdProductsUseCase := use_cases.NewFindProductByIdUseCase(productRepository)
	deleteProductUseCase := use_cases.NewDeleteProductUseCase(productRepository)
	updateProductUseCase := use_cases.NewUpdateProductUseCase(productRepository)

	productController := controllers.NewProductController(
		createProductUseCase,
		listAllProductsUseCase,
		findByIdProductsUseCase,
		deleteProductUseCase,
		updateProductUseCase,
	)

	productGroup := r.Group("/products")

	productGroup.POST("", productController.CreateProduct)
	productGroup.GET("", productController.ListAllProducts)
	productGroup.GET("/:id", productController.FindById)
	productGroup.PUT("/:id", productController.UpdateProduct)
	productGroup.DELETE("/:id", productController.DeleteProduct)

}
