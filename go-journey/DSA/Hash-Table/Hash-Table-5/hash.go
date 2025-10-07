package main

import "fmt"

/*
	define hash(name, length):
		return sum of ASCII values in name % length

	def authorization(target, buctets[indes]):
		for each name in bucket:
			for student == target:
				Return true
		Return false

	Main:
		create an array of empty buckets, size = N
		authorized_names = ["Liam", "Noah", "Olivia", "Emma", "Ava", "Sophia"]

		for each name in names:
			index = hash(name, no. of buckets)
			insert name to buckets[index]

		input scanned_name = "Enter your name:"
		index = buckets[index]
		found = authorization(scanned_name, buckets[index])

		if found 
			print "Access Granted"
		else 
		 Print "Access Denied"	
*/

func hash(name string, length int) int{
	sum := 0
	for _, char := range name {
		sum += int(char)
	}
	return sum % length
}

func authorization(scanned_name string, buckets []string) bool{
	for _, name := range buckets {
		if name == scanned_name {
			return true
		}
	}
	return false
}

func main() {
	buckets := make([][]string, 6)
	authorized_names := []string{"Liam", "Noah", "Olivia", "Emma", "Ava", "Sophia"}

	for _, name := range authorized_names {
		index := hash(name, len(buckets))
		buckets[index] = append(buckets[index], name)
	}

	var scanned_name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&scanned_name)
	index := hash(scanned_name, len(buckets))
	found := authorization(scanned_name, buckets[index])

	if found {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Denied")
	}
}