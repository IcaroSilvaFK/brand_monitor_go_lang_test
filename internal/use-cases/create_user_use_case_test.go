package use_cases_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	use_cases "github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/use-cases"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateUser(t *testing.T) {

	db := database.NewDatabaseConnection()
	userRepo := repositories.NewUserRepository(db)
	createUserUseCase := use_cases.NewCreateUserUseCase(userRepo)

	u, err := createUserUseCase.Execute("email", "username", "password")

	assert.Nil(t, err)
	assert.Equal(t, u.Email, "email")
	assert.Equal(t, u.Name, "username")
	assert.NotNil(t, u.Password)

}
