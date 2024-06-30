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
	setupRendering()
	go postInitialize()
	listen()
}

func setupRendering() {
	Application.LoadHTMLGlob("./views/**/**/*")
	Application.Static("/public", "./public")
	// TEST PURPOSES ONLY
	Application.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.tmpl", gin.H{
			"Title":   "Unique hamada",
			"Link":    "https://www.google.com",
			"Website": "Google",
		})
	})
}

func postInitialize() {
	setupRouters()
}

func setupRouters() {
	routers.SetupRouters(Application)
}

func initEngine() *gin.Engine {
	return gin.Default()
}
func listen() {
	Application.Run("127.0.0.1:" + os.Getenv("PORT"))
}
