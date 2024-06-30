package routers

import (
	"fmt"
	"main/src/controllers"
	"main/src/godash"
	"main/src/middlewares"
	s "strings"

	"github.com/gin-gonic/gin"
)

func AddRoute(Application *gin.Engine, parentMiddlewares []string, route Route, handler gin.HandlerFunc) {
	path := route.Path
	// fmt.Println(" - Adding route: " + route.Method + " " + path + " ---> controllers." + route.Handler)
	parentMiddlewares = godash.If(godash.StringInSlice("*", route.SkipMiddlewares), route.Middlewares, append(parentMiddlewares, route.Middlewares...))
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

func SetupRouters(Application *gin.Engine) {
	for _, route := range routes {
		SetupRoute(Application, route, 1)
	}
}

func SetupRoute(Application *gin.Engine, route Route, count int) {
	fmt.Println(" " + s.Repeat("-", count) + " Setting up route: " + route.Path)
	count += 1
	for _, subRoute := range route.Routes {
		if len(subRoute.Routes) == 0 {
			SetupSubRoute(Application, route.Path, subRoute, count)
		} else {
			subRoute.Path = route.Path + subRoute.Path
			subRoute.Middlewares = append(route.Middlewares, subRoute.Middlewares...)
			subRoute.SkipMiddlewares = append(route.SkipMiddlewares, subRoute.SkipMiddlewares...)
			SetupRoute(Application, subRoute, count)
		}
	}
}

func SetupSubRoute(Application *gin.Engine, path string, route Route, count int) {
	route.Path = path + route.Path
	fmt.Println(" " + s.Repeat("-", count) + " Adding sub-route: " + route.Method + " " + route.Path + " ---> controllers." + route.Handler)
	handler, err := godash.GetMethodByName(route.Handler, &controllers.BaseController{})
	if err != nil {
		fmt.Println(
			"Error setting up route: " + route.Path + " - " + route.Handler + " - " + err.Error(),
		)
	}
	AddRoute(
		Application,
		route.Middlewares,
		route,
		handler,
	)
}
