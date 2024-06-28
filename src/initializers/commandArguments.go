package initializers

import (
	"fmt"
	"main/src/database"
	"os"
)

func HandleCommandArguments(migrateFile bool) {
	if migrateFile {
		database.MIGRATE = true
	}
	if len(os.Args) < 2 {
		return
	}
	for _, command := range os.Args[1:] {
		switch command {
		case "-d", "--drop-tables":
			database.DROP_TABLES = true
			database.REGULAR_STARTUP = false
		case "-m", "--migrate":
			database.MIGRATE = true
			database.REGULAR_STARTUP = false
		case "-md", "-dm":
			database.DROP_TABLES = true
			database.MIGRATE = true
			database.REGULAR_STARTUP = false
		case "-n":
			database.REGULAR_STARTUP = true
		case "-mdn":
			database.DROP_TABLES = true
			database.MIGRATE = true
			database.REGULAR_STARTUP = true
		default:
			fmt.Println("Invalid command:", command)
		}
	}
}
