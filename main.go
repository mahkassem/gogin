package main

import (
	server "main/src"
	"main/src/initializers"
)

func Init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.LoadConfig()
}

func main() {
	Init()
	server.StartServer()
}
