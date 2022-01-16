package notes

import (
	"github.com/Fall-Web-Course/HW3/db"
	"github.com/Fall-Web-Course/HW3/users"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserUsername	string
	User			users.User	`gorm:"references:Username"`
	Text			string
}

type NewNoteInput struct {
	AuthorId	string	`json:"user_id"`
	Text		string	`json:"text"`
}

func InsertToDb(note Note) *gorm.DB {
	err := db.GetDb().Create(&note)
	return err
}

func GetNoteByid(note_id string) Note {
	var note Note
	db.GetDb().Find(&note, note_id)
	return note
}