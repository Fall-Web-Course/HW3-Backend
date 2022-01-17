package main

import (
	init_db "github.com/Fall-Web-Course/HW3/db/init"
	router "github.com/Fall-Web-Course/HW3/router"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	init_db.InitDb()
	router.InitRouter();
}