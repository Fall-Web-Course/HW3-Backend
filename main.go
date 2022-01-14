package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)


func new_note(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "new_note",
	})
}

func get_note(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get_note",
	})
}

func update_note(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update_note",
	})
}

func delete_note(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete_note",
	})
}


func main() {
	PORT := Getenv("PORT", "8080");	HOST_LAN_IP := Getenv("LAN_HOST_IP", "127.0.0.1");
	LISTEN_ADDRESS := fmt.Sprintf("%s:%s", HOST_LAN_IP, PORT);

	r := gin.Default();

	r.GET("/ping", func(c *gin.Context) { // For the sake of testing
		c.JSON(200, gin.H{
			"message": "pong",
		})
	});
	r.POST("/notes/new", new_note);
	r.GET("/notes", get_note); // /note_id
	r.PUT("/notes", update_note); // /note_id
	r.DELETE("/notes", delete_note); // /note_id

	r.Run(LISTEN_ADDRESS); // Listens on 0.0.0.0:8080
}