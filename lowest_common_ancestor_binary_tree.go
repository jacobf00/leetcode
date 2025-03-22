package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func buildTree(values []interface{}) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := findNode(root.Left, val)
	if left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func main() {
	// Test case 1
	values1 := []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	tree1 := buildTree(values1)
	p1 := findNode(tree1, 5)
	q1 := findNode(tree1, 1)
	result1 := lowestCommonAncestor(tree1, p1, q1)
	fmt.Printf("Test 1 LCA: %d\n", result1.Val) // Expected: 3

	// Test case 2
	values2 := []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	tree2 := buildTree(values2)
	p2 := findNode(tree2, 5)
	q2 := findNode(tree2, 4)
	result2 := lowestCommonAncestor(tree2, p2, q2)
	fmt.Printf("Test 2 LCA: %d\n", result2.Val) // Expected: 5
}
