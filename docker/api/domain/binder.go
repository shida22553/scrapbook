package domain

import (
	"gorm.io/gorm"
)

type Binder struct {
	gorm.Model
	Name   string
	UserID uint
	User   User
}
