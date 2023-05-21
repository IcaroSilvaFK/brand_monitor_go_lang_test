package use_cases

import (
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
)

type FindAllUsersUseCase struct {
	userRepo repositories.UserRepositoryInterface
}

type FindAllUsersUseCaseInterface interface {
	Execute() ([]*entities.User, error)
}

func NewFindAllUsersUseCase(
	userRepo repositories.UserRepositoryInterface,
) FindAllUsersUseCaseInterface {

	return &FindAllUsersUseCase{
		userRepo: userRepo,
	}
}

func (urp *FindAllUsersUseCase) Execute() ([]*entities.User, error) {

	users, err := urp.userRepo.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
