package main

import (
	"log"

	"github.com/parikshitg/gomongo"
	"go.mongodb.org/mongo-driver/bson"
)

var client *gomongo.MongoDB
var collection *gomongo.MongoCollection

type Blog struct {
	Title string
	Body  string
}

func main() {

	// Create Mongo Client
	uri := "mongodb://localhost:27017"
	var err error
	client, err = gomongo.Connect(uri)
	if err != nil {
		log.Fatal("Connection Error : ", err)
	}

	// Close Connection
	defer client.Close()

	// Create Collection
	collection = client.CreateCollection("BlogDB", "Blogs")

	// Insert into Collection
	blog1 := Blog{Title: "First Blog Title", Body: "BODY of First Blog."}
	insertResult, err := collection.Insert(&blog1)
	if err != nil {
		log.Fatal("Insert Error : ", err)
	}
	log.Println("InsertResult id : ", insertResult.InsertedID)

	// Read from collection
	id := insertResult.InsertedID
	var result Blog
	res, err := collection.Read(id, result)
	if err != nil {
		log.Fatal("Read Error : ", err)
	}
	log.Printf("Found a document: %+v\n", res)

	// update
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"title", "UPDATED-TITLE"},
		}},
	}

	updateResult, err := collection.Update(filter, update)
	if err != nil {
		log.Fatal("Update Error : ", err)
	}
	log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// delete
	deleteResult, err := collection.Delete(filter)
	if err != nil {
		log.Fatal("Delete Error : ", err)
	}
	log.Printf("Deleted %v documents in the Blogs collection\n", deleteResult.DeletedCount)

}
