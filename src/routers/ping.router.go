package routers

import (
	"main/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPingRouter(Application *gin.Engine) {
	SetupBaseRoute(Application, "ping", &controllers.PingController{})
}
