package services

import (
	"main/src/config"
	"main/src/database/models"
	"main/src/utilities"

	"gorm.io/gorm"
)

func CreateUser(user *models.User) *gorm.DB {
	result := config.DB.Create(&user)
	return result
}

func DeleteUser(id int) *gorm.DB {
	result := config.DB.Delete(&models.User{}, id)
	return result
}

func UpdateUser(id int, data models.User) (*models.User, *gorm.DB, bool) {
	user, _, ok := GetUserById(id)
	if !ok {
		return nil, nil, false
	}
	utilities.AssignDataToUser(data, &user)
	result := config.DB.Updates(&user)
	if result.Error != nil {
		return nil, result, true
	}
	return &user, result, true
}

func GetUserById(id int) (models.User, *gorm.DB, bool) {
	user := models.User{}
	result := config.DB.First(&user, id)
	if user.ID != 0 {
		return user, result, true
	}
	return models.User{}, result, false
}

func GetAllUsers() ([]models.User, *gorm.DB) {
	users := []models.User{}
	result := config.DB.Find(&users)
	return users, result
}
