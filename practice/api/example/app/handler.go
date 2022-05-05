package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type student struct {
	Rollno int    `json:"Roll No" xml:"Rollno"`
	Name   string `json:"Name" xml:"name"`
	Mark   int    `json:"Mark" xml:"mark"`
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "WELCOME.........")
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	stu1 := []student{
		{Rollno: 123, Name: "Divya", Mark: 100},
		{Rollno: 124, Name: "Viji", Mark: 100},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(stu1)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stu1)
	}
}

func GetValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["Rollno"])
}
