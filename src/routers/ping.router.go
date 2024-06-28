package routers

import (
	"main/src/controllers"
	"main/src/utilities"
)

func SetupPingRouter() {
	utilities.SetupBaseRoute("ping", &controllers.PingController{})
}
