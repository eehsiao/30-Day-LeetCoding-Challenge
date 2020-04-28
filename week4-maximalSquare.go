// Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
// Example:
// Input:
// 1 0 1 0 0
// 1 0 1 1 1
// 1 1 1 1 1
// 1 0 0 1 0
// Output: 4

package main

var sw int

func findSquare(matrix [][]byte, from, to int) {
	w, d := len(matrix[0]), len(matrix)
	fx, fy := from%w, from/w
	x, y := to%w, to/w

	for i := x; i >= fx; i-- {
		if matrix[y][i] != '1' {
			return
		}
	}
	for i := y; i >= fy; i-- {
		if matrix[i][x] != '1' {
			return
		}
	}
	if x-fx >= sw {
		sw = x - fx + 1
	}

	if x+1 < w && y+1 < d && matrix[y+1][x+1] == '1' {
		findSquare(matrix, from, to+w+1)
	}

}

func maximalSquare(matrix [][]byte) int {
	w, d := 0, 0
	sw = 0
	if len(matrix) > 0 {
		w, d = len(matrix[0]), len(matrix)
	}
	if w == 0 || d == 0 {
		return 0
	}

	for i := 0; i < w*d; i++ {
		x, y := i%w, i/w
		if matrix[y][x] == '1' && sw == 0 {
			sw = 1
		}
		if x+1 < w && y+1 < d && matrix[y][x] == '1' && matrix[y+1][x+1] == '1' {
			findSquare(matrix, i, i+w+1)
		}
	}

	return sw * sw
}
