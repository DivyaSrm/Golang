package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// struct for storing data
type category struct {
	Categoryid          int    `json:categoryid`
	CategoryName        string `json:categoryname`
	CategoryDescription string `json:categorydescription`
	CreatedBy           string `json:createdby`
	UpdatedBy           string `json:updatedby`
	CategoryStatus      bool   `json:categorystatus`
}
type ResponseError struct {
	ErrorMessage  string `json:"error message"`
	StatusCode    int    `json:"status code"`
	Status        bool   `json:"status"`
	CustomMessage string `json:"customm message"`
}

type Response struct {
	//ErrorMessage  string `json:"error message"`
	StatusCode    int    `json:"status code"`
	Status        bool   `json:"status"`
	CustomMessage string `json:"customm message"`
}

var categoryCollection = db().Database("usecase").Collection("category") // get collection "users" from db() which returns *mongo.Client

// Create Profile or Signup

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var category category
	err := json.NewDecoder(r.Body).Decode(&category) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	var result primitive.M
	err1 := categoryCollection.FindOne(context.TODO(), bson.D{{"categoryid", category.Categoryid}}).Decode(&result)
	fmt.Println("err", err1, "result", result)
	if result == nil {
		insertResult, err := categoryCollection.InsertOne(context.TODO(), category)
		if err != nil {
			log.Fatal(err)
		}
		msg := Response{
			StatusCode:    200,
			Status:        true,
			CustomMessage: "record inserted"}
		fmt.Println("Inserted a single document: ", insertResult)

		json.NewEncoder(w).Encode(msg) // return the mongodb ID of generated document
	} else {
		msg := ResponseError{
			ErrorMessage:  "nil",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "custom message"}
		json.NewEncoder(w).Encode(msg)
	}
}

// Get Profile of a particular User by Name

func GetCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var category category
	e := json.NewDecoder(r.Body).Decode(&category)
	if e != nil {

		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := categoryCollection.FindOne(context.TODO(), bson.D{{"categoryid", category.Categoryid}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}

//Update Profile of User

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		Categoryid          int    `json:"categoryid"`
		CategoryName        string `json:"categoryname"`        //value that has to be matched
		CategoryDescription string `json:"categorydescription"` // value that has to be modified
		UpdatedBy           string `json:"updatedby"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"categoryid", body.Categoryid}} // converting value to BSON type
	after := options.After                            // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"categorydescription", body.CategoryDescription}, {"categoryname", body.CategoryName},
		{"updatedby", body.UpdatedBy}}}}
	updateResult := categoryCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

//Delete Profile of User

func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string
	categoryid, _ := strconv.Atoi(params)
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := categoryCollection.DeleteOne(context.TODO(), bson.D{{"categoryid", categoryid}}, opts)
	_, err1 := SubcategoryCollection.DeleteMany(context.TODO(), bson.D{{"categoryid", categoryid}}, opts)
	_, err2 := BrandCollection.DeleteMany(context.TODO(), bson.D{{"categoryid", categoryid}}, opts)
	if err != nil || err1 != nil || err2 != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                       //slice for multiple documents
	cur, err := categoryCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	fmt.Println(cur)
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}
