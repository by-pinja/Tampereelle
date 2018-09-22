package main

import (
	"net/http"
)

func NewRouter() {
	for _, route := range routes {
		http.HandleFunc(route.Pattern, route.HandlerFunc)
	}
}