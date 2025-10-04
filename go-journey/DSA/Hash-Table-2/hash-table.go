package main

import (
	"fmt"
)

/*
	DEFINE hash function takes a string, integer:
		RETURN length of string % integer

	MAIN:
		Create an array of empty bucket
		DEFINE a list of songs

		for each song in songs:
			index = hash(song, len(buckets))
			insert song to the buckets[index]

		available-song = 'Thriller'
		index = hash(available-song, len(buckets))

		check if available-song is in bucket:
			Print available
		else:
			Print unavailable

*/

func hash(song string, length int) int {
	sum := 0

	for _, char := range song{
		sum += int(char)
	}
	return sum % length
}

func main() {
	buckets := make([][] string, 6)
	songs := []string {"Halo", "ShapeOfYou", "BlindingLights", "Despacito", "Thriller", "RollingInTheDeep"}

	for _, song := range songs {
		index := hash(song, len(songs))
		buckets[index] = append(buckets[index], song)
	}

	target := "Thrille"
	index := hash(target, len(buckets))
	found := false

	for _, song := range buckets[index]{
		if song == target {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Available")
	} else {
		fmt.Println("Unavailable")
	}
}