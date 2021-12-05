package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Connected to MongoDB")
	return client
}

//declear a client database instance
var Client *mongo.Client = DBInstance()

//OpenCollection functon establish a connection to the collection in the mongo database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	DatabaseName := os.Getenv("DATABASE_NAME")
	var collection *mongo.Collection = client.Database(DatabaseName).Collection(collectionName)
	return collection
}
