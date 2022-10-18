package CarParking

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

var carNumber string
var carModel string

func AddCarDetails(){
	fmt.Println("Enter Car Number ")
	fmt.Scanln(&carNumber)
	fmt.Println("Enter Car Model Name ")
	fmt.Scanln(&carModel)

	userData := bson.D{{"First Name",fName},{"Last Name",lName}}
	fmt.Print(userData)

}

func DeleteCarDetails(){}

func UpdateCarDetails(){}

