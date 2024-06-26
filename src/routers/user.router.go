package routers

import (
	"main/src/controllers"
	"main/src/utilities"
)

func SetupUserRouter() {
	utilities.SetupBaseRoute("user", &controllers.UserController{})
}
