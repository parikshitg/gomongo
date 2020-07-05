package gomongo

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCollection struct {
	Collection *mongo.Collection
}

func (c *MongoDB) CreateCollection(dbname, collectionname string) *MongoCollection {

	var col = &MongoCollection{}
	col.Collection = c.Client.Database(dbname).Collection(collectionname)

	log.Println("Collection Created")
	return col
}
