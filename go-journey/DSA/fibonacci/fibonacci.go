package main
import "fmt"

func main() {
	var fibNum int
	fmt.Println("Enter the fibonacci numbers you want to generate: ")
	fmt.Scanln(&fibNum)

	num1 := 0
	num2 := 1
	var Newfib int

	if fibNum >= 1 {
		fmt.Printf("%d\n", num1)
	}

	if fibNum >= 2 {
		fmt.Printf("%d\n", num2)
	}

	for i := 3; i <= fibNum; i++ {
		Newfib = num1 + num2
		fmt.Printf("%d\n", Newfib)

		num1 = num2
		num2 = Newfib
	}
}