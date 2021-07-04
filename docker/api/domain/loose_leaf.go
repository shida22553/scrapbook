package domain

import (
	"gorm.io/gorm"
)

type LooseLeaf struct {
	gorm.Model
	Content string `gorm:"text"`
	UserID  uint
	// BookID uint
	User User
	// Book   Book
}
