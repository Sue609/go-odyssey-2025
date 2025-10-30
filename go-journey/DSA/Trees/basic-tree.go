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
func InorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	result = append(result, InorderTraversal(root.Left)...)
	result = append(result, root.Value)
	result = append(result, InorderTraversal(root.Right)...)
	return result
}


// PreOrder traversal Root → Left → Right
func PreorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}
	result := []int{root.Value}
	result = append(result, PreorderTraversal(root.Left)...)
	result = append(result, PreorderTraversal(root.Right)...)
	return result
}



// PostOrder traversal - Order: Left → Right → Root
func PostorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, PostorderTraversal(root.Left)...)
	result = append(result, PostorderTraversal(root.Right)...)
	result = append(result, root.Value)
	return result
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


// This function is an HTTP handler — it runs whenever someone visits the /tree endpoint (like http://localhost:8080/tree).
func treeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(root)
}


/*
 	TraverseHandler is an HTTP handler function - it handles incoming web requests to your endpoint
 	It reads what kind of traversal the user wants (inorder, preorder, or postorder),
	performs that traversal on your tree, and then sends back the result in JSON format.

	Example: GET /traverse?type=inorder
*/
func traverseHandler(w http.ResponseWriter, r *http.Request) {
	traversalType := r.URL.Query().Get("type")
	w.Header().Set("Content-Type", "application/json")
	
	var result []int

	switch traversalType {
		case "inorder":
			result = InorderTraversal(root)
		
		case "preorder":
			result = PreorderTraversal(root)
		
		case "postorder":
			result = PostorderTraversal(root)
		
		default:
			http.Error(w, "Invalid traversal type", http.StatusBadRequest)
			return
	}

	json.NewEncoder(w).Encode(map[string]interface{} {
		"type": traversalType,
		"result": result,
	})
}


func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/tree", treeHandler)
	http.HandleFunc("/traverse", traverseHandler)


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