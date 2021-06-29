package domain

import (
	"gorm.io/gorm"
)

type Cutting struct {
	gorm.Model
	Note   string `gorm:"text"`
	UserID uint
	User   User
}
