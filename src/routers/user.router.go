package routers

import (
	"main/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(Application *gin.Engine) {
	SetupBaseRoute(Application, "user", &controllers.UserController{})
}
