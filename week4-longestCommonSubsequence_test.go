// Given two strings text1 and text2, return the length of their longest common subsequence.
// A subsequence of a string is a new string generated from the original string with some characters(can be none) deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.
// If there is no common subsequence, return 0.
// Example 1:
// Input: text1 = "abcde", text2 = "ace"
// Output: 3
// Explanation: The longest common subsequence is "ace" and its length is 3.
// Example 2:
// Input: text1 = "abc", text2 = "abc"
// Output: 3
// Explanation: The longest common subsequence is "abc" and its length is 3.
// Example 3:
// Input: text1 = "abc", text2 = "def"
// Output: 0
// Explanation: There is no such common subsequence, so the result is 0.
// Constraints:
// 1 <= text1.length <= 1000
// 1 <= text2.length <= 1000
// The input strings consist of lowercase English characters only.
//    Hide Hint #1
// Try dynamic programming. DP[i][j] represents the longest common subsequence of text1[0 ... i] & text2[0 ... j].
//    Hide Hint #2
// DP[i][j] = DP[i - 1][j - 1] + 1 , if text1[i] == text2[j] DP[i][j] = max(DP[i - 1][j], DP[i][j - 1]) , otherwise

package main

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				text1: "abcde",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				text1: "abc",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("%v longestCommonSubsequence() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
