package initializers

import "main/src/database"

func ConnectDatabase() {
	database.Connect()
}
