package services

import (
	"context"
	"encoding/json"
	"fmt"
	"project1/config"

	//"project1/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func TransactionContext() *mongo.Collection {
	client, _ := config.ConnectDataBase()
	return config.GetCollection(client, "inventory", "students")
}

// func FetchAggregatedTransactions() {
// 	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
// 	matchStage := bson.D{{"$match", bson.D{{"account_id", 278866}}}}
// 	groupStage := bson.D{
// 		{"$group", bson.D{
// 			{"_id", "$account_id"},
// 			{"total_count", bson.D{{"$sum", "$transaction_count"}}},
// 		}}}
// 	result, err := TransactionContext().Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
// 	if err != nil {
// 		fmt.Println(err.Error)
// 	} else {
// 		var showsWithInfo []bson.M
// 		if err = result.All(ctx, &showsWithInfo); err != nil {
// 			panic(err)
// 		}

// 		formatted_data, err := json.MarshalIndent(showsWithInfo, "", " ")
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		} else {
// 			fmt.Println(string(formatted_data))
// 		}

// 	}
// }

func UpdateTransaction(initialvalue int, newvalue int) {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.D{{"account_id", initialvalue}}
	update := bson.D{{"$set", bson.D{{"account_id", newvalue}}}}
	result, err := TransactionContext().UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return result, nil
}
