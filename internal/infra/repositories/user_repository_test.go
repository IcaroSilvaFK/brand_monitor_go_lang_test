package repositories_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/database"
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/infra/repositories"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateUserInDatabase(t *testing.T) {

	u := entities.NewUser("icaro", "Icaro", "123456")
	conn := database.NewDatabaseConnection()
	udb := repositories.NewUserRepository(conn)

	err := udb.Create(u)

	assert.Nil(t, err)
}

func TestShouldCreateFindInDatabase(t *testing.T) {

	u := entities.NewUser("icaro", "Icaro", "123456")
	conn := database.NewDatabaseConnection()
	udb := repositories.NewUserRepository(conn)

	udb.Create(u)

	u, err := udb.FindByEmail("icaro")

	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.Name, "Icaro")
	assert.Equal(t, u.Email, "icaro")
	assert.NotNil(t, u.Password)
}

func TestShouldCreateDeleteInDatabase(t *testing.T) {

	u := entities.NewUser("testedel", "Icaro", "123456")
	conn := database.NewDatabaseConnection()
	udb := repositories.NewUserRepository(conn)

	udb.Create(u)

	err := udb.Delete(u.ID)

	assert.Nil(t, err)
}
