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

type subcategory struct {
	Subcategoryid          int    `json:subcategoryid`
	Categoryid             int    `json:categoryid`
	SubcategoryName        string `json:subcategoryname`
	SubcategoryDescription string `json:subcategorydescription`
	CreatedBy              string `json:createdby`
	UpdatedBy              string `json:updatedby`
}

var SubcategoryCollection = db().Database("usecase").Collection("subcategory") // get collection "users" from db() which returns *mongo.Client

// Create Profile or Signup

func CreateSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var subcategory subcategory
	err := json.NewDecoder(r.Body).Decode(&subcategory) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	insertResult, err := SubcategoryCollection.InsertOne(context.TODO(), subcategory)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the mongodb ID of generated document

}

// Get Profile of a particular User by Name

func GetSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var subcategory subcategory
	e := json.NewDecoder(r.Body).Decode(&subcategory)
	if e != nil {

		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := SubcategoryCollection.FindOne(context.TODO(), bson.D{{"subcategoryid", subcategory.Subcategoryid}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}

//Update Profile of User

func UpdateSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		Categoryid             int    `json:"categoryid"`
		SubCategoryid          int    `json:"subcategoryid"`          //value that has to be matched
		SubCategoryDescription string `json:"subcategorydescription"` // value that has to be modified
		SubCategoryName        string `json:"subcategoryname"`        // value that has to be modified
		UpdatedBy              string `json:"updatedby"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"subcategoryid", body.SubCategoryid}} // converting value to BSON type
	after := options.After                                  // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"categoryid", body.Categoryid}, {"subcategorydescription", body.SubCategoryDescription},
		{"subcategoryname", body.SubCategoryName}, {"updatedby", body.UpdatedBy}}}}
	updateResult := SubcategoryCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

//Delete Profile of User

func DeleteSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	subcategoryid, _ := strconv.Atoi(params)
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := SubcategoryCollection.DeleteOne(context.TODO(), bson.D{{"subcategoryid", subcategoryid}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func GetAllSubCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                          //slice for multiple documents
	cur, err := SubcategoryCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
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
