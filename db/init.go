package db

import (
	"github.com/pyrolass/book-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "host=127.0.0.1 user=las password=00778899 port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db = connection

	err = db.AutoMigrate(&models.Book{})

	if err != nil {
		panic("failed to migrate database")
	}
}

func GetDB() *gorm.DB {
	return db
}
