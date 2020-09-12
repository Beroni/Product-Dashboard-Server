package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Export Client of MongoDB
func GetMongoDB() *mongo.Client {

	url := "mongodb+srv://admin:RSuKTBhkharP8PgP@cluster0.bcaip.mongodb.net/gettydb?retryWrites=true&w=majority"

	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
