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
    // Example trees from the problem
    tree1 := []int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4}
    tree2 := []int{3, 5, 1, 6, 7, 4, 2, -1, -1, -1, -1, 9, 8}

    root1 := NewTreeNode(tree1)
    root2 := NewTreeNode(tree2)

    fmt.Println(leafSimilar(root1, root2)) // Expected output: true
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    // Get the leaf sequences of both trees
    leaves1 := findLeaves(root1)
    leaves2 := findLeaves(root2)

    // Compare the two slices
    if len(leaves1) != len(leaves2) {
        return false
    }
    for i := range leaves1 {
        if leaves1[i] != leaves2[i] {
            return false
        }
    }
    return true
}

func findLeaves(root *TreeNode) []int {
    var leaves []int
    var dfs func(node *TreeNode)

    // Helper function to perform DFS
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        if node.Left == nil && node.Right == nil {
            leaves = append(leaves, node.Val)
            return
        }
        dfs(node.Left)
        dfs(node.Right)
    }

    dfs(root)
    return leaves
}