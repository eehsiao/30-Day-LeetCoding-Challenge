// Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.
// Example 1:
// Input: [0,1]
// Output: 2
// Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
// Example 2:
// Input: [0,1,0]
// Output: 2
// Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
// Note: The length of the given binary array will not exceed 50,000.

package main

import "testing"

func Test_findMaxLength(t *testing.T) {
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
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0, 1, 0},
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{0, 0, 1, 0, 0, 0, 1, 1},
			},
			want: 6,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{0, 1, 0, 0, 1, 1, 0},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("%v findMaxLength() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
