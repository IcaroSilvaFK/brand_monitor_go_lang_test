package use_cases_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	use_cases "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/use-cases"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateProduct(t *testing.T) {

	db := database.NewDatabaseConnection()
	productRepository := repositories.NewProductRepository(db)
	createProductUseCase := use_cases.NewCreateProductUseCase(productRepository)

	p, err := createProductUseCase.Execute("name", "description", nil, 10.0, 1)

	assert.Nil(t, err)
	assert.NotNil(t, p)
}
