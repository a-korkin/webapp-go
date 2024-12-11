package router

import (
	"github.com/a-korkin/webapp/data"
	"github.com/a-korkin/webapp/handlers"
	"github.com/a-korkin/webapp/utils"
	"net/http"
)

type Router struct {
	AppState *data.AppState
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch utils.GetResourcePath(r.RequestURI) {
	case "person":
		handlers.Persons(w, r, router.AppState)
	}
}
