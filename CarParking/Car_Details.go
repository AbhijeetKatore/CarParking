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

	_,err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Car Details Inserted Successfully")

}

func DeleteCarDetails(){
	fmt.Println("Enter Car Details to Delete")
	fmt.Println("Enter Exact Car Number ")
	fmt.Scanln(&carNumber)
	fmt.Println("Enter Exact Car Model Name ")
	fmt.Scanln(&carModel)

	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	result, err := collection.DeleteMany(ctx, bson.D{{"Car Number", carNumber}, {"Car Model Name", carModel}})
	if err != nil {
    	log.Fatal(err)
	}
	if result.DeletedCount == 0{
		fmt.Println("Data didn't Match to Delete")
	}else{
		fmt.Printf("Car Details Deleted Succesfully")
	}

}

func UpdateCarDetails(){}

