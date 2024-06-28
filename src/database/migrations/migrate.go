package migrations

import (
	"fmt"
	"main/src/database/models"

	"gorm.io/gorm"
)

func PerformMigration(DB *gorm.DB) {
	fmt.Println("Performing database migration...")

	DB.AutoMigrate(
		// &models.Entity{},
		&models.User{},
	)

	fmt.Println("Migration complete!")
}

func DropTables(DB *gorm.DB) {
	fmt.Println("Dropping tables...")

	DB.Migrator().DropTable(
		&models.User{},
	)

	fmt.Println("Tables dropped!")
}
