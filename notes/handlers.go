package notes

import (
	"net/http"

	cache "github.com/Fall-Web-Course/HW3/cache"
	"github.com/gin-gonic/gin"
)


func NewNote(c *gin.Context) {
	var new_note NewNoteInput
	c.BindJSON(&new_note)

	user := GetUserByid(new_note.AuthorId)
	err := InsertToDb(Note{Text: new_note.Text, User: user, UserUsername: new_note.AuthorId})
	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H {
			"Message": "Something went wrong",
		})
	}
	c.JSON(http.StatusCreated, gin.H {
		"Message": "Note created",
	})
}

func GetNote(c *gin.Context) {
	note_id := c.Param("note_id")
	value, err := cache.GetKey(note_id)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"Text": value,
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