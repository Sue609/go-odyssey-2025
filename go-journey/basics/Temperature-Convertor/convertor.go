package main
import "fmt"

func toFahrenheit(c float64) float64{
	return ((c * 9) / 5) + 32
}

func toCelsius(f float64) float64{
	return (f - 32) * 5 / 9
}

func main() {
	var choice int
	var temp float64

	fmt.Println("Choose the conversion type:")
	fmt.Println("1: Celsius to Fahrenheit")
	fmt.Println("2: Fahrenheit to Celsius")
	fmt.Print("Enter choice: ")

	fmt.Scanln(&choice)

	fmt.Println("Enter the temperature to convert: ")
	fmt.Scanln(&temp)

	switch choice {
	case 1:
		fmt.Printf("%.2f Celsius is %.2f Fahrenheit.\n", temp, toFahrenheit(temp))
	case 2:
		fmt.Printf("%.2f Fahrenheit is %.2f Celsius.\n", temp, toCelsius(temp))
	default:
    fmt.Println("Invalid choice. Please run the program again.")
}
}
