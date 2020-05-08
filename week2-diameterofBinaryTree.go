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

func depth(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	L := depth(node.Left, ans)
	R := depth(node.Right, ans)
	*ans = max(*ans, L+R+1)
	return max(L, R) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int = 1
	depth(root, &ans)
	return ans - 1
}
