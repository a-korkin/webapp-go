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

func getId(uri string) (int, error) {
	if id := utils.GetResourceId(uri); id != "" {
		resourceId, err := strconv.Atoi(id)
		if err != nil {
			return 0, fmt.Errorf("failed convert id: '%s' to int", id)
		}
		return resourceId, nil
	}
	return 0, nil
}

func Persons(w http.ResponseWriter, r *http.Request, appState *data.AppState) {
	switch r.Method {
	case "GET":
		id, err := getId(r.RequestURI)
		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest)
			return
		}
		if id != 0 {
			getPerson(w, id, appState)
		} else {
			getPersons(w, appState)
		}
	case "POST":
		addPerson(w, r, appState)
	case "PUT":
		id, err := getId(r.RequestURI)
		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest)
			return
		}
		if id != 0 {
			updatePerson(w, r, id, appState)
		} else {
			http.Error(
				w,
				"resource id must be set",
				http.StatusBadRequest)
			return
		}
	case "DELETE":
		id, err := getId(r.RequestURI)
		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest)
			return
		}
		if id != 0 {
			deletePerson(w, id, appState)
		} else {
			http.Error(
				w,
				"resource id must be set",
				http.StatusBadRequest)
			return
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
	pers, err := data.GetPerson(id, appState)
	if err != nil {
		log.Fatalf("failed to get person: %s", err)
	}
	if pers == nil {
		http.Error(w, "person not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
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

func updatePerson(
	w http.ResponseWriter, r *http.Request, id int, appState *data.AppState) {
	pers, err := data.GetPerson(id, appState)
	if err != nil {
		log.Fatalf("failed to get person: %s", err)
	}
	if pers == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	newData := data.Person{}
	if err := json.NewDecoder(r.Body).Decode(&newData); err != nil {
		log.Fatalf("failed serialize person: %s", err)
	}
	w.Header().Set("Content-Type", "applcation/json")
	w.WriteHeader(http.StatusOK)
	newData.UpdatePerson(id, appState)
	if err := json.NewEncoder(w).Encode(&newData); err != nil {
		log.Fatalf("failed serialize person: %s", err)
	}
}

func deletePerson(
	w http.ResponseWriter, id int, appState *data.AppState) {
	data.DeletePerson(id, appState)
	w.WriteHeader(http.StatusNoContent)
}
