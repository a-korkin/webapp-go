package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello from handler\n"))
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", handler)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("couldn't start server")
	}
}
