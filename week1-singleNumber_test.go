package main

// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
// Note:
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
// Example 1:
// Input: [2,2,1]
// Output: 1
// Example 2:
// Input: [4,1,2,1,2]
// Output: 4

import "testing"

func Test_singleNumber(t *testing.T) {
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
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
