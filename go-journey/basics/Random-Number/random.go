package main

import ("fmt"
	"math/rand"
)

func main() {
	secret := rand.Intn(10) + 1

	var guess int
	fmt.Print("Guess a number between 1 and 10: ")
	fmt.Scanln(&guess)

	if guess == secret {
		fmt.Println("🎉 Correct!")
 	}  else {
		fmt.Println("Try again 😅 (the correct number was", secret, ")")
}	}
