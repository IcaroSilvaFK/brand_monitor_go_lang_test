package entities

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidPrice    = errors.New("invalid price")
	ErrInvalidQuantity = errors.New("invalid quantity")
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

func NewProduct(
	name, description, image string, price, quantity int,
) (*Product, error) {

	p := &Product{
		ID:          uuid.NewString(),
		Name:        name,
		Price:       price,
		ImageURL:    image,
		Description: description,
		Quantity:    quantity,
	}

	if err := p.ValidatePrice(); err != nil {
		return nil, err
	}

	if err := p.ValidateQuantity(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) ValidatePrice() error {

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}

func (p *Product) ValidateQuantity() error {
	if p.Quantity <= 0 {
		return ErrInvalidQuantity
	}

	return nil
}
