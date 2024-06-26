package main

import (
	server "main/src"
	"main/src/initializers"
)

func init() {
	initializers.Initialize()
}

func main() {
	server.StartServer()
}
