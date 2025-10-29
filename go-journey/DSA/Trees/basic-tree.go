package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Define a node, which represents a single node in the BST
type Node struct {
	Value int `json:"value"`
	Left *Node `json:"left,omitempty"`
	Right *Node `json:"right,omitempty"`
}

var root *Node

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

// PreOrder traversal Root → Left → Right
func PreOrder(root *Node) {
	if root != nil {
		fmt.Print(root.Value, "")
		PreOrder(root.Left)
		PreOrder(root.Right)
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


// === Handlers ===
func insertHandler(w http.ResponseWriter, r *http.Request) {
	valStr := r.URL.Query().Get("value")
	val, err := strconv.Atoi(valStr)
	
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}
	root = Insert(root, val)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Inserted %d seuccessfully!", val)
}


func deleteHandler(w http.ResponseWriter,  r *http.Request) {
	valStr := r.URL.Query().Get("value")
	val, err := strconv.Atoi(valStr)

	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}
	root = deleteNode(root, val)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted %d successfully!", val)
}


func treeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(root)
}


func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/tree", treeHandler)

	// Logic for clearing tree
	http.HandleFunc("/tree/clear", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			root = nil // Reset global root variable
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Tree cleared successfully"}`))
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("🌲 API running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}