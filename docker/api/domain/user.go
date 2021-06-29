package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uid  string
	Name string
}
