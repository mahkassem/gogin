package users_controller

import (
	"main/src/config"
	"main/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (*UserController) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, result, _ := services.GetUserById(id)

	Message, Error := VerifyResult(result, "User was found", "User not found")

	c.JSON(http.StatusOK, config.Response{Message: Message, Data: user, Error: Error})
}
