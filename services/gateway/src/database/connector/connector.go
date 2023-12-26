package connector

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitConnectionToMongo(ctx context.Context) (*mongo.Client, error) {

	mongo_uri := os.Getenv("MONGODB_URI")
	if mongo_uri == "" {
		return nil, fmt.Errorf("Environment vatiable MONGODB_URI not found")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	log.Println("MONGO connection has been established.")

	if err := startVideoDatabase(client); err != nil {
		return nil, err
	}

	return client, nil
}

func startVideoDatabase(client *mongo.Client) error {
	videosCollection := client.Database("videos").Collection("videos")
	// Check if the 'videos' collection exists
	exists, err := videosCollection.CountDocuments(context.Background(), bson.M{})
	if err == nil && exists > 0 {
		return nil
	}

	// Collection 'videos' does not exist, create the collection
	// (You can also perform other initial operations if needed)
	log.Println("Collection 'videos' does not exist. Creating...")

	// Create the collection by inserting a document
	_, err = videosCollection.InsertOne(context.Background(), bson.M{"title": "__test__", "url": "__test__", "id": 0})

	if err != nil {
		return err
	}

	log.Println("Collection 'videos' created with a sample document.")

	return nil
}
