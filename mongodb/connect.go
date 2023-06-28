package mongodb

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
	"github.com/EnTing0417/go-lib/model"

)

func Init() (client *mongo.Client){

	config := model.ReadConfig()

	// Set the MongoDB connection URI
	connectionURI := config.Database.URI

	// Set options for the MongoDB client
	clientOptions := options.Client().ApplyURI(connectionURI)

	// Create a new MongoDB client
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return client
}


func Connect(client *mongo.Client) {
	// Set a timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB server
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB")
}

func Disconnect(client *mongo.Client) {
	client.Disconnect(context.Background())
	fmt.Println("Disconnected")
}
