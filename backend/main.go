package main

import (
	"google.golang.org/appengine"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
	appengine.Main()
}