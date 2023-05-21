package use_cases

import (
	"context"
	"mime/multipart"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/configs"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UpdateProductUseCase struct {
	productRepository repositories.ProductRepositoryInterface
}

type UpdateProductUseCaseInterface interface {
	Execute(id, name, description string, image *multipart.FileHeader, price float64, quantity int) error
}

func NewUpdateProductUseCase(
	productRepository repositories.ProductRepositoryInterface,
) UpdateProductUseCaseInterface {

	return &UpdateProductUseCase{
		productRepository: productRepository,
	}
}

func (pdr *UpdateProductUseCase) Execute(
	id, name, description string,
	image *multipart.FileHeader,
	price float64,
	quantity int,
) error {

	var p entities.Product

	convertToCents := int(price * 100)
	p.Description = description
	p.Name = name
	p.Price = convertToCents
	p.Quantity = quantity

	if image != nil {
		cdn, err := configs.NewCloudinaryConfig()

		if err != nil {
			return err
		}

		file, _ := image.Open()

		ctx := context.Background()
		cld, _ := cdn.Cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: p.ID})

		p.ImageURL = cld.SecureURL
	}

	err := pdr.productRepository.Update(&p, id)

	if err != nil {
		return err
	}

	return nil
}
