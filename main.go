package main

import (
	init_db "github.com/Fall-Web-Course/HW3/db/init"
	init_cache "github.com/Fall-Web-Course/HW3/cache/init"
	router "github.com/Fall-Web-Course/HW3/router"
	
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load(".env");

	init_db.InitDb();
	init_cache.InitCache();
	router.InitRouter();
}