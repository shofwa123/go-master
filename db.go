package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"

)

func db() *mongo.Client {
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	dbURI := "mongodb://"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")
	clientOptions := options.Client().ApplyURI(dbURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
