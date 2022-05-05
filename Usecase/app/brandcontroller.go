package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// struct for storing data
type brand struct {
	Brandid          int    `json:brandid`
	Categoryid       int    `json:categoryid`
	Subcategoryid    int    `json:subcategoryid`
	BrandName        string `json:brandname`
	BrandDescription string `json:branddescription`
	CreatedBy        string `json:createdby`
	UpdatedBy        string `json:updatedby`
}

var BrandCollection = db().Database("usecase").Collection("brand") // get collection "users" from db() which returns *mongo.Client

// Create Profile or Signup

func CreateBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var brand brand
	err := json.NewDecoder(r.Body).Decode(&brand) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	insertResult, err := BrandCollection.InsertOne(context.TODO(), brand)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the mongodb ID of generated document

}

// Get Profile of a particular User by Name

func GetBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var category category
	e := json.NewDecoder(r.Body).Decode(&category)
	if e != nil {

		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := BrandCollection.FindOne(context.TODO(), bson.D{{"categoryname", category.CategoryName}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}

//Update Profile of User

func UpdateBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		CategoryName        string `json:"categoryname"`        //value that has to be matched
		CategoryDescription string `json:"categorydescription"` // value that has to be modified
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"categoryname", body.CategoryName}} // converting value to BSON type
	after := options.After                                // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"categorydescription", body.CategoryDescription}}}}
	updateResult := BrandCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

//Delete Profile of User

func DeleteBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := BrandCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func GetAllBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                    //slice for multiple documents
	cur, err := BrandCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
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
