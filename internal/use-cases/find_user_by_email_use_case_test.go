package use_cases_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	use_cases "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/use-cases"
	"github.com/stretchr/testify/assert"
)

func TestShouldFindUserByEmail(t *testing.T) {

	db := database.NewDatabaseConnection()
	userRepo := repositories.NewUserRepository(db)
	createUserUseCase := use_cases.NewCreateUserUseCase(userRepo)
	findUserByEmailUseCase := use_cases.NewFindUserByEmailUseCase(userRepo)

	createUserUseCase.Execute("email", "username", "password")

	u, err := findUserByEmailUseCase.Execute("email")

	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.Name, "username")
	assert.Equal(t, u.Email, "email")
	assert.NotNil(t, u.Password)
}
