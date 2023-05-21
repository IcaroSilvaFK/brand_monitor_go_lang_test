package use_cases

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/configs"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/queues"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
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

	queueConnection := queues.NewSendEmailQueue()

	taskQueue, _ := queueConnection.OpenQueue("send_order_emails")
	err = taskQueue.StartConsuming(10, time.Second)

	go func() {
		users, _ := findAllUsersUseCase.Execute()
		for _, u := range users {
			taskQueue.Publish(u.Email)
		}
	}()

	fmt.Println("Já passow")
	if err != nil {
		fmt.Println("Error on consume queue:", err.Error())
	}

	sendEmailConsumer := queues.NewSendAllEmailsQueue()

	_, err = taskQueue.AddConsumer("send_order_emails", sendEmailConsumer)

	if err != nil {
		fmt.Println("Error on consume queue:", err.Error())
	}

	return p, nil
}
