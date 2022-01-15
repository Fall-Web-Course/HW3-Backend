package users

import (
	"github.com/Fall-Web-Course/HW3/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username	string  `json:"username" gorm:"unique"`
	Password	string  `json:"password"`
	IsAdmin		bool    `json:"id_admin"`
}

func InsertToDb(user User) *gorm.DB {
	err := db.GetDb().Create(&user)
	return err
}