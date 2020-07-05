package gomongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (coll *MongoCollection) Update(filter, update interface{}) (*mongo.UpdateResult, error) {

	updateResult, err := coll.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}
