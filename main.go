package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", Router)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed start server: %s", err)
	}
}
