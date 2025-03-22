// Package main implements a solution for LeetCode Path Sum III problem
package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// pathSum returns the number of paths where the sum equals targetSum
func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }
    return findPaths(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func findPaths(node *TreeNode, targetSum int) int {
    if node == nil {
        return 0
    }
    
    count := 0
    if node.Val == targetSum {
        count++
    }
    
    count += findPaths(node.Left, targetSum-node.Val)
    count += findPaths(node.Right, targetSum-node.Val)
    
    return count
}

// createTestTree creates a test tree with the example from LeetCode
func createTestTree() *TreeNode {
	return &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: -2},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val: -3,
			Right: &TreeNode{Val: 11},
		},
	}
}

func main() {
	// Create test cases
	testCases := []struct {
		root      *TreeNode
		targetSum int
		expected  int
	}{
		{createTestTree(), 8, 3},
		{nil, 0, 0},
	}

	// Run test cases
	for i, tc := range testCases {
		result := pathSum(tc.root, tc.targetSum)
		fmt.Printf("Test case %d: target sum = %d, paths found = %d, expected = %d\n",
			i+1, tc.targetSum, result, tc.expected)
	}
}
