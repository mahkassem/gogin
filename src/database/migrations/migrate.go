package migrations

import (
	"fmt"
	"main/src/config"
	"main/src/database/models"
)

func PerformMigration() {
	fmt.Println("Performing database migration...")

	config.DB.AutoMigrate(
		// &models.Entity{},
		&models.User{},
	)

	fmt.Println("Migration complete!")
}

func DropTables() {
	fmt.Println("Dropping tables...")

	config.DB.Migrator().DropTable(
		&models.User{},
	)

	fmt.Println("Tables dropped!")
}
