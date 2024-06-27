package initializers

import (
	"fmt"
	"main/src/config"
	"os"
)

func HandleCommandArguments(migrateFile bool) {
	if migrateFile {
		config.MIGRATE = true
	}
	if len(os.Args) < 2 {
		return
	}
	for _, command := range os.Args[1:] {
		switch command {
		case "-d", "--drop-tables":
			config.DROP_TABLES = true
			config.REGULAR_STARTUP = false
		case "-m", "--migrate":
			config.MIGRATE = true
			config.REGULAR_STARTUP = false
		case "-md", "-dm":
			config.DROP_TABLES = true
			config.MIGRATE = true
			config.REGULAR_STARTUP = false
		case "-n":
			config.REGULAR_STARTUP = true
		default:
			fmt.Println("Invalid command:", command)
		}
	}
}
