package gomongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (coll *MongoCollection) Read(id, content interface{}) (interface{}, error) {

	err := coll.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&content)
	if err != nil {
		return nil, err
	}
	return content, nil
}
