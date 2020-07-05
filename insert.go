package gomongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (coll *MongoCollection) Insert(content interface{}) (*mongo.InsertOneResult, error) {

	insertResult, err := coll.Collection.InsertOne(context.TODO(), content)
	if err != nil {
		return nil, err
	}

	return insertResult, err
}
