package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"promo/src/models"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	var err error
	DB, err = gorm.Open(mysql.Open(os.Getenv("CONN_STRING")), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
