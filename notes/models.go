package notes

import (
	"gorm.io/gorm"
)

type Notes struct {
	gorm.Model
	AuthorRefer 	int
	Author		string	`gorm:"foreignKey:UserRefer"`
	Text	string
}
