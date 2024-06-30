package middlewares

import (
	"fmt"
	"main/src/godash"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Middleware struct{}

func HandleRouteMiddleware(middlewareNames []string, skip []string) []gin.HandlerFunc {
	middlewares := []gin.HandlerFunc{(&Middleware{}).BaseMiddleware()}
	computed := []string{}
	for _, middlewareName := range middlewareNames {
		if godash.StringInSlice(middlewareName, skip) || godash.StringInSlice(middlewareName, computed) {
			continue
		}
		handler, err := getMiddlewareByName(middlewareName, &Middleware{})
		if err != nil {
			fmt.Println(err.Error())
			panic("Middleware not found: " + middlewareName)
		}
		middlewares = append(middlewares, handler)
		computed = append(computed, middlewareName)
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
