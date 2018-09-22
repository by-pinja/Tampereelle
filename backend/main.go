package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"

	"Tampereelle/backend/database"
)

type Test struct {
	Name string
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/test", TodoIndex)
	http.HandleFunc("/db", DbTest)
	appengine.Main()
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	test := Test{Name: "Test"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(test)
}

func DbTest(w http.ResponseWriter, r *http.Request) {
	database.Init(r)

	test := Test{Name: "DB-init"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(test)
}