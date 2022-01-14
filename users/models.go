package users

import (
	"github.com/Fall-Web-Course/HW3/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID			int 	`gorm:"primaryKey"`
	Username	string  `json:"username"`
	Password	string  `json:"password"`
	IsAdmin		bool    `json:"id_admin"`
}

func InsertToDb(user User) {
	db.GetDb().Create(&user)
}