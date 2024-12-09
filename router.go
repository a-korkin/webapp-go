package main

import (
	"github.com/a-korkin/webapp/handlers"
	"log"
	"net/http"
	"strings"
)

func getResourcePath(uri string) string {
	return strings.Split(strings.ToLower(uri), "/")[1]
}

func Router(w http.ResponseWriter, r *http.Request) {
	uri := strings.ToLower(r.RequestURI)
	log.Printf("request uri: %v", uri)
	log.Printf(getResourcePath(uri))
	switch getResourcePath(uri) {
	case "person":
		handlers.GetAll(w, r)
	}
}
