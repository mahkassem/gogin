package main

import (
	server "main/src"
	"main/src/initializers"
)

func init() {
	initializers.Initialize(false)
}

func main() {
	server.StartServer()
}
