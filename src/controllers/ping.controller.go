package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (test *PingController) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"time": time.Now().String()})
}
