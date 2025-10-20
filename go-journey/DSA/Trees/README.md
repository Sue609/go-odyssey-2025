# **🌳 TREE DATA STRUCTURE**
# 📖 Overview
A **Tree** is a hierarchical data structure made up of **nodes** connected by **edges**. It's one of the most important concepts in DSA, used in databases, compilers, networking, and even AI.

We'll learn how trees work from the ground up using  **Go** for implementation.

---

# **🧩 Why Trees matter**
Trees let us:
- Store data **hierachically** like folders and files on a computer).
- Search, inset, and delete data efficicently (better than linear search).
- Represented structured relationships (like HTML, JSON, or organization charts).
- Implement other complex structures like **heaps, tries, and binary search trees**

---

## 🧱 Core Concepts
### **1. Introduction to Trees**
- What a tree is (nodes, edges, root, leaves, height, depth)
- Real life anologies of trees.
- Difference between **Binary Trees, Binary Search Trees, and General Trees**

### **2. Binary Search Tree (BST)**
- Structure and properties of BTS.
- Operations:
    - Insert 🌱
    - Search 🔍
    - Traversals (InOrder, PreOrder, PostOrder)
    - Deletion (leaf, one child, two children)
- Time complexity: O(log n) average, O(n) worst case

### **3. Tree Traversals**
We'll look at three most traversal orders:
- **InOrder Left -> Root -> Right** - returns nodes in sorted order.
- **PreOrder Root -> Left -> Right** - useful for tree copying or expression parsing
- **PostOrder Left -> Right -> Root** - useful for deletion or postfix evaluation

### **4. Balanced Trees (Overview)**
- Problem: Unbalanced trees -> degraded performance
- Introduction to:
    - AVL Trees
    - Red-Black Trees
    - B-Trees (used in databases)

### **5. Advanced Trees (To Be Covered Later)**
- **Heaps** (used in priority queues)
- **Segment Trees & Fenwick Trees** (for range queries)
- **N-ary Trees** (used in compilers and hierarchical models)

---

## ⚙️ Implementation (in Go)
```go
type Node struct {
    Value int
    Left *Node
    Right *Node
}
```

---

## 🧠 Complexity Overview
| Operation | Average Case | Worst Case |
| --------- | ------------ | ---------- |
| Search    | O(log n)     | O(n)       |
| Insert    | O(log n)     | O(n)       |
| Delete    | O(log n)     | O(n)       |
| Traversal | O(n)         | O(n)       |

---

## 🧭 Roadmap

 Understand basic Tree structure

 Implement Binary Search Tree (BST)

 Add traversal functions

 Implement delete operation

 Explore balanced trees

 Extend to Heaps and Tries

 ---

 # **🧑‍💻 Author**
 **Susan K.**
 Learning Go by exploring Data Structures and Algorithms 🌿
 ```
 “Every branch of knowledge grows from a strong root of understanding.”
 ```