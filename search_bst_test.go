package leetcode

import "testing"

func TestSearchBST(t *testing.T) {
    tests := []struct {
        name     string
        root     *TreeNode
        val      int
        expected *TreeNode
    }{
        {
            name: "Example 1",
            root: &TreeNode{
                Val: 4,
                Left: &TreeNode{
                    Val: 2,
                    Left: &TreeNode{Val: 1},
                    Right: &TreeNode{Val: 3},
                },
                Right: &TreeNode{Val: 7},
            },
            val: 2,
            expected: &TreeNode{
                Val: 2,
                Left: &TreeNode{Val: 1},
                Right: &TreeNode{Val: 3},
            },
        },
        {
            name: "Example 2",
            root: &TreeNode{
                Val: 4,
                Left: &TreeNode{
                    Val: 2,
                    Left: &TreeNode{Val: 1},
                    Right: &TreeNode{Val: 3},
                },
                Right: &TreeNode{Val: 7},
            },
            val: 5,
            expected: nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := SearchBST(tt.root, tt.val)
            compareTreeNodes(t, result, tt.expected)
        })
    }
}

func compareTreeNodes(t *testing.T, got, want *TreeNode) {
    if got == nil && want == nil {
        return
    }
    if got == nil || want == nil {
        t.Errorf("Trees don't match: got = %v, want = %v", got, want)
        return
    }
    if got.Val != want.Val {
        t.Errorf("Node values don't match: got = %v, want = %v", got.Val, want.Val)
    }
    compareTreeNodes(t, got.Left, want.Left)
    compareTreeNodes(t, got.Right, want.Right)
}
