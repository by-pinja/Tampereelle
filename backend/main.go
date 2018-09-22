package main

import (
	"google.golang.org/appengine"
	"Tampereelle/backend/database"
)

func main() {
	database.Init()
	NewRouter()
	appengine.Main()
}