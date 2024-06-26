package initializers

import (
	"main/migrations"
	"main/src/config"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(migrate bool) gorm.DB {

	var err error
	// MYSQL
	connection := os.Getenv("DATABASE_URL")
	config.DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	// SQLITE
	// DB, err := gorm.Open(sqlite.Open("/tmp/database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	if migrate {
		migrations.PerformMigration()
		os.Exit(0)
	}

	return *config.DB
}
