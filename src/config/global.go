package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Configuration Config
var Application *gin.Engine
var DB *gorm.DB
