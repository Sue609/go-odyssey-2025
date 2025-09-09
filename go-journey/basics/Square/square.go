package main

import "fmt"

//Square function works for any numeric type int, float, etc
func square[T int | float64](a T) T{
	return a * a
}

func main() {
	var num, intNum float64
	
	fmt.Print("Enter a float number: ")
	fmt.Scanln(&num)
	fmt.Println("Float Result: ", square(num))

	fmt.Print("Enter an integer number: ")
	fmt.Scanln(&intNum)
	fmt.Println("Integer Result: ", square(intNum))

}