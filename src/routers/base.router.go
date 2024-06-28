package routers

import (
	"fmt"
	"main/src/utilities"

	"github.com/gin-gonic/gin"
)

func AddRoute(Application *gin.Engine, method string, path string, handlerName string, handler gin.HandlerFunc) {
	fmt.Println(" - Adding route: " + method + " " + path + " ---> controllers." + handlerName)
	switch method {
	case "GET":
		Application.GET(path, handler)
	case "POST":
		Application.POST(path, handler)
	case "PUT":
		Application.PUT(path, handler)
	case "DELETE":
		Application.DELETE(path, handler)
	default:
		panic("Unsupported HTTP method: " + method)
	}
}

func SetupBaseRoute(Application *gin.Engine, routeName string, controller any) {
	fmt.Println("--> Setting up route: " + routeName)
	for _, route := range routes[routeName].Routes {
		handler, err := utilities.GetMethodByName(route.Handler, controller)
		if err != nil {
			fmt.Println(
				"Error setting up route: " + routeName + " - " + route.Handler + " - " + err.Error(),
			)
			continue
		}
		fullPath := routes[routeName].Path + route.Path
		AddRoute(
			Application,
			route.Method,
			fullPath,
			route.Handler,
			handler,
		)
	}
}
