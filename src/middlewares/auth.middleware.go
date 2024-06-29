package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (*Middleware) AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Authentication logic here
		fmt.Println("Checking authentication...")
		fmt.Println(c.Request.URL)
		c.Next()
	}
}
