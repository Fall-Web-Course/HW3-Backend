package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_DSN string = "host=localhost user=postgres password=webhw3password dbname=hw3 port=5433 sslmode=disable TimeZone=Asia/Tehran"
var db, err = gorm.Open(postgres.Open(DB_DSN), &gorm.Config{})

func GetDb() *gorm.DB{
	return db
}