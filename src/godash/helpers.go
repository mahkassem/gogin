package godash

import (
	"fmt"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetMethodByName(parentStruct any, name string) (gin.HandlerFunc, error) {
	parent := reflect.ValueOf(parentStruct)
	method := parent.MethodByName(name)
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

func If[T any](condition bool, a, b T) T {
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

func Map[A any](source A, destination *A, skipZeroValues ...bool) error {
	for i := 0; i < reflect.ValueOf(*destination).NumField(); i++ {
		name := reflect.ValueOf(*destination).Type().Field(i).Name
		value := reflect.ValueOf(source).Field(i).Interface()
		structValue := reflect.ValueOf(destination).Elem()
		fieldValue := structValue.FieldByName(name)
		if !fieldValue.IsValid() {
			return fmt.Errorf("no such field: %s in obj", name)
		}
		if !fieldValue.CanSet() {
			return fmt.Errorf("cannot set field %s", name)
		}
		val := reflect.ValueOf(value)
		if len(skipZeroValues) > 0 && val.IsZero() {
			continue
		}
		if fieldValue.Type() != val.Type() {
			return fmt.Errorf("provided value type didn't match obj field type")
		}
		fieldValue.Set(val)
	}
	return nil
}
