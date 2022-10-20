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
	var floorNumber int8
	var slotNumber int8
	fmt.Println("Enter New Parking Slot Details")
	fmt.Println("Enter Floor Number")
	fmt.Scan(&floorNumber)
	fmt.Println("Enter Unique Slot Number")
	fmt.Scan(&slotNumber)
		
	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	data :=bson.D{{"Floor Number", floorNumber}, {"Unique Slot Number", slotNumber}}
	result, err := collection.DeleteMany(ctx, data)
	if err != nil {
    	log.Fatal(err)
	}
	if result.DeletedCount == 0{
		fmt.Println("Parking Slot didn't Match to Delete")
	}else{
		fmt.Printf("Parking Slot Deleted Succesfully")
	}

}

func UpdateParkingSlot() {
	var oldFloorNumber int8
	var oldSlotNumber int8
	fmt.Println("Enter Old Parking Slot Details")
	fmt.Println("Enter Exact Floor Number")
	fmt.Scan(&oldFloorNumber)
	fmt.Println("Enter Exact Unique Slot Number")
	fmt.Scan(&oldSlotNumber)

	var newFloorNumber int8
	var newSlotNumber int8
	var newOccupancy string
	var newOccupancyFlag bool
	fmt.Println("Enter New Parking Slot Details")
	fmt.Println("Enter Floor Number")
	fmt.Scan(&newFloorNumber)
	fmt.Println("Enter Unique Slot Number")
	fmt.Scan(&newSlotNumber)
		for{
			fmt.Println("Enter Occupancy As true or false Only")
			fmt.Scan(&newOccupancy)
			if  newOccupancy == "true" ||  newOccupancy == "false"{
				if newOccupancy == "true" {
					newOccupancyFlag=true
				}
				if newOccupancy == "false"{
					newOccupancyFlag=false
				}
				break
			}
		}

	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	prevData :=bson.D{{"Floor Number", oldFloorNumber}, {"Unique Slot Number", oldSlotNumber}}
	newData:=bson.D{{"$set",bson.D{{"Floor Number", newFloorNumber}, {"Unique Slot Number", newSlotNumber},{"Occupancy", newOccupancyFlag}}},}
	result, err := collection.UpdateMany(ctx,prevData,newData)
	if err != nil {
		log.Fatal(err)
	}
	if result.ModifiedCount == 0{
		fmt.Println("Data didn't Match to Update")
	}
	fmt.Printf("Parking Slot Details Updated Succesfully")


}

func GetFreeParkingSlots(){
	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")

	cursor, err := collection.Find(ctx, bson.M{"Occupancy":true})
	if err != nil {
   		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
   		var availableSlots bson.M
    	if err = cursor.Decode(&availableSlots); err != nil {
        	log.Fatal(err)
    	}
    	fmt.Println(availableSlots)
	}
}
func AddNewCarToSlot() {

}
