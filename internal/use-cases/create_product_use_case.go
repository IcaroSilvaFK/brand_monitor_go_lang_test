package use_cases

import (
	"context"
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/configs"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/smtp"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CreateProductUseCase struct {
	productRepository repositories.ProductRepositoryInterface
}

type CreateProductUseCaseInterface interface {
	Execute(name, description string, image *multipart.FileHeader, price float64, quantity int) (*entities.Product, error)
}

func NewCreateProductUseCase(
	productRepository repositories.ProductRepositoryInterface,
) CreateProductUseCaseInterface {

	return &CreateProductUseCase{
		productRepository: productRepository,
	}
}

func (pdr *CreateProductUseCase) Execute(
	name, description string,
	image *multipart.FileHeader,
	price float64,
	quantity int,
) (*entities.Product, error) {

	db := database.NewDatabaseConnection()
	userRepository := repositories.NewUserRepository(db)
	findAllUsersUseCase := NewFindAllUsersUseCase(userRepository)

	newPrice := int(price * 100)
	p, err := entities.NewProduct(name, description, "", newPrice, quantity)

	if err != nil {
		return nil, err
	}

	if image != nil {
		cdn, err := configs.NewCloudinaryConfig()

		if err != nil {
			return nil, err
		}
		file, _ := image.Open()
		ctx := context.Background()
		cld, _ := cdn.Cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: p.ID})
		p.ImageURL = cld.SecureURL
	}

	if err := pdr.productRepository.Create(p); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		users, _ := findAllUsersUseCase.Execute()
		client := smtp.NewSMTPClient()

		defer wg.Done()

		for _, u := range users {
			fmt.Println(u.Email)
			client.SendNewProductAdded(u.Email)
		}

	}()

	wg.Wait()

	return p, nil
}
