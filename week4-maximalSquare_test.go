// Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
// Example:
// Input:
// 1 0 1 0 0
// 1 0 1 1 1
// 1 1 1 1 1
// 1 0 0 1 0
// Output: 4

package main

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]byte{
					{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'},
				},
			},
			want: 4,
		},
		{
			name: "case 1",
			args: args{
				matrix: [][]byte{
					{'0'},
				},
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				matrix: [][]byte{
					{'1'},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("%v maximalSquare() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
