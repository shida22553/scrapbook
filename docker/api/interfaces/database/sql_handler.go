package database

import (
	"gorm.io/gorm"
)

type SqlHandler interface {
	First(out interface{}, where ...interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Order(arg interface{}) *gorm.DB
	Offset(arg int) *gorm.DB
	Limit(arg int) *gorm.DB
	Delete(value interface{}) *gorm.DB
}
