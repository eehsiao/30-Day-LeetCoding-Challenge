// Given an integer array arr, count element x such that x + 1 is also in arr.
// If there're duplicates in arr, count them seperately.
// Example 1:
// Input: arr = [1,2,3]
// Output: 2
// Explanation: 1 and 2 are counted cause 2 and 3 are in arr.
// Example 2:
// Input: arr = [1,1,3,3,5,5,7,7]
// Output: 0
// Explanation: No numbers are counted, cause there's no 2, 4, 6, or 8 in arr.
// Example 3:
// Input: arr = [1,3,2,3,5,0]
// Output: 3
// Explanation: 0, 1 and 2 are counted cause 1, 2 and 3 are in arr.
// Example 4:
// Input: arr = [1,1,2,2]
// Output: 2
// Explanation: Two 1s are counted cause 2 is in arr.
// Constraints:
// 1 <= arr.length <= 1000
// 0 <= arr[i] <= 1000

package main

import "testing"

func Test_countElements(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				arr: []int{1, 1, 3, 3, 5, 5, 7, 7},
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				arr: []int{1, 3, 2, 3, 5, 0},
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				arr: []int{1, 1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countElements(tt.args.arr); got != tt.want {
				t.Errorf("countElements(%v) = %v, want %v", tt.args.arr, got, tt.want)
			}
		})
	}
}
