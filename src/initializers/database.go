package initializers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() gorm.DB {

	var err error
	connection := os.Getenv("DATABASE_URL")
	DB, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	return *DB
}
