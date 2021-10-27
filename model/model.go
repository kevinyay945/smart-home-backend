package model

import "gorm.io/gorm"

var db *gorm.DB

func Init(_db *gorm.DB) {
	db = _db
}
