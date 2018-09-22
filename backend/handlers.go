package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"Tampereelle/backend/database"
)

type Test struct {
	Name string
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	game := database.CreateGame();
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(game)
}

func GetGame(w http.ResponseWriter, r *http.Request) {

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

func TestApi(w http.ResponseWriter, r *http.Request) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, _ := netClient.Get("https://nimiq.mopsus.com/api/quick-stats")
	data, _ := ioutil.ReadAll(response.Body)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(data))
}
