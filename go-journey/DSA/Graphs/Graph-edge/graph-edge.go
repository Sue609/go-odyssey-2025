package main

import "fmt"

func main() {
	// List of edges - undirected graph

	edges := [][2]string{
		{"A", "B"},
		{"A", "C"},
		{"B", "D"},
		{"C", "D"},
	}

	fmt.Println("Edge list representation:")
	for _, edge := range edges {
		fmt.Printf("%s -- %s\n", edge[0], edge[1])
	}
}