package main

import (
	"fmt"
)

/*
	Define hash(plate, length):
		SUM = ASCII of plate % length

	define insert(plate, parking_lot):
		index = hash(plate, length)
		while parking_lot[index] is occupied:
			move to the next slot (index = (index + one) % lot_size)
		place the plate in the parking_lot[index]

	define lookup(plate, parking_lot):
		index = hash(plate, lot_size)
		while paring_lot[index] is not empty:
			if parking_lot[index] == plate:
				return true
			move to next slot(index = (index + one % lot_size))
		return false

	MAIN:
		Create an empty parking_lot of size 7
		insert KBA123, KDC456, KAA789
		Ask for plate_to_serarch
		if lookup(plate_to_search) is true:
			Print "Car found"
		else:
			Print "Car not found"

*/

func hash(plate string, length int) int{
	sum := 0

	for _, char := range plate {
		sum += int(char)
	}

	return sum % length
}

func insert(plate string, parking_lot []string) {
	index := hash(plate, len(parking_lot))

	start := index // To detect if we've looped around
	for parking_lot[index] != "" {
		index = (index + 1) % len(parking_lot)
		if index == start {
			fmt.Println("Parking lot is full")
			return
		}
	}
	parking_lot[index] = plate

}

func lookup(plate string, lot []string) bool {
	index := hash(plate, len(lot))
	start := index

	for lot[index] != "" {
		if lot[index] == plate {
			return true
		}
		index = (index + 1) % len(lot)
		if index == start {
			break
		}
	}
	return false
}

func main() {
	parking_lot := make([]string, 7)
	
	insert("KBA123", parking_lot)
	insert("KDC456", parking_lot)
	insert("KAA789", parking_lot)

	var plate string
	fmt.Println("Enter license plate to search:")
	fmt.Scanln(&plate)

	if lookup(plate, parking_lot) {
		fmt.Println("Car found.")
	} else {
		fmt.Println("Car not found")
	}


}