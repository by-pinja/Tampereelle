package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Test struct {
	Name string
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome1!")
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