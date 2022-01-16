package router

import (
	users "github.com/Fall-Web-Course/HW3/users"
	notes "github.com/Fall-Web-Course/HW3/notes"
	utils "github.com/Fall-Web-Course/HW3/utils"

	"github.com/gin-gonic/gin"
	"fmt"
)

func InitRouter() {
	PORT := utils.Getenv("PORT", "8080");	HOST_LAN_IP := utils.Getenv("LAN_HOST_IP", "127.0.0.1");
	LISTEN_ADDRESS := fmt.Sprintf("%s:%s", HOST_LAN_IP, PORT);

	router := gin.Default();

	router.POST("/notes/new", notes.NewNote);
	router.GET("/notes/:note_id", notes.GetNote);
	router.PUT("/notes/:note_id", notes.UpdateNote); // /note_id
	router.DELETE("/notes/:note_id", notes.DeleteNote); // /note_id
	
	router.POST("/users/register", users.Register);	
	router.POST("/users/login", users.Login);	

	router.Run(LISTEN_ADDRESS); // Listens on 0.0.0.0:8080 by default
}
