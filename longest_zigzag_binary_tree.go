package main

// TreeNode represents a binary tree node
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    maxLength := 0
    
    // DFS function to traverse the tree
    var dfs func(node *TreeNode, goLeft bool, length int)
    dfs = func(node *TreeNode, goLeft bool, length int) {
        if node == nil {
            return
        }
        
        // Update max length
        if length > maxLength {
            maxLength = length
        }
        
        if goLeft {
            // Continue zigzag going left
            dfs(node.Left, false, length+1)
            // Start new zigzag going right
            dfs(node.Right, true, 1)
        } else {
            // Continue zigzag going right
            dfs(node.Right, true, length+1)
            // Start new zigzag going left
            dfs(node.Left, false, 1)
        }
    }
    
    // Start DFS from both directions
    dfs(root, true, 0)
    dfs(root, false, 0)
    
    return maxLength
}

// NewTreeNode creates a new TreeNode with the given value
func NewTreeNode(val int) *TreeNode {
    return &TreeNode{Val: val}
}

func main() {
    // Test Case 1: [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]
    test1 := NewTreeNode(1)
    test1.Right = NewTreeNode(1)
    test1.Right.Left = NewTreeNode(1)
    test1.Right.Right = NewTreeNode(1)
    test1.Right.Right.Left = NewTreeNode(1)
    test1.Right.Right.Right = NewTreeNode(1)
    test1.Right.Right.Left.Right = NewTreeNode(1)

    // Test Case 2: [1,1,1,null,1,null,null,1,1,null,1]
    test2 := NewTreeNode(1)
    test2.Left = NewTreeNode(1)
    test2.Right = NewTreeNode(1)
    test2.Left.Right = NewTreeNode(1)
    test2.Left.Right.Left = NewTreeNode(1)
    test2.Left.Right.Right = NewTreeNode(1)
    test2.Left.Right.Right.Right = NewTreeNode(1)

    // Test Case 3: [1]
    test3 := NewTreeNode(1)

    // Run test cases
    testCases := []*TreeNode{test1, test2, test3}
    expectedResults := []int{3, 4, 0}

    for i, test := range testCases {
        result := longestZigZag(test)
        if result == expectedResults[i] {
            println("Test case", i+1, "passed! Result:", result)
        } else {
            println("Test case", i+1, "failed! Expected:", expectedResults[i], "Got:", result)
        }
    }
}
