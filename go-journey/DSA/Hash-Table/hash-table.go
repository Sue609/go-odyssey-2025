package main

import "fmt"

/*
	define a hash function that takes a string and an integer
		returns length of string MOD no. of buckets

	create an array of empty buckets

	Define a list of names

	For each name in the list:
		Calculate bucket_index using hash function
		Insert name into the corresponding bucket

	For each bucket:
		Print bucket number and the content
*/ 

func hash(name string, buckets int) int {
	return len(name) % buckets
}

func main() {
	buckets := make([][]string, 5)

	names := []string{"Susan", "Alice", "Bob", "Cate", "Sam"}

	for _, name := range names {
		index := hash(name, len(buckets))
		buckets[index] = append(buckets[index], name)
	}

	// Print Contents
	for i, bucket := range buckets {
		fmt.Printf("Bucket %d: %v\n", i, bucket)
	}
}
