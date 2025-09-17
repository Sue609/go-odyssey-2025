package main

import "fmt"

func main() {
	// Array of 5 integers
	var numbers[5] int

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Array:", numbers)
	fmt.Print("First element:", numbers[0])
}