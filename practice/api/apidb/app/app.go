package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/get", GetStudent).Methods(http.MethodGet)
	router.HandleFunc("/post", AddStudent).Methods(http.MethodPost)
}
