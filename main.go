package main

import (
	"go_blog/database"
	"go_blog/route"
)

func main() {
	defer database.DB.Close()
	route.InitRouter()
}
