package CarParking

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
var fName string
var lName string
var age int8

func AddUser() {
	fmt.Println("Enter First Name ")
	fmt.Scanln(&fName)
	fmt.Println("Enter Last Name ")
	fmt.Scanln(&lName)
	fmt.Println("Enter Age")
	fmt.Scanln(&age)


	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	doc := bson.D{{"First Name", fName}, {"Last Name", lName},{"Age", age}}
	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User Details Inserted Successfully")

}
func Deleteuser() {
	fmt.Println("Enter Exact Details of User that you want to Delete  ")
	fmt.Println("Enter Exact First Name  ")
	fmt.Scanln(&fName)
	fmt.Println("Enter Exact Last Name ")
	fmt.Scanln(&lName)
	fmt.Println("Enter Exact Age")
	fmt.Scanln(&age)

	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	result, err := collection.DeleteMany(ctx, bson.D{{"First Name", fName}, {"Last Name", lName},{"Age", age}})
	if err != nil {
    	log.Fatal(err)
	}
	if result.DeletedCount == 0{
		fmt.Println("Data didn't Match to Delete")
	}else{
		fmt.Printf("User Details Deleted Succesfully")
	}
}