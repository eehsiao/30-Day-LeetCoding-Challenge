// Given an array of strings, group anagrams together.
// Example:
// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]
// Note:
// All inputs will be in lowercase.
// The order of your output does not matter.

package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	var (
		mapGroup   = make(map[string]int)
		strGroup   [][]string
		SortString = func(w string) string {
			s := strings.Split(w, "")
			sort.Strings(s)
			return strings.Join(s, "")
		}
	)

	for _, s := range strs {
		ss := SortString(s)
		if _, ok := mapGroup[ss]; !ok {
			mapGroup[ss] = len(strGroup)
			strGroup = append(strGroup, []string{s})
		} else {
			strGroup[mapGroup[ss]] = append(strGroup[mapGroup[ss]], s)
		}
	}

	return strGroup
}
