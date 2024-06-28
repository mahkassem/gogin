package users_controller

import (
	"main/src/config"
	"main/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (*UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result := services.DeleteUser(id)

	Message, Error := VerifyResult(result, "Deleted user successfully", "Error deleting user")

	c.JSON(http.StatusOK, config.Response{Message: Message, Data: nil, Error: Error})
}
