package main

//   Definition for a binary tree node.
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func NewTreeNode(treeArray []int) *TreeNode {
	
}

func main() {
	

}
func maxDepth(root *TreeNode) int {
    // Base case: if root is nil, depth is 0
    if root == nil {
        return 0
    }
    
    // Recursively calculate the depth of left and right subtrees
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    
    // Return the maximum of left and right depths, plus 1 for current node
    return max(leftDepth, rightDepth) + 1
}

// Helper function to get maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}