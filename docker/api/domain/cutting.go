package domain

import (
	"github.com/jinzhu/gorm"
)

type Cutting struct {
	gorm.Model
	Note   string `gorm:"text"`
	UserID uint
	User   User
}

type CuttingPutRequest struct {
	Note string `json:"note"`
}

type CuttingsGetRequest struct {
	UserID   uint
	Page     int
	PageSize int
}
