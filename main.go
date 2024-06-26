package main

import (
	server "main/src"
	"main/src/initializers"
	"main/src/utilities"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB(utilities.DetectMigration())
	initializers.LoadConfig()
}

func main() {
	server.StartServer()
}
