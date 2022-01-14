package notes

import (
	"github.com/Fall-Web-Course/HW3/db"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	AuthorRefer 	int
	Author			string	`gorm:"foreignKey:AuthorRefer"`
	Text			string
}

func InsertToDb(note Note) {
	db.GetDb().Create(&note)
}