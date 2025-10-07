package main

import "fmt"

// Open addressing concept - linear probing

/*
	Define a function hash(name, length):
		Find the sum of ASCII of name
		return sum % length

	Define a function insert(name):
		Find the slot where the name should go using the hash function
		If that slot is empty, place the name there
		If it's already taken, move forward to the next slot
		Keep moving until you find an empty one
		If you reach the end of the table, wrap around to the beginning.const

	Define a function to search for names:
		Use the hash function to find where the name *should* be.
		Check that slot:
			If it's the name we're looking for say found
			If it's empty stop and say, Not found
			If it's different name, move to the next slot and keep checking
		Keep looping until you either find the name or come back to where you started.const

	Main:
		Create an empty table with 5 slots
		Add the names
		When inserting if a slot is already taken, move to the next one
		Ask the system to check whether "Emma" is stored.
		Print 'found' if she is there, otherwie print 'not found	
	
*/

func hash(name string, size int) int{
	sum := 0
	for _, char := range name {
		sum += int(char)
	}
	return sum % 5
}


func insert(table []string, name string) {
	index := hash(name, len(table))
	for table[index] != "" {
		index = (index  + 1) % len(table) // move forward if occupied
	}
	table[index] = name
}

func search(table []string, target string) bool {
	index := hash(target, len(table))
	for table[index] != "" {
		if table[index] == target {
			return true
		}
		index = (index + 1) % len(table)
	}
	return false
}

func main() {
	table := make([]string, 5)

	names := []string{"Ava", "Evan", "Emma"}

	for _, name := range names {
		insert(table, name)
	}

	fmt.Println("Table: ", table)

	target := "Emma"
	if search(table, target) {
		fmt.Println(target, "found ✅")
	} else {
		fmt.Println(target, "not found ❌")
	}
}