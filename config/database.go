package config

import (
	"fmt"
	"golang-crud/models"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%vJakarta", ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_DATABASE, "%2F")

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect Database!")
	}

	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Book{})
	DB = db
	log.Println("Database connected!")
}
