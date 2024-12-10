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

func Persons(w http.ResponseWriter, r *http.Request, appState *data.AppState) {
	switch r.Method {
	case "GET":
		if id := utils.GetResourceId(r.RequestURI); id != "" {
			log.Printf("params: %v", utils.GetQueryParams(r.URL.RawQuery))
			resourceId, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w,
					fmt.Sprintf("failed convert id: '%s' to int", id),
					http.StatusBadRequest)
				return
			}
			getPerson(w, resourceId, appState)
		} else {
			getPersons(w, appState)
		}
	case "POST":
		addPerson(w, r, appState)
	case "DELETE":
		if id := utils.GetResourceId(r.RequestURI); id != "" {
			resourceId, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w,
					fmt.Sprintf("failed convert id: %s to int", id),
					http.StatusBadRequest)
				return
			}
			deletePerson(w, resourceId, appState)
		}
	}
}

func getPersons(w http.ResponseWriter, appState *data.AppState) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data.GetPersons(appState)); err != nil {
		log.Fatalf("failed serialize persons: %s", err)
	}
}

func getPerson(w http.ResponseWriter, id int, appState *data.AppState) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	pers, err := data.GetPerson(id, appState)
	if err != nil {
		log.Fatalf("failed to get person: %s", err)
	}
	if err := json.NewEncoder(w).Encode(pers); err != nil {
		log.Fatalf("failed serialize person: %s", err)
	}
}

func addPerson(w http.ResponseWriter, r *http.Request, appState *data.AppState) {
	person := data.Person{}
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		log.Fatalf("failed deserialize person: %s", err)
	}
	person.AddPerson(appState)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(person); err != nil {
		log.Fatalf("failed serialize person: %s", err)
	}
}

func deletePerson(
	w http.ResponseWriter, id int, appState *data.AppState) {
	data.DeletePerson(id, appState)
	w.WriteHeader(http.StatusNoContent)
}
