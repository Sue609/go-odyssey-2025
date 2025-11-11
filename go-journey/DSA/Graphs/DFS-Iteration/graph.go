package main

import "fmt"

// Defining our own data type, key[int] = node ID, value []int = list of neighbors
// Each key maps to a list of connected nodes
type Graph struct {
	adj map[int][]int
}

// AddEdge adds a directed edge (u -> v)
// (g *Graph) -> method receiver, It belongs to type Graph, and works as a pointer
func (g *Graph) AddEdge(u, v int) {
	if g.adj == nil {
		g.adj = make(map[int][]int)
	}
	g.adj[u] = append(g.adj[u], v)
}

// Iterative DFS
func (g *Graph) DFSIterative(start int) {
	visited := make(map[int]bool)
	stack := []int{start}

	for len(stack) > 0 {
		node := stack[len(stack) - 1] // pop from stack
		stack = stack[:len(stack) - 1]

		// skip if already visited
		if visited[node] {
			continue
		}

		// Mark as visited
		visited[node] = true
		fmt.Printf("%d ", node)

		// Push unvisited neighbours - can reverse order for consistent results
		for _, neighbor := range g.adj[node] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}
}


// Function checks if there is a path from source to destination using DFS
func (g *Graph) HasPath(source, destination int) bool {

	// Base case: If source is the same as destination
	if source == destination {
		return true
	}

	// Visit each neighbour recursively

	for _, neighbor := range g.adj[source] {
		if g.HasPath(neighbor, destination) {
			return true
		}
	}
	return false
}

func main() {
	g := Graph{}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)

	fmt.Println("Iterative DFS starting from node 0:")
	g.DFSIterative(0)

	fmt.Println()

	fmt.Println("\nPath checks")
	fmt.Println("Path 0 -> 6?", g.HasPath(0, 6)) // true
	fmt.Println("Path 2 -> 4?", g.HasPath(2, 4)) // false
	fmt.Println("Path 1 -> 4?", g.HasPath(1, 4)) // true
}