package main

import "fmt"

/*
	define hash(number_plate, length of parking_lot):
		sum = sum of ASCII of number_plate
		return sum % length


	define insert(number_plate, parking_lot):
		index = hash(number_plate, length of parking_lot)
		start = index
		for each number_plate[index] is occupied and parking_lot[index] != DELETED:
			move to the next slot index = (index + one) % length of parking_lot
			if start == index:
				parking lot is full
		place index in the parking lot


	define lookup(number_plate, paring_lot):
		index = hash(number_plate, length of parking_lot)
		start = index
		while parking_lot[index] is not empty:
			if paring_lot[index] == number_plate:
				return true
			move to next slot index = (index + one) % length of parking_lot
		return false


	define delete(number_plate, parking_lot):
		index = hash(number plate, length of parking_lot)
		start = index // to detect if we have checked all slots

		while parking_lot[index] is not empty:
			if parking_lot[index] is equal to number_plate:
				mark slot as "DELETED"
				print "Car removed successfully"
				return
			move to next slot -> index = (index + one) % length of parking_lot
			if index == start:
				print "car not found"
				return
		print "Car not found" -> reached an empty slot, car not there
	

	define printLot(parking_lot):
		Print "Parking lot state"
		for each index i and value v in parking_lot:
			if v is empty:
				print "[i] empty"
			else:
				print "[i] v"


	Main:
		Create an array of empty buckets 'number_plate', size N
		insert each car number_plate
		if lookup == true:
			print "Car found"
		else:
			print "Car not found"
*/

func hash(number_plate string, length int) int {
	sum := 0
	for _, char := range number_plate {
		sum += int(char)
	}
	return sum % length

}


func insert(number_plate string, parking_lot []string) {
	// Prevent duplicates
	if lookup(number_plate, parking_lot) {
		fmt.Println("Car is already parked")
		return
	}

	index := hash(number_plate, len(parking_lot))
	start := index

	for parking_lot[index] != "" && parking_lot[index] != "DELETED" {
		index = (index + 1) % len(parking_lot)

		if index == start {
			fmt.Println("Parking is full")
			return
		}
	}
	parking_lot[index] = number_plate
	fmt.Println("Car parked successfully", number_plate)
}


func lookup(plate string, parking_lot []string) bool{
	index := hash(plate, len(parking_lot))
	start := index

	for parking_lot[index] != "" {
		if parking_lot[index] == plate {
			return true
		}

		index = (index + 1) % len(parking_lot)
		if index == start {
			break
		}
	}
	return false
}



func deleteCar(number_plate string, parking_lot []string) {
	index := hash(number_plate, len(parking_lot))
	start := index

	for parking_lot[index] != "" {
		if parking_lot[index] == number_plate {
			parking_lot[index] = "DELETED"
			fmt.Println("Car removed successfully")
			return
		}
		
		index = (index + 1) % len(parking_lot)
		if index == start {
			fmt.Println("Car not found", number_plate)
			return
		}

	}
	fmt.Println("Car not found", number_plate)
}



func printLot(parking_lot []string) {
	fmt.Println("Parking lost state")
	for i, v := range parking_lot {
		if v == "" {
			fmt.Printf(" [%d] empty\n", i)
		} else {
			fmt.Printf(" [%d] %s\n", i, v)
		}
	}
}


func main() {
	parking_lot := make([]string, 10)
	var choice int

	for {
		fmt.Println("\nChoose an option")
		fmt.Println("1) Park car")
		fmt.Println("2) Find car")
		fmt.Println("3) Show parking lot (debug)")
		fmt.Println("4) Remove a car from parking lot")
		fmt.Println("5) Bye")

		_, err := fmt.Scanln(&choice)

		if err != nil {
			// Clear bad input
			fmt.Println("Please enter a number 1-4")
			var discard string
			fmt.Scanln(&discard)
			continue
		} 

	switch choice {
		case 1:
			fmt.Println("Enter the number plate of your car:")
			var plate string
			fmt.Scanln(&plate)
			insert(plate, parking_lot)
		
		case 2:
			fmt.Println("Enter the plate to lookup:")
			var plate string
			fmt.Scanln(&plate)
			if lookup(plate, parking_lot) {
				fmt.Println("Car found")
			} else {
				fmt.Println("Car not found")
			}
		
		case 3:
			printLot(parking_lot)

		case 4:
			fmt.Println("Enter the number plate to delete:")
			var plate string
			fmt.Scanln(&plate)
			deleteCar(plate, parking_lot)

		case 5:
			fmt.Println("Bye")
			return

		default:
			fmt.Println("Invalid choice, pick 1-4")
	}
}

}