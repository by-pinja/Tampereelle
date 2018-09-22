package database

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"math/rand"
)

type Place struct {
	gorm.Model
	Name string
	Latitude float64
	Longitude float64
}

type Game struct {
	gorm.Model
	State string
	Players []Player
	Questions []Question
}

type Player struct {
	gorm.Model
	Name string
}

type Question struct {
	gorm.Model
	Place Place
	State string
	Game Game
	Answers []Answer
}

type Answer struct {
	gorm.Model
	Player Player
	Question Question
	Angle float64
	PlayerLatitude float64
	PlayerLongitude float64
}

func CreateGame() Game {
	db := getConnection()
	defer db.Close()

	game := Game{State: "PENDING"}
	db.Create(&game)
	return game
}

func GetGame(gameId uint64) Game {
	db := getConnection()
	defer db.Close()

	var game Game
	db.First(&game, gameId)
	return game
}

func UpdateGameState(gameId uint, state string) {
	db := getConnection()
	defer db.Close()

	var game Game
	db.First(&game, gameId)
	game.State = state
	db.Save(game)
}

func CreatePlayer(playerName string) Player {
	db := getConnection()
	defer db.Close()

	player := Player{Name: playerName}
	db.Save(player)
	return player
}

func AddPlayerToGame(playerId uint, gameId uint) {
	db := getConnection()
	defer db.Close()

	var game Game
	db.First(&game, gameId)

	var player Player
	db.First(&player, playerId)
	players := append(game.Players, player)
	game.Players = players
	db.Save(game)
}

func GetQuestion(questionId uint) Question {
	db := getConnection()
	defer db.Close()

	var question Question
	db.First(&question, questionId)
	return question
}

func NextQuestion(gameId uint) Question {
	db := getConnection()
	defer db.Close()

	var game Game
	db.First(&game, gameId)

	var question Question
	db.Where(Question{Game: game, State: "OPEN"}).First(&question)
	if db.Where(Question{Game: game, State: "OPEN"}).First(&question).RecordNotFound() {
		var places []Place
		db.Find(&places)
		place := places[rand.Intn(len(places))]


		question = Question{Place: place, Game: game, State: "OPEN"}
		db.Save(question)
	}

	return question
}

func GetPlayer(playerId uint) Player {
	db := getConnection()
	defer db.Close()

	var player Player
	db.First(&player, playerId)
	return player
}

func CreateAnswer(questionId uint, playerId uint, playerLatitude float64, playerLongitude float64, angle float64) {
	db := getConnection()
	defer db.Close()

	var question Question
	db.First(&question, questionId)

	var player Player
	db.First(&player, playerId)

	answer := Answer{Question: question, Player: player, Angle: angle, PlayerLatitude: playerLatitude, PlayerLongitude: playerLongitude}
	db.Save(answer)

	if len(question.Answers) == len(question.Game.Players) {
		question.State = "CLOSED"
		db.Save(question)
	}
}



func Init() {
	var db = getConnection()
	defer db.Close()
	db.AutoMigrate(&Game{}, &Place{}, &Player{}, &Question{}, &Answer{})
}

func getConnection() *gorm.DB {
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_port := os.Getenv("DB_PORT")
	db_password := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")

	var params string
	if db_port != "" {
		params = fmt.Sprintf(
			"host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
			db_host,
			db_port,
			db_user,
			db_name,
			db_password)
	} else {
		params = fmt.Sprintf("host=%v user=%v dbname=%v password=%v sslmode=disable",
			db_host,
			db_user,
			db_name,
			db_password)
	}

	db, err := gorm.Open("postgres", params)
	if err != nil {
	}
	return db
}