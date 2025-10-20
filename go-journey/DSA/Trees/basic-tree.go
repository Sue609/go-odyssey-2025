package main

import "fmt"

// Define a node
type Node struct {
	Value int
	Left *Node
	Right *Node
}

// Insert function - addsa a new value to the tree
func Insert(root *Node, value int) *Node {
	// If the root is empty, create a new node as root
	if root == nil {
		return &Node{Value : value}
	}

	// If value is smaller, go to the left subtree
	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else if value > root.Value {
		// If value is larger, go to the right subtree
		root.Right = Insert(root.Right, value)
	}

	// Return the root - for recursion
	return root
}


// InOrder traversal (Left → Root → Right)
func InOrder(root *Node) {
	if root != nil {
		InOrder(root.Left)
		fmt.Print(root.Value, " ")
		InOrder(root.Right)
	}
}

func main() {
	// Build a simple tree
	var root *Node

	root = Insert(root, 50)
    root = Insert(root, 30)
    root = Insert(root, 70)
    root = Insert(root, 20)
    root = Insert(root, 40)
    root = Insert(root, 60)
    root = Insert(root, 80)

	fmt.Print("🌳 InOrder Traversal (sorted order): ")
	InOrder(root)
}