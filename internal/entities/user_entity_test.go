package entities_test

import (
	"testing"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewUser(t *testing.T) {

	u := entities.NewUser("email", "name", "password")

	assert.Equal(t, u.Email, "email")
	assert.Equal(t, u.Name, "name")
	assert.NotNil(t, u.Password)
}
