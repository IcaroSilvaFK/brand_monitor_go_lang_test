package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/http/dtos"
	use_cases "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/use-cases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController struct {
	createProductUseCase   use_cases.CreateProductUseCaseInterface
	listAllProductsUseCase use_cases.ListAllProductsUseCaseInterface
	findByIdProductUseCase use_cases.FindProductByIdUseCaseInterface
	deleteProductUseCase   use_cases.DeleteProductUseCaseInterface
	updateProductUseCase   use_cases.UpdateProductUseCaseInterface
}

type ProductControllerInterface interface {
	CreateProduct(ctx *gin.Context)
	ListAllProducts(ctx *gin.Context)
	FindById(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
}

func NewProductController(
	createProductUseCase use_cases.CreateProductUseCaseInterface,
	listAllProductsUseCase use_cases.ListAllProductsUseCaseInterface,
	findByIdProductUseCase use_cases.FindProductByIdUseCaseInterface,
	deleteProductUseCase use_cases.DeleteProductUseCaseInterface,
	updateProductUseCase use_cases.UpdateProductUseCaseInterface,
) ProductControllerInterface {

	return &ProductController{
		createProductUseCase:   createProductUseCase,
		listAllProductsUseCase: listAllProductsUseCase,
		findByIdProductUseCase: findByIdProductUseCase,
		deleteProductUseCase:   deleteProductUseCase,
		updateProductUseCase:   updateProductUseCase,
	}
}

func (pct *ProductController) CreateProduct(ctx *gin.Context) {

	file, _ := ctx.FormFile("image")
	var p dtos.CreateProductInput

	if err := ctx.Bind(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}

	product, err := pct.createProductUseCase.Execute(
		p.Name,
		p.Description,
		file,
		p.Price,
		p.Quantity,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"product": product,
	})

}

func (pct *ProductController) ListAllProducts(ctx *gin.Context) {

	query := ctx.Query("query")

	fmt.Println(query)

	products, err := pct.listAllProductsUseCase.Execute(query)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
	}

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"products": products,
	})

}

func (pct *ProductController) FindById(ctx *gin.Context) {

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

	p, err := pct.findByIdProductUseCase.Execute(id)

	fmt.Println(p)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "product.tmpl", gin.H{
		"product": p,
		"title":   p.Name,
	})
}

func (pct *ProductController) DeleteProduct(ctx *gin.Context) {

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

	err := pct.deleteProductUseCase.Execute(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)

}

func (pct *ProductController) UpdateProduct(ctx *gin.Context) {

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

	var p dtos.CreateProductInput

	if err := ctx.Bind(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}

	file, _ := ctx.FormFile("image")
	err := pct.updateProductUseCase.Execute(
		id,
		p.Name,
		p.Description,
		file,
		p.Price,
		p.Quantity,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"timestamp": time.Now(),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)

}
