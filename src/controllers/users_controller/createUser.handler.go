package users_controller

import (
	"main/src/config"
	"main/src/database/models"
	"main/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*UserController) CreateUser(c *gin.Context) {

	user := models.User{}
	c.Bind(&user)

	result := services.CreateUser(&user)

	Message, Error := VerifyResult(result, "Created user successfully", "Error creating user")

	c.JSON(http.StatusOK, config.Response{Message: Message, Data: user, Error: Error})
}
