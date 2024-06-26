package controllers

import (
	"main/src/config"
	"main/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
