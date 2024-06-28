package users_controller

import (
	"main/src/config"
	"main/src/database/models"
	"main/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (*UserController) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	body := models.User{}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, config.Response{Message: "Invalid request", Data: nil, Error: nil})
		return
	}

	user, result, found := services.UpdateUser(id, body)

	if !found {
		c.JSON(http.StatusNotFound, config.Response{Message: "User not found", Data: nil, Error: nil})
		return
	}

	Message, Error := VerifyResult(result, "Updated user successfully", "Error updating user")

	c.JSON(http.StatusOK, config.Response{Message: Message, Data: user, Error: Error})
}
