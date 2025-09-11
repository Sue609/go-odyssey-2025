package main

import "fmt"

//func multiplicationTable(num float64) float64{

//}

func main() {
	var num int

	fmt.Println("Enter a number: ")
	fmt.Scanln(&num)

    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", num, i, num*i)
    }
}