package main

import "fmt"

type Graph struct {
	adj map[int][]int
}

// AddEdge adds a directed edge u -> v

func (g *Graph) AddEdge(u, v int) {
	if g.adj == nil {
		g.adj = make(map[int][]int)
	}
	g.adj[u] = append(g.adj[u], v)
}

// BFS traversal starting from 'start' node
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start} // Initialize queue with start node

	fmt.Printf("BFS starting from node %d:\n", start)

	for len(queue) > 0 {
		node := queue[0] // dequeue from fron
		queue = queue[1:] // remove first element

		if visited[node] {
			continue
		}

		visited[node] = true
		fmt.Printf("%d ", node)

		// Enqueue all visited neighbours
		for _, neighbor := range g.adj[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	g := Graph{}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)

	g.BFS(0)
}