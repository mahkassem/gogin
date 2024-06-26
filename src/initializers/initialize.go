package initializers

import (
	"main/src/utilities"
)

func Initialize() {
	LoadEnv()
	LoadConfig()
	Connect(utilities.DetectMigration())
}
