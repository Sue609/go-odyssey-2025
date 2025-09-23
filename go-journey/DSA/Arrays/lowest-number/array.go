package main 
import "fmt"
/*
	variable minValue = array[0]
	for each element in array:
		if currentElement < minValue
			minValue = currentElement
*/

func main() {
	
	numbers := []int {7, 12, 9, 4, 11};
	minValue := numbers[0]
	length := len(numbers)

	for element := 1; element < length; element++ {
		if numbers[element] < minValue {
			minValue = numbers[element]
		}
	}
	fmt.Println("The lowest number of the array is: ", minValue)
}