package main

import (
	"fmt"
	"math"
)

// TreeNode represents a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxLevelSum performs a level-order traversal to find the level with maximum sum
// Time Complexity: O(N) where N is the number of nodes in the tree
// Space Complexity: O(W) where W is the maximum width of the tree
func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	maxLevelSum := math.MinInt
	maxLevel := 1
	currentLevel := 1
	// continue loop on each level with BFS until the queue (and levels) are depleted
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevelSum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevelSum += node.Val
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if currentLevelSum > maxLevelSum {
			maxLevelSum = currentLevelSum
			maxLevel = currentLevel
		}
		currentLevel++
	}
	return maxLevel
}

// Helper function to create a tree from array representation
func createTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		if i < len(arr) && arr[i] != nil {
			node.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(arr) && arr[i] != nil {
			node.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func main() {
	// Test cases
	testCases := [][]interface{}{
		{1, 2, 3, nil, 5, nil, 4},
		{1, nil, 3},
		{},
	}

	for i, test := range testCases {
		root := createTree(test)
		result := maxLevelSum(root)
		fmt.Printf("Test case %d: %v\n", i+1, result)
	}
}
