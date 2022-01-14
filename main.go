package main

import (
	"fmt"

	users "github.com/Fall-Web-Course/HW3/users"
	"github.com/gin-gonic/gin"
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

func init_router(LISTEN_ADDRESS string) {
	router := gin.Default();

	router.POST("/notes/new", new_note);
	router.GET("/notes", get_note); // /note_id
	router.PUT("/notes", update_note); // /note_id
	router.DELETE("/notes", delete_note); // /note_id
	
	router.POST("/users/register", users.Register);	
	router.POST("/users/login", users.Login);	

	router.Run(LISTEN_ADDRESS); // Listens on 0.0.0.0:8080 by default
}

func main() {
	PORT := Getenv("PORT", "8080");	HOST_LAN_IP := Getenv("LAN_HOST_IP", "127.0.0.1");
	LISTEN_ADDRESS := fmt.Sprintf("%s:%s", HOST_LAN_IP, PORT);

	init_router(LISTEN_ADDRESS)
}