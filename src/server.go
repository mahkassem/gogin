package server

import (
	"main/src/config"
	"main/src/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	config.Application = initEngine()
	gin.SetMode(gin.ReleaseMode)
	go postInitialize()
	listen()
}

func postInitialize() {
	setupRouters()
}

func setupRouters() {
	routers.SetupPingRouter()
	routers.SetupUserRouter()
}

func initEngine() *gin.Engine {
	return gin.Default()
}
func listen() {
	config.Application.Run(":" + os.Getenv("PORT"))
}
