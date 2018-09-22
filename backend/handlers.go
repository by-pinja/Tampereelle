package main

import (
	"Tampereelle/backend/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Test struct {
	Name string
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	game := database.CreateGame()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(game)
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameId, _ := strconv.ParseUint(params["id"], 10, 64)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	game := database.GetGame(gameId)
	json.NewEncoder(w).Encode(game)
}

type PlayerData struct {
	Name string `json:"name"`
}

type GameState struct {
	State string `json:"state"`
}

type AnswerDto struct {
	PlayerID uint `json:"playerID"`
	Angle float64 `json:"angle"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func JoinGame(w http.ResponseWriter, r *http.Request) {
	var playerData PlayerData 
		_ = json.NewDecoder(r.Body).Decode(&playerData);
	player := database.CreatePlayer(playerData.Name)

	params := mux.Vars(r)
	gameID, _ := strconv.ParseUint(params["id"], 10, 64)
	database.AddPlayerToGame(player.ID, uint(gameID))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetGameState(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameID, _ := strconv.ParseUint(params["id"], 10, 64)
	game := database.GetGame(gameID)
	gameState := GameState{State: game.State}
	json.NewEncoder(w).Encode(gameState)
}

func UpdateGameState(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameID, _ := strconv.ParseUint(params["id"], 10, 64)
	var gameState GameState 
		_ = json.NewDecoder(r.Body).Decode(&gameState);
	database.UpdateGameState(uint(gameID), gameState.State)
}

func AnswerQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	questionID, _ := strconv.ParseUint(params["questionID"], 10, 64)
	var answer AnswerDto 
		_ = json.NewDecoder(r.Body).Decode(&answer);
	database.CreateAnswer(uint(questionID), answer.PlayerID, answer.Latitude, answer.Longitude, answer.Angle)

}

func GetNextQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameID, _ := strconv.ParseUint(params["id"], 10, 64)
	question := database.NextQuestion(uint(gameID))
	json.NewEncoder(w).Encode(question)
}

func GetQuestionResult(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	questionID, _ := strconv.ParseUint(params["questionID"], 10, 64)
	scores := database.GetPlayerScores(uint(questionID))
	json.NewEncoder(w).Encode(scores)
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
