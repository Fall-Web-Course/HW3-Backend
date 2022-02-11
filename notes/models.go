package notes

import (
	"strconv"

	"github.com/Fall-Web-Course/HW3/db"
	"github.com/Fall-Web-Course/HW3/users"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserUsername string
	User         users.User `gorm:"references:Username"`
	Text         string
}

type NewNoteInput struct {
	AuthorId string `json:"user_id"`
	Text     string `json:"text"`
}

func InsertToDb(note Note) (*gorm.DB, uint) {
	res := db.GetDb().Create(&note)
	return res, note.ID
}

func GetNoteById(note_id string) Note {
	var note Note
	db.GetDb().Find(&note, note_id)
	return note
}

func UpdateNoteById(note_id string, text string) {
	note := GetNoteById(note_id)
	note.Text = text
	db.GetDb().Save(&note)
}

func GetUserNotesById(user_id string) (res string) {
	var notes []Note
	user := users.GetUserByid(user_id)
	db.GetDb().Where("user_username = ?", user.Username).Find(&notes)

	len_notes := len(notes)
	for i := 0; i < len_notes; i++ {
		res += strconv.FormatUint(uint64(notes[i].ID), 10)
		if i != len_notes-1 {
			res += " "
		}
	}
	return res
}
