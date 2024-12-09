package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/a-korkin/webapp/data"
	"github.com/a-korkin/webapp/utils"
)

func Persons(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if id := utils.GetResourceId(r.RequestURI); id != "" {
			log.Printf("params: %v", utils.GetQueryParams(r.URL.RawQuery))
			resourceId, err := strconv.Atoi(id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("failed convert id: '%s' to int", id)))
				return
			}
			getPerson(w, resourceId)
		} else {
			getPersons(w)
		}
	case "POST":
		addPerson(w, r)
	}
}

func getPersons(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data.GetPersons()); err != nil {
		log.Fatalf("failed serialize persons: %s", err)
	}
}

func getPerson(w http.ResponseWriter, id int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data.GetPerson(id)); err != nil {
		log.Fatalf("failed serialize person: %s", err)
	}
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	person := data.Person{}
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		log.Fatalf("failed deserialize person: %s", err)
	}
	person.AddPerson()
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(person); err != nil {
		log.Fatalf("failed serialize person: %s", err)
	}
}
