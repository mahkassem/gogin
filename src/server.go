package server

import (
	"main/src/routers"
	"os"

	"github.com/gin-gonic/gin"
)

var Application *gin.Engine

func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	Application = initEngine()
	go postInitialize()
	listen()
}

func postInitialize() {
	setupRouters()
}

func setupRouters() {
	routers.SetupPingRouter(Application)
	routers.SetupUserRouter(Application)
}

func initEngine() *gin.Engine {
	return gin.Default()
}
func listen() {
	Application.Run("127.0.0.1:" + os.Getenv("PORT"))
}
