package main

import "fmt"

// Define a node, which represents a single node in the BST
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


func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return root
	}

	// Traverse the tree
	if value < root.Value {
		root.Left = deleteNode(root.Left, value)
	} else if value > root.Value {
		root.Right = deleteNode(root.Right, value)
	} else {
		// Node found
		// Case 1: No child
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// Case 2: One child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Case 3: Two children
		// Find the inorder successor - smallest value in the right subtree
		minNode := findMin(root.Right)
		root.Value = minNode.Value

		// Delete the inordr successor
		root.Right = deleteNode(root.Right, minNode.Value)
	}
	return root

}


// Helper function to find the smallest node
func findMin(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
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
	fmt.Println()

	// Delete a node
	root = deleteNode(root, 60)

	fmt.Println("Inorder after deleting 60")
	InOrder(root)
	fmt.Println()
}