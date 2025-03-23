package main

import "fmt"

// TreeNode represents a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rightSideView returns an array of values visible from the right side of the binary tree
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		
		// Process each level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// If it's the rightmost node in the level, add to result
			if i == levelSize-1 {
				result = append(result, node.Val)
			}

			// Add children to queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
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
		result := rightSideView(root)
		fmt.Printf("Test case %d: %v\n", i+1, result)
	}
}
