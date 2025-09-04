package main

import "fmt"

func showVariables(){
	var a int = 6 // explicit type
	b := "hello, go" // type inferred
	const pi = 3.14 // constant

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", pi)
}