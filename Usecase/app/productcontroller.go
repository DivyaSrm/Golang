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
type product struct {
	Productid   int            `json:productid`
	ProductName string         `json:productname`
	Product     map[string]int `json:product`
	CreatedBy   string         `json:createdby`
	UpdatedBy   string         `json:updatedby`
}
type ids struct {
	categoryid    int `json:categoryid`
	subcategoryid int `json:subcategoryid`
}

var ProductCollection = db().Database("usecase").Collection("products") // get collection "users" from db() which returns *mongo.Client

// Create Profile or Signup

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var product product
	err := json.NewDecoder(r.Body).Decode(&product) // storing in person variable of type user
	if err == nil {
		fmt.Print(err)
	}
	var result primitive.M
	err1 := ProductCollection.FindOne(context.TODO(), bson.D{{"productid", product.Productid}}).Decode(&result)
	fmt.Println(err1)
	if err1 == nil {
		msg := ResponseError{
			ErrorMessage:  "nil",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "id already exist"}
		json.NewEncoder(w).Encode(msg)
	} else {
		insertResult, err := ProductCollection.InsertOne(context.TODO(), product)
		if err != nil {
			log.Fatal(err)
		}

		msg := Response{
			StatusCode:    200,
			Status:        true,
			CustomMessage: "record inserted"}
		fmt.Println("Inserted a single document: ", insertResult)

		json.NewEncoder(w).Encode(msg)

	}

}

// Get Profile of a particular User by Name

func GetProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var product product
	e := json.NewDecoder(r.Body).Decode(&product)
	if e != nil {

		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := ProductCollection.FindOne(context.TODO(), bson.D{{"productid", product.Productid}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}

//Update Profile of User

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		Productid          int    `josn:"productid"`
		ProductName        string `json:"productname"`        //value that has to be matched
		ProductDescription string `json:"productdescription"` // value that has to be modified
		UpdatedBy          string `json:"updatedby"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"productid", body.Productid}} // converting value to BSON type
	after := options.After                          // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"productname", body.ProductName}, {"updatedby", body.UpdatedBy}}}}
	updateResult := ProductCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

//Delete Profile of User

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	productid, _ := strconv.Atoi(params)
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := ProductCollection.DeleteOne(context.TODO(), bson.D{{"productid", productid}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                      //slice for multiple documents
	cur, err := ProductCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
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
