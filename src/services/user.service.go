package services

import (
	"main/src/config"
	"main/src/database/models"
	"reflect"
)

func GetUserById(id int) (models.User, bool) {
	user := models.User{}
	config.DB.First(&user, id)
	if user.ID != 0 {
		return user, true
	}
	return models.User{}, false
}

func GetAllUsers() ([]models.User, bool) {
	users := []models.User{}
	config.DB.Find(&users)
	return users, true
}

func AssignUserToAnotherUser(user models.User, anotherUser models.User) {
	v1 := reflect.ValueOf(anotherUser)
	v2 := reflect.ValueOf(anotherUser)
	for i := 0; i < v1.NumField(); i++ {
		v2.Field(i).Set(v1.Field(i))
	}
}
