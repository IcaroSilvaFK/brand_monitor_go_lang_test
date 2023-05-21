package use_cases_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	use_cases "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/use-cases"
	"github.com/stretchr/testify/assert"
)

func TestShouldFindAllUsers(t *testing.T) {

	db := database.NewDatabaseConnection()
	userRepo := repositories.NewUserRepository(db)
	createUserUseCase := use_cases.NewFindAllUsersUseCase(userRepo)

	u, err := createUserUseCase.Execute()

	assert.Nil(t, err)
	assert.NotNil(t, u)
}
