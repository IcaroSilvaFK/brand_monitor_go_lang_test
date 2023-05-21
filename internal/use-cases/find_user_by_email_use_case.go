package use_cases

import (
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
)

type FindUserByEmailUseCase struct {
	userRepo repositories.UserRepositoryInterface
}

type FindUserByEmailUseCaseInterface interface {
	Execute(email string) (*entities.User, error)
}

func NewFindUserByEmailUseCase(userRepo repositories.UserRepositoryInterface) FindUserByEmailUseCaseInterface {
	return &FindUserByEmailUseCase{
		userRepo: userRepo,
	}
}

func (useCase *FindUserByEmailUseCase) Execute(email string) (*entities.User, error) {
	return useCase.userRepo.FindByEmail(email)
}
