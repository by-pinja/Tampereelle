package main

import (
	"net/http"
)

type Route struct {
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"/api/games",
		CreateGame,
	},
	Route{
		"/api/games/{id}",
		GetGame,
	},
	Route{
		"/",
		Index,
	},
	Route{
		"/test",
		TodoIndex,
	},
	Route{
		"/api/test",
		TestApi,
	},
}