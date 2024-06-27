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

func (*UserController) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, ok := services.GetUserById(id)
	if !ok {
		c.JSON(http.StatusNotFound, config.Response{Message: "User not found", Data: nil, Error: nil})
		return
	}
	c.JSON(http.StatusOK, config.Response{Message: "User found successfully", Data: user, Error: nil})
}

func (*UserController) ShowAllUsers(c *gin.Context) {
	users, ok := services.GetAllUsers()
	if !ok {
		c.JSON(http.StatusNotFound, config.Response{Message: "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, config.Response{Message: "Successfully fetched all users", Data: users, Error: nil})
}

func (*UserController) CreateUser(c *gin.Context) {

	var body struct {
		Username string
		Email    string
	}
	c.Bind(&body)

	user := models.User{Username: body.Username, Email: body.Email}
	result := config.DB.Create(&user)

	Message, Error := verifyResult(result, "Created user successfully", "Error creating user")

	c.JSON(http.StatusOK, config.Response{Message: Message, Data: user, Error: Error})
}

func (*UserController) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var body struct {
		Username *string
		Email    *string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, config.Response{Message: "Invalid request", Data: nil, Error: nil})
		return
	}

	user, ok := services.GetUserById(id)

	if !ok {
		c.JSON(http.StatusNotFound, config.Response{Message: "User not found", Data: nil, Error: nil})
		return
	}

	assignBodyToUser(body, &user)

	result := config.DB.Updates(&user)

	Message, Error := verifyResult(result, "Updated user successfully", "Error updating user")

	c.JSON(http.StatusOK, config.Response{Message: Message, Data: user, Error: Error})
}

func (*UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, ok := services.GetUserById(id)

	if !ok {
		c.JSON(http.StatusNotFound, config.Response{Message: "User not found", Data: nil, Error: nil})
		return
	}

	result := config.DB.Delete(&user)

	Message, Error := verifyResult(result, "Deleted user successfully", "Error deleting user")

	c.JSON(http.StatusOK, config.Response{Message: Message, Data: user, Error: Error})
}

func verifyResult(result *gorm.DB, successMessage, errorMessage string) (string, *config.Error) {
	var Message string
	if result.Error != nil {
		Message = errorMessage
	} else {
		Message = successMessage
	}

	var Error *config.Error
	if result.Error != nil {
		Error = &config.Error{Message: Message, Error: result.Error}
	}

	return Message, Error
}

func assignBodyToUser(body struct {
	Username *string
	Email    *string
}, user *models.User) {
	if body.Email != nil {
		user.Email = *body.Email
	}

	if body.Username != nil {
		user.Username = *body.Username
	}
}
