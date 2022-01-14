package users

import "github.com/Fall-Web-Course/HW3/db"

type User struct {
	// gorm.Model
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	IsAdmin		bool	`json:"id_admin"`
}

func InsertToDb(user User) {
	db.GetDb().Create(&user)
}