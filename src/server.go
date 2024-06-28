package server

import (
	"main/src/config"
	"main/src/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	config.Application = initEngine()
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
	config.Application.Run("127.0.0.1:" + os.Getenv("PORT"))
}
