package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Student struct {
	Rollno  int
	Name    string
	Mark1   int
	Mark2   int
	Mark3   int
	Mark4   int
	Mark5   int
	Total   int
	Average float32
	Grade   string
	FailCount int
	FailedSubjects string

}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
	var Student []student
	params := mux.Vars(r)["age"]
	age, _ := strconv.Atoi(params)
	filter := bson.M{"age": age}
	Collection := Db().Database("testing").Collection("users")
	cur, err := Collection.Find(context.TODO(), filter)
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {

		var stus student
		err := cur.Decode(&stus)
		if err != nil {
			log.Fatal(err)
		}
		stu = append(stu, stus)
	}
	json.NewEncoder(w).Encode(stu)
	fmt.Println("error", err)
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	stu := Student{}
	json.NewDecoder(r.Body).Decode(&stu)
	stu.Total = stu.Mark1 + stu.Mark2 + stu.Mark3 + stu.Mark4 + stu.Mark5
	stu.Average = float32(stu.Total) / 5
	if stu.Average > 90 {
		stu.Grade = "O Grade"
	} else if stu.Average > 80 {
		stu.Grade = "A Grade"
	} else if stu.Average > 70 {
		stu.Grade = "B Grade"
	} else if stu.Average > 60 {
		stu.Grade = "C Grade"
	} else if stu.Average >= 50 {
		stu.Grade = "D Grade"
	} else {
		stu.Grade = "E Grade"
	}

	json.NewEncoder(w).Encode(stu)

}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
