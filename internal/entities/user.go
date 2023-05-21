package entities

import (
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/pkg/utils"
	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser(email, name, password string) *User {

	hash, err := utils.GenerateHash(password)

	if err != nil {
		return nil
	}

	return &User{
		ID:       uuid.NewString(),
		Email:    email,
		Name:     name,
		Password: hash,
	}
}

func (u *User) VerifyPassword(pass string) error {

	if err := utils.VerifyHash(pass, u.Password); err != nil {
		return err
	}

	return nil
}
