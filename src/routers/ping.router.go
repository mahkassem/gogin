package routers

import (
	controllers "main/src/controllers/ping_controller"
	"main/src/utilities"
)

func SetupPingRouter() {
	utilities.SetupBaseRoute("ping", &controllers.PingController{})
}
