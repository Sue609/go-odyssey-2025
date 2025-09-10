package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num % 2 == 0 {
		fmt.Printf("The number %d is an even number.\n", num)
	} else {
		fmt.Printf("The number %d is an odd number.\n", num)
	}
}