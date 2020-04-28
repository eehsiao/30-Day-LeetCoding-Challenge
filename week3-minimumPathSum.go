// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
// Note: You can only move either down or right at any point in time.
// Example:
// Input:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// Output: 7
// Explanation: Because the path 1→3→1→1→1 minimizes the sum.

package main

var min, road [][]int

func minPath(grid [][]int, from, to int) {
	w, d := len(grid[0]), len(grid)
	x, y := from%w, from/w
	tx, ty := to%w, to/w

	road[ty][tx] = 1
	defer func() {
		road[ty][tx] = 0
	}()

	min[ty][tx] = min[y][x] + grid[ty][tx]
	currValue := min[ty][tx]
	if to == w*d-1 {
		// fmt.Println("-----")
		// for i := 0; i < d; i++ {
		// 	fmt.Println(road[i])
		// }
		// fmt.Println("-----")
		// for i := 0; i < d; i++ {
		// 	fmt.Println(min[i])
		// }

		return
	}
	if tx+1 < w && (road[ty][tx+1] == 0) && (min[ty][tx+1] == 0 || min[ty][tx+1] > currValue+grid[ty][tx+1]) {
		minPath(grid, to, to+1)
	}
	if ty+1 < d && (road[ty+1][tx] == 0) && (min[ty+1][tx] == 0 || min[ty+1][tx] > currValue+grid[ty+1][tx]) {
		minPath(grid, to, to+w)
	}
	// if tx > 0 && (road[ty][tx-1] == 0) && (min[ty][tx-1] == 0 || min[ty][tx-1] > currValue+grid[ty][tx-1]) {
	// 	minPath(grid, to, to-1)
	// }
	// if ty > 0 && (road[ty-1][tx] == 0) && (min[ty-1][tx] == 0 || min[ty-1][tx] > currValue+grid[ty-1][tx]) {
	// 	minPath(grid, to, to-w)
	// }
}

func minPathSum(grid [][]int) int {
	w, d := 0, 0
	min, road = nil, nil
	if len(grid) > 0 {
		w, d = len(grid[0]), len(grid)
	}
	if w == 0 || d == 0 {
		return 0
	}
	if w == 1 && d == 1 {
		return grid[0][0]
	}
	for i := 0; i < d; i++ {
		min = append(min, make([]int, w))
		road = append(road, make([]int, w))
	}
	min[0][0], road[0][0] = grid[0][0], 1
	minPath(grid, 0, 1)

	if d > 1 && (min[1][0] == 0 || min[1][0] > grid[0][0]+grid[1][0]) {
		min[0][0], road[0][0] = grid[0][0], 1
		minPath(grid, 0, w)
	}
	return min[d-1][w-1]
}
