package main
import "fmt"

func main () {
	// Creating a slice

	fruits := []string{"Apple", "Banana", "Kiwi", "Pineapple"}

	// Adding an element using append

	fruits = append(fruits, "Orange")

	// Printing all fruits
	fmt.Println("Fruits:", fruits)

	// Loop through slice
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i, fruit)
	}
}