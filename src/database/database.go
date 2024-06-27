package database

import (
	"main/src/config"
	"main/src/database/migrations"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() gorm.DB {

	var err error
	// MYSQL
	connection := os.Getenv("DATABASE_URL")
	config.DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	// SQLITE
	// DB, err := gorm.Open(sqlite.Open("/tmp/database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	if config.DROP_TABLES {
		migrations.DropTables()
	}

	if config.MIGRATE {
		migrations.PerformMigration()
	}

	if !config.REGULAR_STARTUP && (config.DROP_TABLES || config.MIGRATE) {
		os.Exit(0)
	}

	return *config.DB
}
