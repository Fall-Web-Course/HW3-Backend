package notes

import (
	"github.com/Fall-Web-Course/HW3/db"
	"github.com/Fall-Web-Course/HW3/users"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserRefer	int
	User		users.User	`gorm:"foreignKey:UserRefer"`
	Text		string
}

type NewNoteInput struct {
	AuthorId	int		`json:"user_id"`
	Text		string	`json:"text"`
}

func InsertToDb(note Note) {
	db.GetDb().Create(&note)
}

func GetUserByid(user_id int) users.User {
	var user users.User
	db.GetDb().Find(&user, user_id)
	return user
}