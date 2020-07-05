package gomongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (coll *MongoCollection) Delete(filter interface{}) (*mongo.DeleteResult, error) {

	deleteResult, err := coll.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}
