package CarParking

import "fmt"

var carNumber string
var carModel string

func AddCarDetails(){
	fmt.Println("Enter Car Number ")
	fmt.Scanln(&carNumber)
	fmt.Println("Enter Car Model Name ")
	fmt.Scanln(&carModel)
}

func DeleteCarDetails(){}

func UpdateCarDetails(){}