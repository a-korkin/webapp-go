package main

import (
	"github.com/a-korkin/webapp/handlers/persons"
	"github.com/a-korkin/webapp/utils"
	"log"
	"net/http"
)

func Router(w http.ResponseWriter, r *http.Request) {
	log.Printf("params: %v", utils.GetQueryParams(r.URL.RawQuery))
	switch utils.GetResourcePath(r.RequestURI) {
	case "person":
		persons.GetAll(w, r)
	}
}
