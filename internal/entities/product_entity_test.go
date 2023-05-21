package entities_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewProduct(t *testing.T) {

	u, err := entities.NewProduct("name", "description", "image", 10, 10)

	assert.Nil(t, err)
	assert.Equal(t, u.Name, "name")
	assert.Equal(t, u.Description, "description")
	assert.NotNil(t, u.ImageURL)
}

func TestShouldNotCreateNewProductInvalidPrice(t *testing.T) {

	p, err := entities.NewProduct("name", "description", "image", -10, 10)

	assert.NotNil(t, err)
	assert.Equal(t, err, entities.ErrInvalidPrice)
	assert.Nil(t, p)

}

func TestShouldNotCreateNewProductInvalidQuantity(t *testing.T) {

	p, err := entities.NewProduct("name", "description", "image", 10, 0)

	assert.NotNil(t, err)
	assert.Equal(t, err, entities.ErrInvalidQuantity)
	assert.Nil(t, p)

}
