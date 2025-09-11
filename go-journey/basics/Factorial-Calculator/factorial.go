package main

import "fmt"

func main() {
	var num int
	fact := 1

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	for i := 1; i <= num; i++ {
		fact *= i
	}

	fmt.Printf("Factorial of %d is: %d", num, fact)
}