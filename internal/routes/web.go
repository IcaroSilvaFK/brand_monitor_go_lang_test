package routes

import (
	"net/http"
	"time"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	use_cases "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/use-cases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewWebRoutes(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.tmpl", gin.H{})
	})

	r.GET("/create", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "create_account.tmpl", gin.H{})
	})

	r.GET("/admin", func(ctx *gin.Context) {
		db := database.NewDatabaseConnection()
		productRepository := repositories.NewProductRepository(db)
		listAllProductsUseCase := use_cases.NewListAllProductsUseCase(productRepository)

		products, _ := listAllProductsUseCase.Execute("")

		ctx.HTML(http.StatusOK, "admin.tmpl", gin.H{
			"products": products,
		})
	})

	r.GET("admin/products/new/:id", func(ctx *gin.Context) {
		db := database.NewDatabaseConnection()
		productRepository := repositories.NewProductRepository(db)
		findByIDProductUseCase := use_cases.NewFindProductByIdUseCase(productRepository)

		id, ok := ctx.Params.Get("id")

		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":     "id not found",
				"timestamp": time.Now(),
			})
			return
		}

		if _, err := uuid.Parse(id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":     err.Error(),
				"timestamp": time.Now(),
			})
			return
		}

		product, _ := findByIDProductUseCase.Execute(id)

		ctx.HTML(http.StatusOK, "edit_product.tmpl", gin.H{
			"product": product,
		})
	})

	r.GET("/admin/products/new", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "product_new.tmpl", gin.H{})
	})

}
