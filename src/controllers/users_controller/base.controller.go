package users_controller

import (
	"main/src/config"

	"gorm.io/gorm"
)

type UserController struct{}

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
