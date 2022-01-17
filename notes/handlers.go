package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	cache "github.com/Fall-Web-Course/HW3/cache"
	users "github.com/Fall-Web-Course/HW3/users"
	db "github.com/Fall-Web-Course/HW3/db"
)


func NewNote(c *gin.Context) {
	var new_note NewNoteInput
	c.BindJSON(&new_note)

	user := users.GetUserByid(new_note.AuthorId)
	err := InsertToDb(Note{Text: new_note.Text, User: user, UserUsername: new_note.AuthorId})
	if (err.Error != nil) {
		c.JSON(http.StatusBadRequest, gin.H {
			"Message": "Something went wrong",
			"ErrorMesssage": err.Error,
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
		c.JSON(http.StatusOK, gin.H{
			"Message": "Something went wrong",
		})
		return
	}
	var text string
	if text = value.GetValue(); text == "-1" {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Miss cache",
		})
		text = GetNoteById(note_id).Text
		go cache.SetKey(note_id, text)
	}

	c.JSON(http.StatusOK, gin.H{
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
	note_id := c.Param("note_id")

	out := db.GetDb().Delete(&Note{}, note_id)
	if out.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something bad happend",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete_note",
		})
	}
}