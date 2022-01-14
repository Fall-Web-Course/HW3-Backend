package notes

import (
	"github.com/gin-gonic/gin"
)


func NewNote(c *gin.Context) {
	var new_note NewNoteInput
	c.BindJSON(&new_note)

	user := GetUserByid(new_note.AuthorId)
	InsertToDb(Note{Text: new_note.Text, User: user, UserRefer: new_note.AuthorId})

	c.JSON(200, gin.H{
		"message": "new_note",
	})
}

func GetNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get_note",
	})
}

func UpdateNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update_note",
	})
}

func DeleteNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete_note",
	})
}