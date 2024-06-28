package routers

import (
	controllers "main/src/controllers/users_controller"
	"main/src/utilities"
)

func SetupUserRouter() {
	utilities.SetupBaseRoute("user", &controllers.UserController{})
}
