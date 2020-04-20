// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// Example:
// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Follow up:
// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

package main

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{-2, -1},
			},
			want: -1,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{-1, -2},
			},
			want: -1,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{-1, 2},
			},
			want: 2,
		},
		{
			name: "case 5",
			args: args{
				nums: []int{1, 2},
			},
			want: 3,
		},
		{
			name: "case 6",
			args: args{
				nums: []int{1, -2},
			},
			want: 1,
		},
		{
			name: "case 7",
			args: args{
				nums: []int{1, 2, -1, -2, 2, 1, -2, 1},
			},
			want: 3,
		},
		{
			name: "case 8",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "case 9",
			args: args{
				nums: []int{0, -3, 1, 1},
			},
			want: 2,
		},
		{
			name: "case 10",
			args: args{
				nums: []int{0, -3, -2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("%v maxSubArray(%v) = %v, want %v", tt.name, tt.args, got, tt.want)
			}
		})
	}
}
