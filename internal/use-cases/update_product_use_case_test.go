package use_cases_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	use_cases "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/use-cases"
	"github.com/stretchr/testify/assert"
)

func TestShouldUpdateProduct(t *testing.T) {

	db := database.NewDatabaseConnection()
	productRepository := repositories.NewProductRepository(db)
	createProductUseCase := use_cases.NewCreateProductUseCase(productRepository)
	updateProductUseCase := use_cases.NewUpdateProductUseCase(productRepository)
	findProductById := use_cases.NewFindProductByIdUseCase(productRepository)

	p, _ := createProductUseCase.Execute("Test1", "Test1", nil, 100, 1)

	err := updateProductUseCase.Execute(p.ID, "Update", "Update", nil, 120, 2)

	assert.Nil(t, err)

	product, _ := findProductById.Execute(p.ID)

	assert.NotNil(t, product)
	assert.Equal(t, product.Name, "Update")
	assert.Equal(t, product.Description, "Update")
	assert.Equal(t, product.Price/100, 120)
	assert.Equal(t, product.Quantity, 2)
	assert.NotNil(t, product.ImageURL)

}
