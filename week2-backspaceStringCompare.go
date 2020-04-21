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

import (
	"strings"
)

func backspaceCompare(S string, T string) bool {
	var (
		s        []string = strings.Split(S, "")
		t        []string = strings.Split(T, "")
		compareS string   = ""
		compareT string   = ""
	)
	for si, ti, sbi, tbi := 1, 1, 0, 0; ; si, ti = si+1, ti+1 {
		if compareS != compareT || (si > len(s) && ti > len(t)) {
			break
		}
		for si <= len(s) {
			if s[len(s)-si] == "#" {
				sbi++
			} else if sbi > 0 {
				sbi--
			} else {
				break
			}
			si++
		}

		for ti <= len(t) {
			if t[len(t)-ti] == "#" {
				tbi++
			} else if tbi > 0 {
				tbi--
			} else {
				break
			}
			ti++
		}
		if si <= len(s) {
			compareS = s[len(s)-si] + compareS
		}
		if ti <= len(t) {
			compareT = t[len(t)-ti] + compareT
		}
	}
	return compareS == compareT
}
