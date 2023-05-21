package use_cases

import "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"

type DeleteProductUseCase struct {
	productRepository repositories.ProductRepositoryInterface
}

type DeleteProductUseCaseInterface interface {
	Execute(id string) error
}

func NewDeleteProductUseCase(
	productRepository repositories.ProductRepositoryInterface,
) DeleteProductUseCaseInterface {

	return &DeleteProductUseCase{
		productRepository: productRepository,
	}

}

func (dpuc *DeleteProductUseCase) Execute(id string) error {
	return dpuc.productRepository.Delete(id)
}
