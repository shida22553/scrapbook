package domain

import (
	"gorm.io/gorm"
)

const Limit = 30

type Binder struct {
	gorm.Model
	Name   string
	UserID uint
	User   User
}
