package main

import (
	init_db "github.com/Fall-Web-Course/HW3/db/init"
	router "github.com/Fall-Web-Course/HW3/router"
)

func main() {
	init_db.InitDb()
	router.InitRouter();
}