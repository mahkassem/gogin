package migrations

import (
	"fmt"
	"main/src/config"
	"main/src/models"
)

func PerformMigration() {
	fmt.Println("Performing database migration...")

	config.DB.AutoMigrate(&models.Entity{}, &models.User{})

	fmt.Println("Migration complete!")
}
