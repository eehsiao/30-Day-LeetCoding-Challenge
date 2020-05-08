package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertLevelOrder(nums []interface{}, i int) (node *TreeNode) {
	if len(nums) > i && nums[i] != nil {
		node = &TreeNode{Val: nums[i].(int)}
		node.Left = insertLevelOrder(nums, 2*i+1)
		node.Right = insertLevelOrder(nums, 2*i+2)
	}

	return node
}

func inOrder(node *TreeNode) {
	if node != nil {
		inOrder(node.Left)
		fmt.Print(node.Val, " > ")
		inOrder(node.Right)
	}
}

func preOrder(node *TreeNode) {
	if node != nil {
		fmt.Print(node.Val, " > ")
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

func postOrder(node *TreeNode) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		fmt.Print(node.Val, " > ")
	}
}
