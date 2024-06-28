package controllers

import (
	"main/src/database/models"
	"main/src/services"
	"main/src/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (*UserController) CreateUser(c *gin.Context) {

	user, ok := bindBodyWithResponse[models.User](c)
	if !ok {
		return
	}

	result := services.CreateUser(&user)

	verifyAndRespond(c, http.StatusOK, user, result.Error, "Created user successfully", "Error creating user")
}

func (*UserController) DeleteUser(c *gin.Context) {
	id := getIntParam(c, "id")

	result := services.DeleteUser(id)

	verifyAndRespond(c, http.StatusOK, nil, result.Error, "Deleted user successfully", "Error deleting user")
}

func (*UserController) GetAllUsers(c *gin.Context) {
	users, result := services.GetAllUsers()

	verifyAndRespond(c, http.StatusOK, users, result.Error, "Showing all users", "Error showing all users")
}

func (*UserController) GetUserById(c *gin.Context) {
	id := getIntParam(c, "id")
	user, result, _ := services.GetUserById(id)

	verifyAndRespond(c, http.StatusOK, user, result.Error, "User was found", "User not found")
}

func (*UserController) UpdateUser(c *gin.Context) {
	id := getIntParam(c, "id")

	body, ok := bindBodyWithResponse[models.User](c)
	if !ok {
		return
	}

	user, result, found := services.UpdateUser(id, body)

	var message = utilities.Ternary(found, "User updated successfully", "User not found")

	verifyAndRespond(c, http.StatusOK, user, result.Error, message, "Error updating user")
}
