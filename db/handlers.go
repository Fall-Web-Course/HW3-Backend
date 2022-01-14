package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	users "github.com/Fall-Web-Course/HW3/users"
	notes "github.com/Fall-Web-Course/HW3/notes"
)

var DB_DSN string = "host=localhost user=hw3 password=webhw3password dbname=hw3 port=5433 sslmode=disable TimeZone=Asia/Tehran"
var db, err = gorm.Open(postgres.Open(DB_DSN), &gorm.Config{})

func init_db() {
	db.AutoMigrate(&users.User{}, &notes.Note{})
}