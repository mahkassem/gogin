package database

import (
	"main/src/database/migrations"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// STATE
var DROP_TABLES = false
var MIGRATE = false
var REGULAR_STARTUP = true

func Connect() gorm.DB {

	var err error
	// MYSQL
	connection := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	// SQLITE
	// DB, err := gorm.Open(sqlite.Open("/tmp/database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	if DROP_TABLES {
		migrations.DropTables(DB)
	}

	if MIGRATE {
		migrations.PerformMigration(DB)
	}

	if !REGULAR_STARTUP && (DROP_TABLES || MIGRATE) {
		os.Exit(0)
	}

	return *DB
}
