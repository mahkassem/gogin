package controllers

import (
	"main/src/config"
	"main/src/utilities"
	"strconv"

	"github.com/gin-gonic/gin"
)

func bindBody[T interface{}](c *gin.Context) (T, error) {
	var obj T
	err := c.Bind(&obj)
	return obj, err
}
func bindBodyWithResponse[T interface{}](c *gin.Context) (T, bool) {
	var obj T
	err := c.Bind(&obj)
	var ok = err == nil
	if !ok {
		respond(c, 400, nil, "Invalid request", &config.Error{Error: err})
	}
	return obj, ok
}

func getIntParam(c *gin.Context, param string) int {
	value, _ := strconv.Atoi(c.Param(param))
	return value
}

func verifyAndRespond(c *gin.Context, responseCode int, data interface{}, err error, successMessage, errorMessage string) {

	var Message = utilities.Ternary(err == nil, successMessage, errorMessage)
	var error = utilities.Ternary(err != nil, &config.Error{Error: err}, nil)

	respond(c, responseCode, data, Message, error)
}

func respond(c *gin.Context, responseCode int, data interface{}, message string, err *config.Error) {
	if c == nil {
		return
	}
	c.JSON(responseCode, config.Response{
		Message: message,
		Data:    data,
		Error:   err,
	})
}
