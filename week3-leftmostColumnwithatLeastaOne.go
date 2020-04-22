// (This problem is an interactive problem.)
// A binary matrix means that all elements are 0 or 1. For each individual row of the matrix, this row is sorted in non-decreasing order.
// Given a row-sorted binary matrix binaryMatrix, return leftmost column index(0-indexed) with at least a 1 in it. If such index doesn't exist, return -1.
// You can't access the Binary Matrix directly.  You may only access the matrix using a BinaryMatrix interface:
// BinaryMatrix.get(x, y) returns the element of the matrix at index (x, y) (0-indexed).
// BinaryMatrix.dimensions() returns a list of 2 elements [n, m], which means the matrix is n * m.
// Submissions making more than 1000 calls to BinaryMatrix.get will be judged Wrong Answer.  Also, any solutions that attempt to circumvent the judge will result in disqualification.
// For custom testing purposes you're given the binary matrix mat as input in the following four examples. You will not have access the binary matrix directly.
// Example 1:
// Input: mat = [[0,0],[1,1]]
// Output: 0
// Example 2:
// Input: mat = [[0,0],[0,1]]
// Output: 1
// Example 3:
// Input: mat = [[0,0],[0,0]]
// Output: -1
// Example 4:
// Input: mat = [[0,0,0,1],[0,0,1,1],[0,1,1,1]]
// Output: 1
// Constraints:
// 1 <= mat.length, mat[i].length <= 100
// mat[i][j] is either 0 or 1.
// mat[i] is sorted in a non-decreasing way.
//    Hide Hint #1
// 1. (Binary Search) For each row do a binary search to find the leftmost one on that row and update the answer.
//    Hide Hint #2
// 2. (Optimal Approach) Imagine there is a pointer p(x, y) starting from top right corner. p can only move left or down. If the value at p is 0, move down. If the value at p is 1, move left. Try to figure out the correctness and time complexity of this algorithm.

package main

type BinaryMatrix struct {
	matrix [][]int
}

func (b BinaryMatrix) Get(x int, y int) int {
	return b.matrix[x][y]
}

func (b BinaryMatrix) Dimensions() (dim []int) {
	return append(dim, len(b.matrix), len(b.matrix[0]))
}

func binarySearch(b BinaryMatrix, x, left, right int) int {
	if right > 0 {
		mid := left + (right-left)/2
		l, r := b.Get(x, mid), b.Get(x, mid+1)

		if l == 0 && r == 1 {
			return mid + 1
		}
		if l == 1 {
			return binarySearch(b, x, left, mid)
		}

		return binarySearch(b, x, mid, right)
	}

	if b.Get(x, 0) == 1 {
		return 0
	} else {
		return -1
	}
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	var dim []int = binaryMatrix.Dimensions()
	var x, leftMost int = 0, dim[1] - 1
	var hit bool
	for ; x < dim[0]; x++ {
		if binaryMatrix.Get(x, leftMost) == 1 {
			hit = true
			leftMost = binarySearch(binaryMatrix, x, 0, leftMost)
		}
	}
	if !hit {
		leftMost = -1
	}
	return leftMost
}
