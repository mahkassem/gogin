package initializers

import (
	"main/src/utilities"
)

func Initialize() {
	LoadEnv()
	LoadConfig()
	utilities.DetectMigration()
}
