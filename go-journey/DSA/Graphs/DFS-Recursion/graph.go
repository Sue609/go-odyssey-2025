package main

import "fmt"

type Graph struct {
	adj map[int][]int
}

//AddEdge adds a directed edge u -> v
func (g *Graph) AddEdge(u, v int) {
	if g.adj == nil {
		g.adj = make(map[int][]int)
	}
	g.adj[u] = append(g.adj[u], v)
}

// Recursive DFS helper
func (g *Graph) DFSRecursive(node int, visited map[int] bool) {
	if visited[node] {
		return
	}

	visited[node] = true
	fmt.Printf("%d ", node)

	for _, neighbor := range g.adj[node] {
		g.DFSRecursive(neighbor, visited)
	}
}

// Wrapper for DFSrECURSIVE

func (g Graph) DFS(start int) {
	visited := make(map[int]bool)
	fmt.Printf("Recursive DFS starting from node %d:\n", start)
	g.DFSRecursive(start, visited)
}

func main() {
	g := Graph{}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)

	g.DFS(0)
}