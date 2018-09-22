package database

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"math/rand"
	"math"
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
	Players []Player `gorm:"foreignkey:GameId"`
	Questions []Question `gorm:"foreignkey:GameId"`
}

type Player struct {
	gorm.Model
	Name string
	GameId uint
}

type Question struct {
	gorm.Model
	Place Place  `gorm:"foreignkey:PlaceId"`
	State string
	Game Game `gorm:"foreignkey:GameId"`
	Answers []Answer
	GameId uint
	PlaceId uint
}

type Answer struct {
	gorm.Model
	Player Player `gorm:"foreignkey:PlayerId"`
	Question Question `gorm:"foreignkey:QuestionId"`
 	Angle float64
	PlayerLatitude float64
	PlayerLongitude float64
	PlayerId uint
	QuestionId uint
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
	db.Save(&game)
}

func CreatePlayer(playerName string) Player {
	db := getConnection()
	defer db.Close()

	player := Player{Name: playerName}
	db.Save(&player)
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
	db.Save(&game)
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
	db.Where(Question{GameId: game.ID, State: "OPEN"}).First(&question)
	if db.Where(Question{GameId: game.ID, State: "OPEN"}).First(&question).RecordNotFound() {
		var places []Place
		db.Find(&places)
		place := places[rand.Intn(len(places))]

		question = Question{Place: place, Game: game, State: "OPEN"}
		db.Save(&question)
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
	db.Save(&answer)

	if len(question.Answers) == len(question.Game.Players) {
		question.State = "CLOSED"
		db.Save(question)
	}
}


type PlayerScore struct {
	Player Player
	Score float64
}

func dist(dLat float64, dLon float64) float64 {
	return math.Sqrt(dLat * dLat + dLon * dLon)
}

func angle(x1 float64, x2 float64, y1 float64, y2 float64) float64 {
	distX := dist(x1, y1)
	distY := dist(x2, y2)
	return math.Acos((x1 * x2 + y1 * y2) / (distX  * distY))

}

func getPlayerScore(question Question, answer Answer) float64 {
	realDLat := question.Place.Latitude - answer.PlayerLatitude
	realDLon := question.Place.Longitude - answer.PlayerLongitude

	realAngle := angle(realDLon, 1.0, realDLat, 0)
	fmt.Println(realAngle)

	ansRotated := answer.Angle + 90
	if ansRotated > 360 {
		ansRotated =  ansRotated - 360
	}
	ansRotated = 360 - ansRotated
	ansRad := ((ansRotated) / 360.0) * 2 * math.Pi
	fmt.Println(answer.Angle)
	fmt.Println(ansRad)

	ansDLat := math.Sin(ansRad)
	ansDLon := math.Cos(ansRad)

	return angle(realDLat, ansDLat, realDLon, ansDLon)
}

func GetPlayerScores(questionId uint) []PlayerScore {
	db := getConnection()
	defer db.Close()

	var question Question

	if db.First(&question, questionId).RecordNotFound() {
		panic("No such question")
	}

	var result []PlayerScore

	if question.State == "OPEN" {
		return result
	}

	var game Game
	db.Model(&question).Related(&game, "GameId").Row()
	var players []Player
	db.Model(&game).Related(&players, "Players")
	for i := 0 ; i < len(players) ; i++ {
		player := players[i]
		var answer Answer
		var playerScore PlayerScore
		if db.Where(Answer{PlayerId: player.ID}).First(&answer).RecordNotFound() {
			score := getPlayerScore(question, answer)
			playerScore = PlayerScore{Player: player, Score: score}
		} else {
			playerScore = PlayerScore{Player: player, Score: 99999}
		}
		result = append(result, playerScore)
	}

	return result
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
