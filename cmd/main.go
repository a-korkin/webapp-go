package main

import (
	"database/sql"
	"fmt"
	"github.com/a-korkin/webapp/config"
	"github.com/a-korkin/webapp/data"
	"github.com/a-korkin/webapp/router"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var AppState = data.AppState{}

func main() {
	var err error
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnv("DB_USR"), config.GetEnv("DB_PWD"), config.GetEnv("DB_NAME"))
	AppState.Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %s", err)
	}
	defer func() {
		log.Printf("db connection close")
		if err := AppState.Db.Close(); err != nil {
			log.Fatalf("failed to close connection to postgres: %s", err)
		}
	}()

	server := http.Server{
		Addr: ":8080",
	}
	router := router.Router{
		AppState: &AppState,
	}
	http.Handle("/", router)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed start server: %s", err)
	}
}
