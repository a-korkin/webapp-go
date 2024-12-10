package main

import (
	"github.com/a-korkin/webapp/handlers"
	"github.com/a-korkin/webapp/utils"
	"net/http"
)

type Router struct {
	State AppState
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch utils.GetResourcePath(r.RequestURI) {
	case "person":
		handlers.Persons(w, r)
	}
}
