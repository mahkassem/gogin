package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Configuration map[string]RouteConfig
var Application *gin.Engine
var DB *gorm.DB

// STATE
var DROP_TABLES = false
var MIGRATE = false
var REGULAR_STARTUP = true
