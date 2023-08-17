package main

import (
	"project1/services"
	// "project1/services"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
    mongoClient *mongo.Client
)

func main() {
   // fmt.Println("MongoDB successfully connected...")

    // products := []interface{}{
    //  models.Product{ID: primitive.NewObjectID(), Name: "OnePlus", Price: 1000000, Description: "Budget Phone"},
    //  models.Product{ID: primitive.NewObjectID(), Name: "Vivo", Price: 100000, Description: "China based Phone"},
    services.UpdateTransaction()

}

