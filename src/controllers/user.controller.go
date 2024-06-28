package controllers

import (
	"main/src/config"
	"main/src/database/models"
	"main/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct{}

func (*UserController) CreateUser(c *gin.Context) {

	user := models.User{}
	c.Bind(&user)

	result := services.CreateUser(&user)

	Message, Error := VerifyResult(result, "Created user successfully", "Error creating user")

	c.JSON(http.StatusOK, config.Response{Message: Message, Data: user, Error: Error})
}

func (*UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result := services.DeleteUser(id)

	Message, Error := VerifyResult(result, "Deleted user successfully", "Error deleting user")

	c.JSON(http.StatusOK, config.Response{Message: Message, Data: nil, Error: Error})
}

func (*UserController) GetAllUsers(c *gin.Context) {
	users, result := services.GetAllUsers()

	Message, Error := VerifyResult(result, "Showing all users", "Error showing all users")
	c.JSON(http.StatusOK, config.Response{Message: Message, Data: users, Error: Error})
}

func (*UserController) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, result, _ := services.GetUserById(id)

	Message, Error := VerifyResult(result, "User was found", "User not found")

	c.JSON(http.StatusOK, config.Response{Message: Message, Data: user, Error: Error})
}

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

func VerifyResult(result *gorm.DB, successMessage, errorMessage string) (string, *config.Error) {
	var Message string
	var Error *config.Error
	Message = successMessage
	if result.Error != nil {
		Error = &config.Error{Message: Message, Error: result.Error}
		Message = errorMessage
	}

	return Message, Error
}
