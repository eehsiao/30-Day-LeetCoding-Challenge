// Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Example:
// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Note:
// You must do this in-place without making a copy of the array.
// Minimize the total number of operations.
//    Hide Hint #1
// In-place means we should not be allocating any space for extra array. But we are allowed to modify the existing array. However, as a first step, try coming up with a solution that makes use of additional space. For this problem as well, first apply the idea discussed using an additional array and the in-place solution will pop up eventually.
//    Hide Hint #2
// A two-pointer approach could be helpful here. The idea would be to have one pointer for iterating the array and another pointer that just works on the non-zero elements of the array.

package main

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0, 1, 0},
			},
			want: []int{1, 0, 0},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{0, 0, 1},
			},
			want: []int{1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("%v maxSubArray() = %v, want %v", tt.name, tt.args.nums, tt.want)
			}
		})
	}
}
