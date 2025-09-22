package main
import "fmt"

var count int = 2

func fibonacci(num1, num2 int) {
	if (count <= 18) {
		newFib := num1 + num2
		fmt.Println(newFib)

		num1 = num2
		num2 = newFib

		count ++
		fibonacci(num1, num2)
	} 
}


func main() {
	fmt.Println("0")
	fmt.Println("1")
	fibonacci(0, 1)
}