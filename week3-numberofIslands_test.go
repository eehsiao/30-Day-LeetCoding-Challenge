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

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				grid: [][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				grid: [][]byte{},
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				grid: [][]byte{
					{'1', '1', '1'},
					{'0', '1', '0'},
					{'1', '1', '1'},
				},
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				grid: [][]byte{
					{'1', '0', '1', '1', '1'},
					{'1', '0', '1', '0', '1'},
					{'1', '1', '1', '0', '1'},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("%v numIslands() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
