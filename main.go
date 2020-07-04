package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Blog struct {
	Title string
	Body  string
}

var client *mongo.Client
var collection *mongo.Collection

// Create database and collection
func CreateCollection(dbname, collectionname string) *mongo.Collection {
	collection = client.Database(dbname).Collection(collectionname)
	log.Println("Collection Created")
	return collection
}

// create Document in the collection
func InsertDocument(content interface{}, collectionn *mongo.Collection) (*mongo.InsertOneResult, error) {
	insertResult, err := collectionn.InsertOne(context.TODO(), content)
	if err != nil {
		return nil, err
	}
	log.Println("document created")
	return insertResult, err
}

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Connection Error : ", err)
	}

	log.Println("Connected to MongoDB!")
}

func main() {
	// Create a collection
	collectionn := CreateCollection("Pokemon", "Trainers")

	blog1 := Blog{Title: "Title One", Body: "BODY of First Blog1"}

	insertResult, err := InsertDocument(&blog1, collectionn)
	if err != nil {
		log.Fatal("Insert document error : ", err)
	}
	log.Println("insertresult id : ", insertResult.InsertedID)

	// Read All
	// cursor, err := collectionn.Find(context.TODO(), bson.M{})
	// if err != nil {
	// 	log.Fatal("cursor err :", err)
	// }
	// var episodes []bson.M
	// if err = cursor.All(context.TODO(), &episodes); err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("=========", episodes)

	// cursor, err := collectionn.Find(context.TODO(), bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cursor.Close(context.TODO())
	// for cursor.Next(context.TODO()) {
	// 	var episode bson.M
	// 	if err = cursor.Decode(&episode); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println("++++", episode)
	// }

	// read by id -------------
	id := insertResult.InsertedID
	var result Blog

	err = collectionn.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found a single document: %+v\n", result)
	//--------------------------

	// var readblog []bson.M
	// if err = filterCursor.All(context.TODO(), &readblog); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("read blog :", readblog)

	// Update by id-------------
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"Title", "JOHN-CENA-TITLE"},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("Update eerrorrrr: ------- : ", err)
	}
	log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	//-----------------

	// Delete by id-----------------
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// -----------------------------

}

// func Read(collectionn *mongo.Collection, id interface{}) []bson.M {
// 	filterCursor, err := collectionn.Find(context.TODO(), bson.M{"_id": id})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var readblog []bson.M
// 	if err = filterCursor.All(context.TODO(), &readblog); err != nil {
// 		log.Fatal(err)
// 	}
// 	return readblog
// }
