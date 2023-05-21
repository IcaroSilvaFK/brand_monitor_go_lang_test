package use_cases

import (
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
)

type FindProductByIdUseCase struct {
	productRepository repositories.ProductRepositoryInterface
}

type FindProductByIdUseCaseInterface interface {
	Execute(id string) (*entities.Product, error)
}

func NewFindProductByIdUseCase(
	productRepository repositories.ProductRepositoryInterface,
) FindProductByIdUseCaseInterface {

	return &FindProductByIdUseCase{
		productRepository: productRepository,
	}

}

func (f *FindProductByIdUseCase) Execute(id string) (*entities.Product, error) {

	p, err := f.productRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	return p, nil
}
