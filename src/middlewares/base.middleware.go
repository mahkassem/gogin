package middlewares

import (
	"fmt"
	"main/src/utilities"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Middleware struct{}

func HandleRouteMiddleware(middlewareNames []string, skip []string) []gin.HandlerFunc {
	middlewares := []gin.HandlerFunc{(&Middleware{}).BaseMiddleware()}
	for _, middlewareName := range middlewareNames {
		if utilities.StringInSlice(middlewareName, skip) {
			continue
		}
		handler, err := getMiddlewareByName(middlewareName, &Middleware{})
		if err != nil {
			fmt.Println(err.Error())
			panic("Middleware not found: " + middlewareName)
		}
		middlewares = append(middlewares, handler)
	}
	return middlewares
}

func (*Middleware) BaseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func (*Middleware) TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Testing middleware...")
		c.Next()
	}
}

func getMiddlewareByName(name string, controller any) (gin.HandlerFunc, error) {
	ctrl := reflect.ValueOf(controller)
	method := ctrl.MethodByName(name)
	if !method.IsValid() {
		return nil, fmt.Errorf("no middleware with name %s found", name)
	}
	handlerFunc, ok := method.Interface().(func() gin.HandlerFunc)
	if !ok {
		return nil, fmt.Errorf("function %s has wrong signature", name)
	}
	return handlerFunc(), nil
}