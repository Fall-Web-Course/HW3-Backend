package init

import (
	db "github.com/Fall-Web-Course/HW3/db"
	users "github.com/Fall-Web-Course/HW3/users"
	notes "github.com/Fall-Web-Course/HW3/notes"
)

func InitDb() {
	db.GetDb().AutoMigrate(&users.User{}, &notes.Note{});
}