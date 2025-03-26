package main

// TreeNode represents a binary tree node
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        // Case 1: leaf node
        if root.Left == nil && root.Right == nil {
            return nil
        }
        // Case 2: one child
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }
        // Case 3: two children
        // Find smallest value in right subtree (successor)
        successor := findMin(root.Right)
        root.Val = successor.Val
        // Delete the successor
        root.Right = deleteNode(root.Right, successor.Val)
    }
    return root
}

// findMin returns the node with minimum value in the tree
func findMin(node *TreeNode) *TreeNode {
    current := node
    for current.Left != nil {
        current = current.Left
    }
    return current
}
