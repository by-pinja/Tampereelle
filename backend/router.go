package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() {
	r := mux.NewRouter()
	for _, route := range routes {
		r.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}
	http.Handle("/", r)
}