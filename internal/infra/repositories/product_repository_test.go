package repositories_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewProduct(t *testing.T) {

	p, _ := entities.NewProduct("Teste", "Teste", "Teste", 100, 1)
	conn := database.NewDatabaseConnection()

	productRepository := repositories.NewProductRepository(conn)

	err := productRepository.Create(p)

	assert.Nil(t, err)

}

func TestShouldFindProductById(t *testing.T) {

	p, _ := entities.NewProduct("Teste", "Teste", "Teste", 100, 1)
	conn := database.NewDatabaseConnection()

	productRepository := repositories.NewProductRepository(conn)

	productRepository.Create(p)

	p, err := productRepository.FindById(p.ID)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.Name, "Teste")
	assert.Equal(t, p.Description, "Teste")
	assert.Equal(t, p.ImageURL, "Teste")
}

func TestShouldDeleteProduct(t *testing.T) {

	p, _ := entities.NewProduct("Teste", "Teste", "Teste", 100, 1)
	conn := database.NewDatabaseConnection()

	productRepository := repositories.NewProductRepository(conn)

	productRepository.Create(p)

	err := productRepository.Delete(p.ID)

	assert.Nil(t, err)
}

func TestShouldFindAllProducts(t *testing.T) {

	conn := database.NewDatabaseConnection()

	productRepository := repositories.NewProductRepository(conn)

	products, err := productRepository.FindAll("")

	assert.Nil(t, err)
	assert.NotNil(t, products)
}
