package users_controller

import (
	"main/src/config"
	"main/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*UserController) GetAllUsers(c *gin.Context) {
	users, result := services.GetAllUsers()

	Message, Error := VerifyResult(result, "Showing all users", "Error showing all users")
	c.JSON(http.StatusOK, config.Response{Message: Message, Data: users, Error: Error})
}
