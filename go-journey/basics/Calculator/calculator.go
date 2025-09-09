package main

import "fmt"

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64{
	return a - b
}

func multiply(a, b float64) float64{
	return a * b
}

func division(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Division by 0 ")
		return 0
	}
	return a / b
}


func main() {
	var num1, num2 float64
	var operator string

	for {
		fmt.Printf("Enter first number(or type 0 to quit): ")
		fmt.Scanln(&num1)

		if num1 == 0 {
			fmt.Println("Goodbye 👋")
			break
		}
		
		fmt.Print("Enter operator(+, -, /, *): ")
		fmt.Scanln(&operator)

		fmt.Print("Enter the second number: ")
		fmt.Scanln(&num2)

		switch operator {
		case "+":
			fmt.Println("Result:", add(num1, num2))

		case "-":
			fmt.Println("Result:", subtract(num1, num2))

		case "*":
			fmt.Println("Result:", multiply(num1, num2))

		case "/":
			fmt.Println("Result:", division(num1, num2))

		default:
			fmt.Println("Invalid Operator")
		}
	}
}