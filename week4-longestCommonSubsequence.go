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

import (
	"strings"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	var dp [][]int
	st1, st2 := strings.Split(text1, ""), strings.Split(text2, "")
	w, d := len(st1), len(st2)
	if w == 0 || d == 0 {
		return 0
	}

	for i := 0; i <= w; i++ {
		dp = append(dp, make([]int, d+1))
	}

	for i := w; i >= 0; i-- {
		for j := d; j >= 0; j-- {
			if i == w || j == d {
				dp[i][j] = 0
			} else if st1[i] == st2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}
