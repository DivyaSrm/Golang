package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/greet", Greet)

	router.HandleFunc("/get", GetStudent).Methods(http.MethodGet)
	router.HandleFunc("/getstudent/{Rollno:[0-9]+}", GetValue).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
