package main

import (
	"CarParking/CarParking"
	"fmt"
)
func main(){
for{
	fmt.Print(`
 	Type 1 for Adding User
 	Type 2 for Deleting User
 	Type 3 for Adding Car Details
 	Type 4 for Deleting Car Details
 	Type 5 for Updating Car Details
 	Type 6 for Adding New Parking Slots
 	Type 7 for Deleting Parking Slots
 	Type 8 for Updating Parking Slots
 	Type 9 for To Know Free Parking Slots
 	Type 10 for Adding New Car to Slot along with Time In
 	Type 11 or Removing Car from Slot along with Time Out
`)
	

 	var selector int8
 	fmt.Scanln(&selector)
 	switch selector{
 	case 1: CarParking.AddUser()
 	case 2: CarParking.Deleteuser()
 	case 3 : CarParking.AddCarDetails()
 	case 4: CarParking.DeleteCarDetails()
	case 5: CarParking.UpdateCarDetails()
 	case 6: CarParking.AddParkingSlots()
 	case 7: CarParking.DeleteParkingSlots()
 	case 8: CarParking.UpdateParkingSlot()
	case 9: CarParking.GetFreeParkingSlots()
 	case 10: CarParking.AddNewCarToSlot()
 	case 11: CarParking.RemoveCarFromSlot()
 	default : fmt.Println("Selct proper input from the above list")
	}

}}