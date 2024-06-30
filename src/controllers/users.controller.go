package controllers

import (
	"main/src/database/models"
	"main/src/godash"
	"main/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*BaseController) CreateUser(c *gin.Context) {

	user, ok := bindBodyWithResponse[models.User](c)
	if !ok {
		return
	}

	result := services.CreateUser(&user)

	verifyAndRespond(c, http.StatusOK, user, result.Error, "Created user successfully", "Error creating user")
}

func (*BaseController) DeleteUser(c *gin.Context) {
	id := getIntParam(c, "id")

	result := services.DeleteUser(id)

	verifyAndRespond(c, http.StatusOK, nil, result.Error, "Deleted user successfully", "Error deleting user")
}

func (*BaseController) GetAllUsers(c *gin.Context) {
	users, result := services.GetAllUsers()

	verifyAndRespond(c, http.StatusOK, users, result.Error, "Showing all users", "Error showing all users")
}

func (*BaseController) GetUserById(c *gin.Context) {
	id := getIntParam(c, "id")
	user, result, _ := services.GetUserById(id)

	verifyAndRespond(c, http.StatusOK, user, result.Error, "User was found", "User not found")
}

func (*BaseController) UpdateUser(c *gin.Context) {
	id := getIntParam(c, "id")

	body, ok := bindBodyWithResponse[models.User](c)
	if !ok {
		return
	}

	user, result, found := services.UpdateUser(id, body)

	var message = godash.If(found, "User updated successfully", "User not found")

	verifyAndRespond(c, http.StatusOK, user, result.Error, message, "Error updating user")
}

/*
	1. GET / index

	2. GET / new
	3. POST / create

	4. GET /:id show

	7. DELETE /:id destroy


	5. GET /:id/edit edit
	6. PUT /:id update
*/
