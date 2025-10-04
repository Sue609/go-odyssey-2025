package main

import "fmt"

/*
	DEFINE hash function, takes string and an integer:
		RETURN string % integer

	DEFINE verify(target, bucket)
		for name in bucket:
			IF name == target:
				RETURN true
		RETURN false

	MAIN function:
		CREATE empty buckets array with size N
		CREATE a list of names

		FOR name in names:
			index = hash(name, len(names))
			insert song to the buckets[index]

		target = "Paul"
		index = hash(target, len(bucket))
		found = verify(target, buckets[index])

		IF found:
			PRINT "Access Granted"
		ELSE:
			PRINT "Access Denied"
*/

func hash(name string, length int) int{
	sum := 0

	for _, char := range name {
		sum += int(char)
	}
	return sum % length
}

func verify(target string, bucket []string) bool {
	for _, name := range bucket {
		if name == target {
			return true
		}
	}
	return false
}

func main() {
	buckets := make([][]string, 15)
	names := []string{"Alice", "Bob", "Cate", "Daniel", "Eve", "Susan", "Mark", "Paul", "Rita", "John"}

	for _, name := range names {
		index := hash(name, len(buckets))
		buckets[index] = append(buckets[index], name)
	}

	target := "Paul"
	index := hash(target, len(buckets))
	found := verify(target, buckets[index])

	if found {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Denied")
	}
}