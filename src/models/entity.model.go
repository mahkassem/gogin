package models

import "gorm.io/gorm"

type Entity struct {
	gorm.Model
	Id string
}
