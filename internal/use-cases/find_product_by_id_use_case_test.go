package use_cases_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	use_cases "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/use-cases"
	"github.com/stretchr/testify/assert"
)

func TestShouldFindProductById(t *testing.T) {

	db := database.NewDatabaseConnection()
	productRepo := repositories.NewProductRepository(db)
	findByIdProductUseCase := use_cases.NewFindProductByIdUseCase(productRepo)
	createProductUseCase, _ := use_cases.NewCreateProductUseCase(productRepo).Execute(
		"Name",
		"Description",
		nil,
		12.2,
		1,
	)

	p, err := findByIdProductUseCase.Execute(createProductUseCase.ID)

	assert.Nil(t, err)
	assert.NotNil(t, p)

}
