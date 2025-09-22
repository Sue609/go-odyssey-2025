package main
import "fmt"

func main() {
	num1 := 0
	num2 := 1
	var Newfib int

	fmt.Printf("%d\n", num1)
	fmt.Printf("%d\n", num2)

	for fib := 1; fib < 18; fib ++ {
		Newfib = num1 + num2
		fmt.Printf("%d\n", Newfib)
		num1 = num2
		num2 = Newfib
	}
}