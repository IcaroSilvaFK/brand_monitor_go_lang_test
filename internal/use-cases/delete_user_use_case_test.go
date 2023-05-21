package use_cases_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	use_cases "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/use-cases"
	"github.com/stretchr/testify/assert"
)

func TestShouldDeleteUser(t *testing.T) {

	db := database.NewDatabaseConnection()
	userRepo := repositories.NewUserRepository(db)
	createUserUseCase := use_cases.NewCreateUserUseCase(userRepo)
	deleteUserUseCase := use_cases.NewDeleteUserUseCase(userRepo)

	u, _ := createUserUseCase.Execute("email", "username", "password")

	err := deleteUserUseCase.Execute(u.ID)

	assert.Nil(t, err)
}
