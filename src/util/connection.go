package utils

import (
	"cms/config"

	"go.mongodb.org/mongo-driver/mongo"
)

func MongoConnection(collectionName string) *mongo.Collection {

	client := config.GetMongoDB()

	collection := client.Database("dev").Collection(collectionName)

	return collection

}
