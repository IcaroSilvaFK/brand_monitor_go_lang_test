package use_cases

import (
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
)

type ListAllProductsUseCase struct {
	repository repositories.ProductRepositoryInterface
}

type ListAllProductsUseCaseInterface interface {
	Execute(query string) ([]*entities.Product, error)
}

func NewListAllProductsUseCase(
	repository repositories.ProductRepositoryInterface,
) ListAllProductsUseCaseInterface {

	return &ListAllProductsUseCase{
		repository: repository,
	}
}

func (pdr *ListAllProductsUseCase) Execute(query string) ([]*entities.Product, error) {

	products, err := pdr.repository.FindAll(query)

	return products, err
}
