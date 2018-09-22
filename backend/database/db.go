package database

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
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
	game := Game{State: "PENDING"}
	db.Create(&game)
	return game
}

func GetGame(gameId uint64) Game {
	db := getConnection()
	var game Game
	db.First(&game, gameId)
	return game
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