package main

import (
	db "github.com/Fall-Web-Course/HW3/db"
	router "github.com/Fall-Web-Course/HW3/router"
)

func main() {
	db.InitDb();
	router.InitRouter();
}