package utilities

import (
	"fmt"
	"main/src/config"
	"main/src/database/models"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func Assign(_1 interface{}, _2 []interface{}) {
	// Reflect value of the existing user
	v := reflect.ValueOf(_1).Elem()

	// Iterate over the request body and update fields dynamically
	for key, value := range _2 {
		fmt.Println(key, value)
		field := v.FieldByName(string(key))
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(value))
		}
	}
}

func HashPassword(password string) string {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func ComparePassword(hashedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func AssignDataToUser(data models.User, user *models.User) {
	if data.Email != "" {
		user.Email = data.Email
	}

	if data.Username != "" {
		user.Username = data.Username
	}
}
