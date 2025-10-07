package main

import "fmt"

/*
	Define hash(students, length):
		return students % length

	define attendance(target, bucket):
		for each student in bucket:
			if student == lookup:
				output true
		output false

	define MAIN:
		create an array of empty bucket of size N
		store a list of students name

		for each student in students:
			index = hash(student, no. of bucket)
			insert student to bucket[index]

		lookup = "David"
		index = hash(lookup, no.of bucket)
		found = attendance(lookup, bucket[index])

		if found:
			print "present"
		else:
			print "absent"
*/

func hash(student string, length int) int {
	sum := 0
	for _, stdnt := range student {
		sum += int(stdnt)
	}
	return sum % length
}

func attendance(target string, buckets [] string) bool {
	for _, student := range buckets {
		if student == target {
			return true
		}
	}
	return false
}

func main() {
	buckets := make([][] string, 6)

	students := [] string{"Alice", "Brian", "Clara", "David", "Eva"}

	for _, student := range students {
		index := hash(student, len(buckets))
		buckets[index] = append(buckets[index], student)
	}

	target := "David"
	index := hash(target, len(buckets))
	found := attendance(target, buckets[index])

	if found {
		fmt.Println("Present")
	} else {
		fmt.Println("Absent")
	}

}