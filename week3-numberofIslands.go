// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
// Example 1:
// Input:
// 11110
// 11010
// 11000
// 00000
// Output: 1
// Example 2:
// Input:
// 11000
// 11000
// 00100
// 00011
// Output: 3

package main

func island(grid [][]byte, i int, n int) {
	w, d := len(grid[0]), len(grid)
	x, y := i%w, i/w

	if y > d-1 || grid[y][x] == '0' {
		return
	}
	if grid[y][x] == '1' {
		grid[y][x] = 'x'
	}
	if x+1 < w && grid[y][x+1] == '1' {
		island(grid, i+1, n)
	}
	if x > 0 && grid[y][x-1] == '1' {
		island(grid, i-1, n)
	}
	if y+1 < d {
		island(grid, i+w, n)
	}
	if y > 0 && grid[y-1][x] == '1' {
		island(grid, i-w, n)
	}
}

func numIslands(grid [][]byte) int {
	w, d, n := 0, 0, 0
	if len(grid) > 0 {
		w, d = len(grid[0]), len(grid)
	}
	if w == 0 || d == 0 {
		return 0
	}

	for i := 0; i < w*d; i++ {
		x, y := i%w, i/w
		if grid[y][x] == '1' {
			n++
			island(grid, i, n)
		}
	}

	return int(n)
}
