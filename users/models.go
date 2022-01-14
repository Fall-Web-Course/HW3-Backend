package user

import (
	"gorm.io/gorm"
)

type Notes struct {
	gorm.Model
	AuthorRefer 	int
	Author		string	`gorm:"foreignKey:UserRefer"`
	Text	string
}

type User struct {
	// gorm.Model
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	IsAdmin		bool	`json:"id_admin"`
}