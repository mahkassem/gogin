package routers

import (
	"fmt"
	"main/src/middlewares"
	"main/src/utilities"

	"github.com/gin-gonic/gin"
)

func AddRoute(Application *gin.Engine, path string, parentMiddlewares []string, route Route, handler gin.HandlerFunc) {
	fmt.Println(" - Adding route: " + route.Method + " " + path + " ---> controllers." + route.Handler)
	parentMiddlewares = utilities.Ternary(utilities.StringInSlice("*", route.SkipMiddlewares), route.Middlewares, append(parentMiddlewares, route.Middlewares...))
	handlers := append(middlewares.HandleRouteMiddleware(parentMiddlewares, route.SkipMiddlewares), handler)
	switch route.Method {
	case "GET":
		Application.GET(path, handlers...)
	case "POST":
		Application.POST(path, handlers...)
	case "PUT":
		Application.PUT(path, handlers...)
	case "DELETE":
		Application.DELETE(path, handlers...)
	default:
		panic("Unsupported HTTP method: " + route.Method)
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
			fullPath,
			routes[routeName].Middlewares,
			route,
			handler,
		)
	}
}
