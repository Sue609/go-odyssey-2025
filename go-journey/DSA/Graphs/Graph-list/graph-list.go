package main

import "fmt"

func main() {
    // Graph using adjacency list
    graph := map[string][]string{
        "A": {"B", "C"},
        "B": {"A", "D"},
        "C": {"A", "D"},
        "D": {"B", "C"},
    }

    fmt.Println("Adjacency List Representation:")
    for node, neighbors := range graph {
        fmt.Printf("%s -> %v\n", node, neighbors)
    }
}
