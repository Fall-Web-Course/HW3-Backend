package main

import (
	users "github.com/Fall-Web-Course/HW3/users"
	notes "github.com/Fall-Web-Course/HW3/notes"

	"github.com/gin-gonic/gin"

	"fmt"
)

func init_router(LISTEN_ADDRESS string) {
	router := gin.Default();

	router.POST("/notes/new", notes.NewNote);
	router.GET("/notes", notes.GetNote); // /note_id
	router.PUT("/notes", notes.UpdateNote); // /note_id
	router.DELETE("/notes", notes.DeleteNote); // /note_id
	
	router.POST("/users/register", users.Register);	
	router.POST("/users/login", users.Login);	

	router.Run(LISTEN_ADDRESS); // Listens on 0.0.0.0:8080 by default
}

func main() {
	PORT := Getenv("PORT", "8080");	HOST_LAN_IP := Getenv("LAN_HOST_IP", "127.0.0.1");
	LISTEN_ADDRESS := fmt.Sprintf("%s:%s", HOST_LAN_IP, PORT);

	init_router(LISTEN_ADDRESS)
}