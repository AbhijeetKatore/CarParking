package CarParking

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var carNumber string
var carModel string

func AddCarDetails(){
	fmt.Println("Enter Car Number ")
	fmt.Scanln(&carNumber)
	fmt.Println("Enter Car Model Name ")
	fmt.Scanln(&carModel)

	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	doc := bson.D{{"Car Number", carNumber}, {"Car Model Name", carModel}}
	
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

}

func DeleteCarDetails(){}

func UpdateCarDetails(){}

