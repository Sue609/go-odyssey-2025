package main

import "fmt"
/*
	DEFINE a hash function, takes in a string and an integer:
		convert string into a number - sum of Unicode values
		returns number MOD integer
	
	MAIN:
		Create a list of empty buckets, size = no. of names

		define our list of names

		for each name in the list:
			calculate index = hash(name, no. of buckets)
			insert name to corresponding bucket

		Check if 'Susan' exists:
			compute target_ index = hash("Susan", number of buckets)

			if 'susan' is in buckets[target_index]:
				return True
			else:
				return False

*/

func hash(name string, bucketSize int) int {
	sum := 0

	for _, char := range name {
		sum += int(char)
	}
	return sum % bucketSize
}

func main() {
	buckets := make([][]string, 5)

	names := []string{"Susan", "Alice", "Bob", "Cate", "Sam"}

	for _, name := range names {
		index := hash(name, len(buckets))
		buckets[index] = append(buckets[index], name)
	}

	target := "Susan"
	index := hash(target, len(buckets))
	found := false
	
	for _, v := range buckets[index] {
		if v == target {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Found ", target)
	} else {
		fmt.Println("Not Found ", target)
	}
}