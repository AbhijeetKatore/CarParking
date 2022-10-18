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

func AddUser() {
	fmt.Println("Enter First Name ")
	fmt.Scanln(&fName)
	fmt.Println("Enter Last Name ")
	fmt.Scanln(&lName)

	var client *mongo.Client 
	var ctx context.Context
	client,ctx = ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	doc := bson.D{{"First Name", fName}, {"Last Name", lName}}
	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User Details Inserted Successfully")


}
func Deleteuser() {

}