# Pseudocode for Insert and InOrder traversal
We create a struct - which is a blueprint for defining nodes, it stores data and pointers

Insert() - Adds value in the correct BTS order

InOrder() - Prints values in sorter order


```
DEFINE funtion Insert(node, value):
    IF node IS NULL:
        CREATE  a now Node with Value = value
        RETURN that new Node

    IF value < node.Value:
        node.Left = Insert(node.Left, value) // go left

    ELSE IF value > node.Value:
        node.Right = Insert(node.Right, value) // go right

    RETURN node // always return the current node back up


DEFINE function InOrder(node):
    IF node IS NULL:
        RETURN

    InOrder(node.Left)
    PRINT node.Value
    InOrder(node.Right)


DEFINE function deleteNode(root, value):
    if root is null:
        return root

    if value < root.Value
        root.Left = deleteNode(root.Right, value)

    else if value > root.value:
        root.Right = deleteNode(root.Right, value)

    else:
        if root.Left == null:
            return root.Right
        else if root.Right == null:
            return root.Left

        root.Value = minValue(root.Right)
        root.Right = deleteNode(root.Right, root.Value)

    return root


function minValue(node):
    current = node
    while current.Left != null:
        current = current.Left
    return current.Value
