package app

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type student struct {
	FullName string `json:"Name" xml:"name"`
	Age      int    `json:"Age" xml:"mark"`
}

func GetStudent(w http.ResponseWriter, r *http.Request) {

}

var client *mongo.Client

func AddStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	var stu student
	json.NewDecoder(r.Body).Decode(&stu)
	usersCollection := client.Database("testing").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	results, _ := usersCollection.InsertOne(ctx, stu)
	json.NewEncoder(w).Encode(results)
}
