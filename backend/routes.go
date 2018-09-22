package main

import (
	"net/http"
)

type Route struct {
	Pattern string
	HandlerFunc http.HandlerFunc
	Method string
}

type Routes []Route

var routes = Routes{
	Route{
		"/api/games",
		CreateGame,
		"POST",
	},
	Route{
		"/api/games/{id}",
		GetGame,
		"GET",
	},
	Route{
		"/api/games/{id}",
		JoinGame,
		"POST",
	},
	Route{
		"/api/games/{id}/state",
		GetGameState,
		"GET",
	},
	Route{
		"/api/games/{id}/state",
		GetGameState,
		"GET",
	},
	Route{
		"/api/games/{id}/state",
		UpdateGameState,
		"PUT",
	},
	Route{
		"/api/games/{id}/questions/next",
		GetNextQuestion,
		"GET",
	},
	Route{
		"/api/games/{id}/questions/{questionID}/answers",
		AnswerQuestion,
		"POST",
	},
	Route{
		"/api/games/{id}/questions/{questionID}",
		GetQuestion,
		"GET",
	},
}