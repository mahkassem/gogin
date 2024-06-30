package services

import (
	"main/src/database"
	"main/src/database/models"
	"main/src/godash"

	"gorm.io/gorm"
)

func CreateUser(user *models.User) *gorm.DB {
	result := database.DB.Create(&user)
	return result
}

func DeleteUser(id int) *gorm.DB {
	result := database.DB.Delete(&models.User{}, id)
	return result
}

func UpdateUser(id int, data models.User) (*models.User, *gorm.DB, bool) {
	user, result, ok := GetUserById(id)
	if !ok {
		return nil, result, false
	}
	// Update user with new data from request
	godash.Map(data, &user, true)
	result = database.DB.Updates(&user)
	if result.Error != nil {
		return nil, result, true
	}
	return &user, result, true
}

func GetUserById(id int) (models.User, *gorm.DB, bool) {
	user := models.User{}
	result := database.DB.First(&user, id)
	if user.ID != 0 {
		return user, result, true
	}
	return models.User{}, result, false
}

func GetAllUsers() ([]models.User, *gorm.DB) {
	users := []models.User{}
	result := database.DB.Find(&users)
	return users, result
}
