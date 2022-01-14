package notes

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	AuthorRefer 	int
	Author		string	`gorm:"foreignKey:UserRefer"`
	Text	string
}
