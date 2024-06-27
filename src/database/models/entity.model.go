package models

import "gorm.io/gorm"

type Entity struct {
	gorm.Model
	ID uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
}
