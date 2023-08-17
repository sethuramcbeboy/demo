package config

//import constants and options
//planning to provde context
// run the context for given ammount of time 

import(

	"context"
	"project1/Constants"
	"fmt"
	
	"time"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
	
)
func ConnectDataBase() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoConnection := options.Client().ApplyURI(Constants.ConnectionString)
	mongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return mongoClient, nil
}

func GetCollection(client *mongo.Client,dbName string,collectionName string) *mongo.Collection{
	client,err := ConnectDataBase()
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}