package users_controller

import (
	"main/src/config"

	"gorm.io/gorm"
)

type UserController struct{}

func VerifyResult(result *gorm.DB, successMessage, errorMessage string) (string, *config.Error) {
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
