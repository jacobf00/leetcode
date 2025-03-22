package main

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		values   []interface{}
		p, q     int
		expected int
	}{
		{
			name:     "Test case 1",
			values:   []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
			p:        5,
			q:        1,
			expected: 3,
		},
		{
			name:     "Test case 2",
			values:   []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
			p:        5,
			q:        4,
			expected: 5,
		},
		{
			name:     "Test case 3",
			values:   []interface{}{1, 2},
			p:        1,
			q:        2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.values)
			p := findNode(root, tt.p)
			q := findNode(root, tt.q)
			result := lowestCommonAncestor(root, p, q)
			if result.Val != tt.expected {
				t.Errorf("got %v, want %v", result.Val, tt.expected)
			}
		})
	}
}
