package main

import (
	"reflect"
	"testing"
)

func TestMaximumDepthBinaryTree(t *testing.T) {
	var tests = []struct {
		input    *TreeNode
		expected int
	}{
		// 		 3
		// 		/ \
		//	 9  20
		// 		 /  \
		// 	 15   7
		{&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}, 3},
		// 		 1
		// 		  \
		//	     2
		{&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		}, 2},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {

			got := maxDepth(test.input)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("maxDepth(tree) got %v, want %v", got, test.expected)
			}
		})
	}
}
