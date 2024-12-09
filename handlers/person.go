package handlers

import (
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("list of persons\n"))
}
