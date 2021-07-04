package domain

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string
	UserID uint
	User   User
}
