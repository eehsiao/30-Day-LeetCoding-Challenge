// Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.
// Example 1:
// Input: [5,7]
// Output: 4
// Example 2:
// Input: [0,1]
// Output: 0

package main

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				m: 5,
				n: 7,
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				m: 0,
				n: 1,
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				m: 1,
				n: 2,
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				m: 1,
				n: 1,
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				m: 0,
				n: 0,
			},
			want: 0,
		},
		{
			name: "case 6",
			args: args{
				m: 3,
				n: 3,
			},
			want: 3,
		},
		{
			name: "case 7",
			args: args{
				m: 6,
				n: 7,
			},
			want: 6,
		},
		{
			name: "case 8",
			args: args{
				m: 12,
				n: 14,
			},
			want: 12,
		},
		{
			name: "case 9",
			args: args{
				m: 12,
				n: 140,
			},
			want: 0,
		},
		{
			name: "case 10",
			args: args{
				m: 10,
				n: 11,
			},
			want: 10,
		},
		{
			name: "case 11",
			args: args{
				m: 20,
				n: 22,
			},
			want: 20,
		},
		{
			name: "case 12",
			args: args{
				m: 32,
				n: 57,
			},
			want: 32,
		},
		{
			name: "case 13",
			args: args{
				m: 64,
				n: 113,
			},
			want: 64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("%v rangeBitwiseAnd(%v, %v) = %v, want %v", tt.name, tt.args.m, tt.args.n, got, tt.want)
			}
		})
	}
}
