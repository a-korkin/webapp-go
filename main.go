package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type AppState struct {
	Db *sql.DB
}

var appState = AppState{}

func init() {
	var err error
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		GetEnv("DB_USR"), GetEnv("DB_PWD"), GetEnv("DB_NAME"))
	appState.Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %s", err)
	}
	defer func() {
		if err = appState.Db.Close(); err != nil {
			log.Fatalf("failed to close connection to postgres: %s", err)
		}
	}()
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	router := Router{
		State: appState,
	}
	http.Handle("/", router)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed start server: %s", err)
	}
}
