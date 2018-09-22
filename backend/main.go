package main

import (
	"encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type Test struct {
	Name string
}

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/test", TodoIndex)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	test := Test{Name: "Test"}
	json.NewEncoder(w).Encode(test);
}

