package database

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"
	"net/http"
)

type SomeTable struct {
	gorm.Model
	value 	string
}


func Init(r *http.Request) {
	var db = getConnection(r)
	defer db.Close()

	db.CreateTable(&SomeTable{})
}

func getConnection(r *http.Request) *gorm.DB {
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
	ctx := appengine.NewContext(r)

	db, err := gorm.Open("postgres", params)
	if err != nil {
		log.Errorf(ctx, "Error opening db connection", err)
	}
	return db
}