package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
	// MongoDB connection string
	connectionString := "mongodb://79.118.145.37:27017/teamCompDB"

	// Create a context with a timeout
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect to MongoDB
	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}
