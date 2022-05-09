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
type variant struct {
	Variantid   int    `json:variantid`
	Variantname string `json:variantname`
	CreatedBy   string `json:createdby`
	UpdatedBy   string `json:updatedby`
}

var VariantCollection = db().Database("usecase").Collection("variant") // get collection "users" from db() which returns *mongo.Client

// Create Profile or Signup

func CreateVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var variant variant
	err := json.NewDecoder(r.Body).Decode(&variant) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	errvariant := VariantCollection.FindOne(context.TODO(), bson.D{{"variantid", variant.Variantid}}).Decode(&result)
	if errvariant != nil {
		_, err := VariantCollection.InsertOne(context.TODO(), variant)
		if err != nil {
			log.Fatal(err)
		}
		msg := Response{
			StatusCode:    200,
			Status:        true,
			CustomMessage: "record inserted"}

		json.NewEncoder(w).Encode(msg) // return the mongodb ID of generated document
	} else {
		msg := ResponseError{
			ErrorMessage:  "nil",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "record not inserted"}
		json.NewEncoder(w).Encode(msg)
	}

}

// Get Profile of a particular User by Name

func GetVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var variant variant
	e := json.NewDecoder(r.Body).Decode(&variant)
	if e != nil {

		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := VariantCollection.FindOne(context.TODO(), bson.D{{"varianatid", variant.Variantid}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}

//Update Profile of User

func UpdateVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		VariantId   int    `json:"variantid"`
		VariantName string `json:"variantname"` //value that has to be matched
		CreatedBy   string `json:"createdby"`   // value that has to be modified
		UpdatedBy   string `json:"updatedby"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"variantid", body.VariantId}} // converting value to BSON type
	after := options.After                          // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"variantname", body.VariantName},
		{"updatedby", body.UpdatedBy}}}}
	updateResult := VariantCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)
	if result == nil {
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode("record not available")
	}

}

//Delete Profile of User

func DeleteVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	variantid, _ := strconv.Atoi(params) // convert params to mongodb Hex ID

	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := VariantCollection.DeleteOne(context.TODO(), bson.D{{"variantid", variantid}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func GetAllVariant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                      //slice for multiple documents
	cur, err := VariantCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
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
