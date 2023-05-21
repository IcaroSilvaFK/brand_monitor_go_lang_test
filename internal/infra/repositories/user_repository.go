package repositories

import (
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

type UserRepositoryInterface interface {
	Create(u *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	Delete(id string) error
	FindAll() ([]*entities.User, error)
}

func NewUserRepository(conn *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		conn: conn,
	}
}

func (udb *UserRepository) Create(u *entities.User) error {
	return udb.conn.Create(u).Error
}

func (udb *UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	tx := udb.conn.Find(&user, "email = ?", email)

	return &user, tx.Error
}

func (udb *UserRepository) FindAll() ([]*entities.User, error) {
	var users []*entities.User

	if err := udb.conn.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (udb *UserRepository) Delete(id string) error {

	err := udb.conn.Where("id = ?", id).Delete(&entities.User{}).Error

	return err
}
