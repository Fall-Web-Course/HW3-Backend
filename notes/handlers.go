package notes

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	jwt "github.com/Fall-Web-Course/HW3/authorization"
	cache "github.com/Fall-Web-Course/HW3/cache"
	db "github.com/Fall-Web-Course/HW3/db"
	users "github.com/Fall-Web-Course/HW3/users"
)

func NewNote(c *gin.Context) {
	var new_note NewNoteInput
	c.BindJSON(&new_note)

	user := users.GetUserByid(new_note.AuthorId)
	res, id := InsertToDb(Note{Text: new_note.Text, User: user, UserUsername: new_note.AuthorId})
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message":       "Something went wrong",
			"ErrorMesssage": res.Error,
		})
		return
	}
	updateUserNotes(new_note.AuthorId, strconv.FormatUint(uint64(id), 10))
	c.JSON(http.StatusCreated, gin.H{
		"Message": "Note created",
	})
}

func GetNote(c *gin.Context) {
	ibearToken := c.Request.Header.Get("Authorization")
	str_split := strings.Split(ibearToken, " ")
	var user_id string
	var errs error
	if len(str_split) == 2 {
		user_id, errs = jwt.GetUserID(str_split[1])
		if errs != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"Message": "Authentication token not provided.",
			})
			return
		}
	} else {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"Message": "Authentication token not provided.",
		})
		return
	}

	fmt.Print(user_id)
	note_id := c.Param("note_id")
	value, _ := cache.GetKey(user_id)
	user_notes := value.GetValue()
	if user_notes == "-1" {
		user_notes = GetUserNotesById(user_id)
		go cache.SetKey(fmt.Sprintf("u%s", user_id), user_notes)
	}
	if !strings.Contains(user_notes, note_id) {
		c.JSON(http.StatusForbidden, gin.H{
			"Message": "Access denied",
		})
		return
	}
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

	c.JSON(http.StatusOK, gin.H{"Text": text})
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

func updateUserNotes(user_id string, new_note_id string) {
	notes, err := cache.GetKey(fmt.Sprintf("u%s", user_id))
	if err != nil {
		panic(err)
	}
	cache.SetKey(fmt.Sprintf("u%s", user_id), fmt.Sprintf("%s %s", notes, new_note_id))
}
