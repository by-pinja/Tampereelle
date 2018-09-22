package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

type Test struct {
	Name string
}

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/test", TodoIndex)

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
