package database

import (
	"gorm.io/gorm"
)

type SqlHandler interface {
	First(out interface{}, where ...interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Order(arg interface{}) *gorm.DB
}
