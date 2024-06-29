package utilities

import (
	"fmt"
	"main/src/database/models"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetMethodByName(name string, controller any) (gin.HandlerFunc, error) {
	ctrl := reflect.ValueOf(controller)
	method := ctrl.MethodByName(name)
	if !method.IsValid() {
		return nil, fmt.Errorf("no function with name %s found", name)
	}
	handlerFunc, ok := method.Interface().(func(*gin.Context))
	if !ok {
		return nil, fmt.Errorf("function %s has wrong signature", name)
	}
	return gin.HandlerFunc(handlerFunc), nil
}

func DetectMigration() bool {
	if len(os.Args) < 2 {
		return false
	}
	command := os.Args[1]
	return command == "migrate" || command == "-m" || command == "--migrate"
}

func Assign(_1 interface{}, _2 []interface{}) {
	// Reflect value of the existing user
	v := reflect.ValueOf(_1).Elem()

	// Iterate over the request body and update fields dynamically
	for key, value := range _2 {
		fmt.Println(key, value)
		field := v.FieldByName(string(key))
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(value))
		}
	}
}

func HashPassword(password string) string {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func ComparePassword(hashedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func AssignDataToUser(data models.User, user *models.User) {
	if data.Email != "" {
		user.Email = data.Email
	}

	if data.Username != "" {
		user.Username = data.Username
	}
}

func Ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

func StringInSlice(str string, slice []string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
