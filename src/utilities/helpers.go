package utilities

import (
	"fmt"
	"main/src/config"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
)

func GetMethodByName(name string, controller any) (gin.HandlerFunc, error) {
	ctrl := reflect.ValueOf(controller)
	method := ctrl.MethodByName(name)
	if !method.IsValid() {
		return nil, fmt.Errorf("no function with name %s found", name)
	}
	handlerFunc, ok := method.Interface().(func(*gin.Context))
	if !ok {
		return nil, fmt.Errorf("function %s has wrong signature", name)
	}
	return gin.HandlerFunc(handlerFunc), nil
}

func AddRoute(method string, path string, handlerName string, handler gin.HandlerFunc) {
	fmt.Println(" - Adding route: " + method + " " + path + " ---> controllers." + handlerName)
	switch method {
	case "GET":
		config.Application.GET(path, handler)
	case "POST":
		config.Application.POST(path, handler)
	case "PUT":
		config.Application.PUT(path, handler)
	case "DELETE":
		config.Application.DELETE(path, handler)
	default:
		panic("Unsupported HTTP method: " + method)
	}
}

func SetupBaseRoute(routeName string, controller any) {
	fmt.Println("--> Setting up route: " + routeName)
	for _, route := range config.Configuration.Routes[routeName].Routes {
		handler, err := GetMethodByName(route.Handler, controller)
		if err != nil {
			fmt.Println(
				"Error setting up route: " + routeName + " - " + route.Handler + " - " + err.Error(),
			)
			continue
		}
		fullPath := config.Configuration.Routes[routeName].Path + route.Path
		AddRoute(
			route.Method,
			fullPath,
			route.Handler,
			handler,
		)
	}
}

func DetectMigration() bool {
	if len(os.Args) < 2 {
		return false
	}
	command := os.Args[1]
	return command == "migrate" || command == "-m" || command == "--migrate"
}

// func HandleCommandArguments() {
// 	if len(os.Args) < 2 {
// 		return
// 	}
// 	command := os.Args[1]
// 	switch command {
// 	case "migrate", "-m", "--migrate":
// 		initializers.DB.AutoMigrate(&models.Entity{})
// 		// migrations.PerformMigration()
// 		os.Exit(0)
// 	default:
// 		fmt.Println("Invalid command!, Argument[1]: " + os.Args[1])
// 	}
// }
