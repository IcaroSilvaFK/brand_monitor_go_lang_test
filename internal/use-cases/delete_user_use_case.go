package use_cases

import "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"

type DeleteUserUseCase struct {
	userRepo repositories.UserRepositoryInterface
}

type DeleteUserUseCaseInterface interface {
	Execute(id string) error
}

func NewDeleteUserUseCase(
	userRepo repositories.UserRepositoryInterface,
) DeleteUserUseCaseInterface {

	return &DeleteUserUseCase{
		userRepo: userRepo,
	}
}

func (u *DeleteUserUseCase) Execute(id string) error {
	return u.userRepo.Delete(id)
}
