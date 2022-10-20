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
			fmt.Println("Enter Occupancy As true for Free or false for Occupied Only")
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
	filter := bson.D{{"Occupancy",true}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	var result []bson.D

	cursor.All(ctx, &result)
	cursor.Decode(&result);
	fmt.Println("Below are Free Parking Slots Available Now")
	
	for _,item:=range result {   
		fmt.Println("Floor No",item[1].Value,"and Slot Number",item[2].Value,"is Free Now ")
	}

}
var carNumberV string
var firstNameV string
var lastNameV string
var floorNumberV string
var slotNumberV string

func AddNewCarToSlot() {
	 
	// start := time.Now()
	// fmt.Println(start.Hour(),":",start.Minute())

	// fmt.Println("Enter User's First Name")
	// fmt.Scan(&firstNameV)
	// fmt.Println("Enter User's Last Name")
	// fmt.Scan(&lastNameV)
	fmt.Println("Enter Car Number to Add as Parked ")
	fmt.Scan(&carNumberV)
	// fmt.Println("Enter Floor Number to Block Parking Slot")
	// fmt.Scan(&floorNumber)
	// fmt.Println("Enter Slot Number to Block Parking Slot")
	// fmt.Scan(&slotNumber)

	//check user available in database

	//check car number in databse
	checkCarInDatabase()
	//check slot available in database

}

func checkCarInDatabase(){
	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	var result bson.M
	err := collection.FindOne(ctx, bson.D{{"Car Number",carNumber}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Car Number Not Found in Existing Database please add Car Details First or Enter a Correct One")
		}
		
	}
}


func checkUserInDatabase(){
	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	var result bson.M
	err := collection.FindOne(ctx, bson.D{{"First Name",firstNameV},{"Last Name",lastNameV}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Users Name Not Found in Existing Database please add User Details First or Enter a Correct One")
		}
		
	}

}
func checkParkingSlotInDatabase(){

}

func RemoveCarFromSlot(){
	// var client *mongo.Client 
	// var ctx context.Context
	// client,ctx = ConnectDatabase()
	// collection := client.Database("CarParking").Collection("ParkingSlots")
}