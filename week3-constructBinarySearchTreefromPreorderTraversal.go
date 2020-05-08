// Return the node node of a binary search tree that matches the given preorder traversal.
// (Recall that a binary search tree is a binary tree where for every node, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)
// Example 1:
// Input: [8,5,1,7,10,12]
// Output: [8,5,10,1,7,null,12]
//      8
// 	   / \
//    5   10
//   / \    \
//  1 	7    12
// Note:
// 1 <= preorder.length <= 100
// The values of preorder are distinct.

package main

// defin in week2-diameterofBinaryTree.go
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

var order []int

func bstFromPreorder(preorder []int) *TreeNode {
	var bTree *TreeNode

	order = make([]int, 0)
	for _, v := range preorder {
		newNode := &TreeNode{
			Val: v,
		}

		if bTree == nil {
			bTree = newNode
			continue
		}
		node := bTree
		for {
			if node.Val > v {
				if node.Left != nil {
					node = node.Left
				} else {
					node.Left = newNode
					break
				}
			} else {
				if node.Right != nil {
					node = node.Right
				} else {
					node.Right = newNode
					break
				}
			}
		}
	}

	return bTree
}
