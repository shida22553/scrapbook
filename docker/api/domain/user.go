package domain

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uid  string
	Name string
}
