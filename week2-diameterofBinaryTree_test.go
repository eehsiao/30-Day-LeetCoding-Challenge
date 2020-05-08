// Diameter of Binary Tree
// Solution
// Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
// Example:
// Given a binary tree
//           1
//          / \
//         2   3
//        / \
//       4   5
// Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
// Note: The length of path between two nodes is represented by the number of edges between them.

package main

import (
	"fmt"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		nums []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []interface{}{1, 2, 3, 4, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := insertLevelOrder(tt.args.nums, 0)
			postOrder(root)
			fmt.Println()
			if got := diameterOfBinaryTree(root); got != tt.want {
				t.Errorf("%v diameterOfBinaryTree() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
