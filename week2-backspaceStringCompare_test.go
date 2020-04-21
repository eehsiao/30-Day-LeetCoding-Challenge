// Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.
// Note that after backspacing an empty text, the text will continue empty.
// Example 1:
// Input: S = "ab#c", T = "ad#c"
// Output: true
// Explanation: Both S and T become "ac".
// Example 2:
// Input: S = "ab##", T = "c#d#"
// Output: true
// Explanation: Both S and T become "".
// Example 3:
// Input: S = "a##c", T = "#a#c"
// Output: true
// Explanation: Both S and T become "c".
// Example 4:
// Input: S = "a#c", T = "b"
// Output: false
// Explanation: S becomes "c" while T becomes "b".
// Note:
// 1 <= S.length <= 200
// 1 <= T.length <= 200
// S and T only contain lowercase letters and '#' characters.
// Follow up:
// Can you solve it in O(N) time and O(1) space?

package main

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				S: "ab#c",
				T: "ad#c",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				S: "ab##",
				T: "c#d#",
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				S: "a##c",
				T: "#a#c",
			},
			want: true,
		},
		{
			name: "case 4",
			args: args{
				S: "a#c",
				T: "b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("%v backspaceCompare(%v,%v) = %v, want %v", tt.name, tt.args.S, tt.args.T, got, tt.want)
			}
		})
	}
}
