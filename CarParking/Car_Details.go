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
	
	data := bson.D{{"Car Number", carNumber}, {"Car Model Name", carModel}}

	_,err := collection.InsertOne(ctx, data)
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
	data :=bson.D{{"Car Number", carNumber}, {"Car Model Name", carModel}}
	result, err := collection.DeleteMany(ctx, data)
	if err != nil {
    	log.Fatal(err)
	}
	if result.DeletedCount == 0{
		fmt.Println("Data didn't Match to Delete")
	}else{
		fmt.Printf("Car Details Deleted Succesfully")
	}

}

func UpdateCarDetails(){

	var oldCarNumber string
	var oldCarModel string
	var newCarNumber string
	var newCarModel string

	fmt.Println("Enter Car Details that you want to Update")
	fmt.Println("Enter Exact Car Number ")
	fmt.Scanln(&oldCarNumber)
	fmt.Println("Enter Exact Car Model Name ")
	fmt.Scanln(&oldCarModel)

	fmt.Println("Enter New Car Details")
	fmt.Println("Enter New Car Number ")
	fmt.Scanln(&newCarNumber)
	fmt.Println("Enter New Car Model Name ")
	fmt.Scanln(&newCarModel)

	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	prevData :=bson.D{{"Car Number", oldCarNumber}, {"Car Model Name", oldCarModel}}
	newData:=bson.D{{"$set",bson.D{{"Car Number", newCarNumber}, {"Car Model Name", newCarModel}}},}
	result, err := collection.UpdateMany(ctx,prevData,newData)
	if err != nil {
		log.Fatal(err)
	}
	if result.ModifiedCount == 0{
		fmt.Println("Data didn't Match to Update")
	}
	fmt.Printf("New Car Details Updated Succesfully")

}