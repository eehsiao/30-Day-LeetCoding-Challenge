// Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.
// Example 1:
// Input: [0,1]
// Output: 2
// Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
// Example 2:
// Input: [0,1,0]
// Output: 2
// Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
// Note: The length of the given binary array will not exceed 50,000.

package main

func findMaxLength(nums []int) int {
	var (
		m           map[int]int = make(map[int]int)
		maxLen, cnt int         = 0, 0
	)
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		cnt += iif(nums[i] == 1, 1, -1).(int)
		if _, ok := m[cnt]; ok {
			maxLen = max(maxLen, i-m[cnt])
		} else {
			m[cnt] = i
		}
	}
	return maxLen
}
