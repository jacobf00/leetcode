package main

import "fmt"

//   Definition for a binary tree node.
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

// NewTreeNode constructs a binary tree from a slice of integers.
// Use -1 to represent null nodes.
func NewTreeNode(treeArray []int) *TreeNode {
    if len(treeArray) == 0 {
        return nil
    }

    var root *TreeNode
    nodes := make([]*TreeNode, len(treeArray))
    for i, val := range treeArray {
        if val != -1 {
            nodes[i] = &TreeNode{Val: val}
        }
        if i == 0 {
            root = nodes[i]
		}
    }

    for i, node := range nodes {
        if node != nil {
            leftIdx := 2*i + 1
            rightIdx := 2*i + 2
            if leftIdx < len(treeArray) {
                node.Left = nodes[leftIdx]
            }
            if rightIdx < len(treeArray) {
                node.Right = nodes[rightIdx]
            }
        }
    }

    return root
}

func main() {
    // Example tree from the problem
    tree1 := []int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4}
    // tree2 := []int{3, 5, 1, 6, 7, 4, 2, -1, -1, -1, -1, 9, 8}

    root1 := NewTreeNode(tree1)
    // root2 := NewTreeNode(tree2)

    fmt.Println(goodNodes(root1))
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(node *TreeNode, maxSoFar int) int
	// Helper function to perform DFS
	dfs = func(node *TreeNode, maxSoFar int) int {
		if node == nil {
			return 0
		}

		// Count the current node as good if its value is >= maxSoFar
		good := 0
		if node.Val >= maxSoFar {
			good = 1
		}

		// Update maxSoFar for the child nodes
		newMax := maxSoFar
		if node.Val > maxSoFar {
			newMax = node.Val
		}

		// Recur for left and right children
		return good + dfs(node.Left, newMax) + dfs(node.Right, newMax)
	}

	// Start DFS with the root node and its value as the initial maxSoFar
	return dfs(root, root.Val)
}