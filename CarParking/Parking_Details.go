package CarParking

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddParkingSlots() {
	var floorNumber int8
	var slotNumber int8
	var occupancy string
	var occupancyFlag bool
	fmt.Println("Enter New Parking Slot Details")
	fmt.Println("Enter Floor Number")
	fmt.Scan(&floorNumber)
	fmt.Println("Enter Unique Slot Number")
	fmt.Scan(&slotNumber)

	
		for{
		fmt.Println("Enter Occupancy As true or false Only")
		fmt.Scan(&occupancy)
		if  occupancy == "true" ||  occupancy == "false"{
			if occupancy == "true" {
				occupancyFlag=true
			}
			if occupancy == "false"{
				occupancyFlag=false
			}
			break
		}
	}
	
	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	data := bson.D{{"Floor Number", floorNumber}, {"Unique Slot Number", slotNumber},{"Occupancy", occupancyFlag}}
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Parking Slot Details Inserted Successfully")


}

func DeleteParkingSlots() {

}

func GetFreeParkingSlots() {

}

func AddNewCarToSlot() {

}
