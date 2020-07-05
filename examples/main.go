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

	uri := "mongodb://localhost:27017"

	var err error
	client, err = gomongo.Connect(uri)
	if err != nil {
		log.Fatal("Connection Error : ", err)
	}
	log.Println("client : ", client)

	// Create Collection
	collection = client.CreateCollection("Cricket", "India")
	log.Println("collection : ", collection)

	// Insert into Collection
	blog1 := Blog{Title: "Title One", Body: "BODY of First Blog ONe"}
	insertResult, err := collection.Insert(&blog1)
	if err != nil {
		log.Fatal("Insert document error : ", err)
	}
	log.Println("insertresult id : ", insertResult.InsertedID)

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
			{"title", "JOHN-CENA-TITLE"},
		}},
	}

	updateResult, err := collection.Update(filter, update)
	if err != nil {
		log.Fatal("Update eerrorrrr: ------- : ", err)
	}
	log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
