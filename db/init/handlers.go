package init

import (
	"fmt"

	db "github.com/Fall-Web-Course/HW3/db"
	notes "github.com/Fall-Web-Course/HW3/notes"
	users "github.com/Fall-Web-Course/HW3/users"
	utils "github.com/Fall-Web-Course/HW3/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDB() (db *gorm.DB, err error) {
	DB_PORT := utils.Getenv("DB_PORT", "8080")
	DB_ADDRESS := utils.Getenv("DB_ADDRESS", "127.0.0.1")
	DB_USER := utils.Getenv("DB_USER", "db_user")
	DB_NAME := utils.Getenv("DB_NAME", "db_name")
	DB_PASSWORD := utils.Getenv("DB_PASSWORD", "db_password")

	DB_DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", DB_ADDRESS, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	
	db, err = gorm.Open(postgres.Open(DB_DSN), &gorm.Config{})
	return db, err
}

func InitDb() {
	database, err := connectToDB()
	if err != nil {
		panic(err)
	}

	db.SetDB(database)
	db.GetDb().AutoMigrate(&users.User{}, &notes.Note{})
}