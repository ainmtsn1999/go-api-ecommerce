package db

import (
	"fmt"

	"github.com/ainmtsn1999/go-api-ecommerce/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_HOST = config.GetEnvVariable("DB_HOST")
	DB_PORT = config.GetEnvVariable("DB_PORT")
	DB_USER = config.GetEnvVariable("DB_USER")
	DB_PASS = config.GetEnvVariable("DB_PASS")
	DB_NAME = config.GetEnvVariable("DB_NAME")
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME,
	)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = gormDB
}
