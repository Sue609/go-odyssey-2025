package main

import "fmt"

func main() {
	nodes := []string{"A", "B", "C", "D"}
	n := len(nodes)

	//Initialize 4x4 matrix (A,B,C,D)
	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Add edges - undirected
	edges := [][2]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}
	for _, e := range edges {
		u, v := e[0], e[1]
		matrix[u][v] = 1
		matrix[v][u] = 1 // undirected
	}

	fmt.Println("Adjacency matrix representation:")
	fmt.Printf(" %v\n", nodes)
	for i, row := range matrix {
		fmt.Printf("%s: %v\n", nodes[i], row)
	}
}