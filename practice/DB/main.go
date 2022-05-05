package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	usersCollection := client.Database("testing").Collection("users")

	// insert a single document into a collection

	// create a bson.D object

	user := bson.D{primitive.E{Key: "fullName", Value: "User 1"}, {Key: "age", Value: 30}}

	// insert the bson object using InsertOne()

	result, err := usersCollection.InsertOne(context.TODO(), user)

	// check for errors in the insertion

	if err != nil {

		panic(err)

	}

	string

	// display the id of the newly inserted object

	fmt.Println(result.InsertedID)

	// insert multiple documents into a collection

	// create a slice of bson.D objects

	users := []interface{}{

		bson.D{primitive.E{Key: "fullName", Value: "User 2"}, {Key: "age", Value: 25}},

		bson.D{primitive.E{Key: "fullName", Value: "User 3"}, {Key: "age", Value: 20}},

		bson.D{primitive.E{Key: "fullName", Value: "User 4"}, {Key: "age", Value: 28}},
	}

	// insert the bson object slice using InsertMany()

	results, err := usersCollection.InsertMany(context.TODO(), users)

	// check for errors in the insertion

	if err != nil {

		panic(err)

	}

	// display the ids of the newly inserted objects

	fmt.Println(results.InsertedIDs)

	cursor, err := usersCollection.Find(context.TODO(), bson.D{})

	// convert the cursor result to bson
	var result1 []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &result1); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results in a collection")
	for _, result1 := range result1 {
		fmt.Println(result1)
	}

}
