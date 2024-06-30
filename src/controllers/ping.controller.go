package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (test *BaseController) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"time": time.Now().String()})
}
