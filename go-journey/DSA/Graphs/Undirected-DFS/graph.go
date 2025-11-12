package main 

import "fmt"

type Graph struct {
	adj map[int][]int
}


// AddEdge adds an undirected edge, u <-> v, connections go both way
func (g *Graph) AddEdge(u, v int) {
	if g.adj == nil {
		g.adj = make(map[int][]int)
	}
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

// DFS - RECURSION - Recursive helper
func (g *Graph) DFSRecursive (node int, visited map[int]bool) {
	if visited[node] {
		return
	}

	visited[node] = true
	fmt.Printf("%d ", node)

	for _, neighbor := range g.adj[node] {
		g.DFSRecursive(neighbor, visited)
	}
}


// Wrapper for DFSRecursive - we are not using the pointer because this function does not modify the graph.
// It only reads form it - to traverse and print. It only receives a copy of the graph

func (g Graph) DFS(start int) {
	visited := make(map[int]bool)
	fmt.Printf("Recursive DFS startting from node %d: \n", start)
	g.DFSRecursive(start, visited)
	fmt.Println()
}


// Recursive path checker for undirected path
func (g *Graph) HasPath(source, destination int) bool {
	visited := make(map[int]bool)
	return g.hasPathRecursive(source, destination, visited)
}

// Recursive helper for HasPath
func (g *Graph) hasPathRecursive(current, destination int, visited map[int] bool) bool {
	// base case
	if current == destination {
		return true
	}

	// Avoid revisiting
	if visited[current] {
		return false
	}
	visited[current] = true

	// Explore neighbors
	for _, neighbor := range g.adj[current] {
		if g.hasPathRecursive(neighbor, destination, visited){
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

	g.DFS(1)

	fmt.Println("Path checks:")
	fmt.Println("Path 0 -> 6?", g.HasPath(0, 6)) // ✅ true
	fmt.Println("Path 2 -> 4?", g.HasPath(2, 4)) // ✅ true (undirected)
	fmt.Println("Path 5 -> 3?", g.HasPath(5, 3)) // ✅ true
	fmt.Println("Path 5 -> 10?", g.HasPath(5, 10)) // ❌ false
}