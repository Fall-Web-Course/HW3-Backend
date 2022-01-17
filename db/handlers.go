package db

import "gorm.io/gorm"

var db *gorm.DB

func GetDb() *gorm.DB {
	return db
}

func SetDB(database *gorm.DB) {
	db = database
}