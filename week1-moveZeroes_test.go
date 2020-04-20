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
