package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	cache "github.com/Fall-Web-Course/HW3/cache"
	users "github.com/Fall-Web-Course/HW3/users"
)


func NewNote(c *gin.Context) {
	var new_note NewNoteInput
	c.BindJSON(&new_note)

	user := users.GetUserByid(new_note.AuthorId)
	err := InsertToDb(Note{Text: new_note.Text, User: user, UserUsername: new_note.AuthorId})
	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H {
			"Message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H {
		"Message": "Note created",
	})
}

func GetNote(c *gin.Context) {
	note_id := c.Param("note_id")
	value, err := cache.GetKey(note_id)
	if err != nil {
		c.JSON(200, gin.H{
			"Message": "Something went wrong",
		})
		return
	}
	var text string
	if text = value.GetValue(); text == "-1" {
		c.JSON(200, gin.H{
			"Message": "Miss cache\n",
		})
		text = GetNoteById(note_id).Text
		go cache.SetKey(note_id, text)
	}

	c.JSON(200, gin.H{
		"Text": text,
	})
}

func UpdateNote(c *gin.Context) {
	var note NewNoteInput
	c.BindJSON(&note)

	note_id := c.Param("note_id")
	go cache.SetKey(note_id, note.Text)
	go UpdateNoteById(note_id, note.Text)

	c.JSON(200, gin.H{
		"message": "Note updated",
	})
}

func DeleteNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete_note",
	})
}