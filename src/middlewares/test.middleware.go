package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (*Middleware) Test2Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Testing 2 middleware...")
		c.Next()
	}
}
