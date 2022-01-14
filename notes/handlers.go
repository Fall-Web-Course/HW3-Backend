package notes

import (
	"github.com/gin-gonic/gin"
)


func NewNote(c *gin.Context) {
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