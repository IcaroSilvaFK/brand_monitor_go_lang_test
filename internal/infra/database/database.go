package database

import (
	"log"

	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_DATABASE = "DB_DATABASE"
	DB_USERNAME = "DB_USERNAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
)

func NewDatabaseConnection() *gorm.DB {

	conn, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	autoMigrate(conn)

	return conn
}

func autoMigrate(db *gorm.DB) {

	db.AutoMigrate(&entities.User{}, &entities.Product{})

}
