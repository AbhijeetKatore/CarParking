package CarParking

import (
	"context"
	"fmt"
	"log"
	"time"

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
	filter := bson.D{{"Occupancy",false}}
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
var floorNumberV int8
var slotNumberV int8
var carFound bool
var userFound bool
var slotFound bool

func AddNewCarToSlot() {
	
	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots") 

	fmt.Println("Enter User's First Name")
	fmt.Scan(&firstNameV)
	fmt.Println("Enter User's Last Name")
	fmt.Scan(&lastNameV)
	userFound = checkUserInDatabase()

	fmt.Println("Enter Car Number to Add as Parked ")
	fmt.Scan(&carNumberV)
	carFound = checkCarInDatabase()

	fmt.Println("Enter Floor Number to Block Parking Slot")
	fmt.Scan(&floorNumberV)
	fmt.Println("Enter Slot Number to Block Parking Slot")
	fmt.Scan(&slotNumberV)
	slotFound = checkParkingSlotInDatabase()

	if userFound == true && carFound == true && slotFound == true{
		prevData :=bson.D{{"Floor Number", floorNumberV}, {"Unique Slot Number", slotNumberV}}
		newData:=bson.D{{"$set",bson.D{{"Floor Number", floorNumberV}, {"Unique Slot Number", slotNumberV},{"Occupancy", true},{"Car Number",carNumberV},{"First Name",firstNameV},{"Last Name",lastNameV},{"TimeIn",time.Now()}}},}
		result, err := collection.UpdateMany(ctx,prevData,newData)
		if err != nil {
			log.Fatal(err)
		}
		if result.ModifiedCount == 0{
			fmt.Println("Data didn't Match to Update")
		}
		fmt.Printf("Parking Slot Details Updated Succesfully")
	}

}

func checkCarInDatabase()(bool){
	var found bool 
	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	var result bson.M
	err := collection.FindOne(ctx, bson.D{{"Car Number",carNumberV}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Car Number Not Found in Existing Database please add Car Details First or Enter a Correct One",result)
		}
		
	}else{
		found = true
	}
	return found
}


func checkUserInDatabase()(bool){
	var found bool
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
	}else{
		found = true
	}
	return found

}
func checkParkingSlotInDatabase()(bool){
	var client *mongo.Client 
	var ctx context.Context
	var slotFound bool
	var isFree bool 
	var allTrue bool

	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	var result bson.M
	err := collection.FindOne(ctx, bson.D{{"Floor Number",floorNumberV},{"Unique Slot Number",slotNumberV}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Car Slot Not Found in Existing Database please add Slot Details First or Enter a Correct One")
		}
	}else{
		slotFound =true
	}
	result=nil
	nerr := collection.FindOne(ctx, bson.D{{"Floor Number",floorNumberV},{"Unique Slot Number",slotNumberV},{"Occupancy",false}}).Decode(&result)
	if nerr != nil {
		if nerr == mongo.ErrNoDocuments {
			fmt.Println("Car Slot is Occupied Now Please Enter a Different One")
		}
	}else{
		isFree =true
	}

	if slotFound == true && isFree == true {
		allTrue = true
	}
	return allTrue
	
}


func RemoveCarFromSlot(){

	fmt.Println("Enter Car Number to Remove From Parking ")
	fmt.Scan(&carNumber)

	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")

	prevData :=bson.D{{"Car Number",carNumber}}
	newData:=bson.D{{"$set",bson.D{{"Occupancy", true},{"outTime",time.Now()}}},}
	result, err := collection.UpdateMany(ctx,prevData,newData)
	if err != nil {
		log.Fatal(err)
	}
	if result.ModifiedCount == 0{
		fmt.Println("Data didn't Match to Update",result)
	}else{

	filter := bson.D{{"Car Number",carNumber}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var results []bson.D
	cursor.All(ctx, &results)
	cursor.Decode(&results);
	for _,item:=range results {   
		data := bson.D{{"Floor Number", item[1].Value}, {"Unique Slot Number", item[2].Value},{"Occupancy", false}}
		_, err := collection.InsertOne(ctx, data)
		if err != nil {
			log.Fatal(err)
		}

	}
	fmt.Printf("Parking Slot Details Updated Succesfully")
	}


}