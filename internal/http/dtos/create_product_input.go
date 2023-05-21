package dtos

type CreateProductInput struct {
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	Quantity    int     `json:"quantity" form:"quantity"`
}
