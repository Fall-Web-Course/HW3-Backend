package router

import (
	notes "github.com/Fall-Web-Course/HW3/notes"
	middlewares "github.com/Fall-Web-Course/HW3/router/middleware"
	users "github.com/Fall-Web-Course/HW3/users"
	utils "github.com/Fall-Web-Course/HW3/utils"

	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	PORT := utils.Getenv("PORT", "8080")
	HOST_LAN_IP := utils.Getenv("LAN_HOST_IP", "127.0.0.1")
	LISTEN_ADDRESS := fmt.Sprintf("%s:%s", HOST_LAN_IP, PORT)

	router := gin.Default()

	router.Use(middlewares.RateLimit)

	router.POST("/notes/new", notes.NewNote)
	router.GET("/notes/:note_id", notes.GetNote)
	router.PUT("/notes/:note_id", notes.UpdateNote)
	router.DELETE("/notes/:note_id", notes.DeleteNote)

	router.POST("/users/register", users.Register)
	router.POST("/users/login", users.Login)

	router.Run(LISTEN_ADDRESS) // Listens on 0.0.0.0:8080 by default
}
