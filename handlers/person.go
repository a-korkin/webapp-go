package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/a-korkin/webapp/data"
	"github.com/a-korkin/webapp/utils"
	"log"
	"net/http"
	"strconv"
)

func Persons(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Printf("params: %v", utils.GetQueryParams(r.URL.RawQuery))
		if id := utils.GetResourceId(r.RequestURI); id != "" {
			resourceId, err := strconv.Atoi(id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("couldn't convert id: '%s' to int", id)))
				return
			}
			getPerson(w, resourceId)
		} else {
			getPersons(w)
		}
	}
}

func getPersons(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data.GetPersons()); err != nil {
		log.Fatalf("couldn't serialize persons: %s", err)
	}
}

func getPerson(w http.ResponseWriter, id int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data.GetPerson(id)); err != nil {
		log.Fatalf("couldn't serialize person: %s", err)
	}

}
