// Given a non-empty binary tree, find the maximum path sum.
// For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.
// Example 1:
// Input: [1,2,3]
//        1
//       / \
//      2   3
// Output: 6
// Example 2:
// Input: [-10,9,20,null,null,15,7]
//    -10
//    / \
//   9  20
//     /  \
//    15   7
// Output: 42

package main

// defin in week2-diameterofBinaryTree.go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxValue int

func retriveMax(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftMax := max(retriveMax(node.Left), 0)
	rightMax := max(retriveMax(node.Right), 0)
	maxValue = max(maxValue, node.Val+leftMax+rightMax)

	return node.Val + max(leftMax, rightMax)
}

func maxPathSum(root *TreeNode) int {
	// maxValue = -(int(^uint(0) >> 1)) - 1
	maxValue = root.Val
	retriveMax(root)

	return maxValue
}
