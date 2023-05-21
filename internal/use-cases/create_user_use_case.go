package use_cases

import (
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
)

type CreateUserUseCase struct {
	userRepo repositories.UserRepositoryInterface
}

type CreateUserUseCaseInterface interface {
	Execute(email, username, password string) (*entities.User, error)
}

func NewCreateUserUseCase(userRepo repositories.UserRepositoryInterface) CreateUserUseCaseInterface {
	return &CreateUserUseCase{
		userRepo: userRepo,
	}
}

func (useCase *CreateUserUseCase) Execute(email, username, password string) (*entities.User, error) {

	u := entities.NewUser(email, username, password)

	if err := useCase.userRepo.Create(u); err != nil {
		return nil, err
	}

	return u, nil
}
