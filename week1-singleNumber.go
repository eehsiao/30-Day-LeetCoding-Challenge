// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
// Note:
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
// Example 1:
// Input: [2,2,1]
// Output: 1
// Example 2:
// Input: [4,1,2,1,2]
// Output: 4

package main

func singleNumber(nums []int) int {
	var (
		singleN   = make(map[int]struct{}, 0)
		multiN    = make(map[int]struct{}, 0)
		returnKey int
	)

	for _, n := range nums {
		if _, ok := multiN[n]; ok {
			continue
		}
		if _, ok := singleN[n]; ok {
			multiN[n] = struct{}{}
			delete(singleN, n)
		} else {
			singleN[n] = struct{}{}
		}
	}

	for k, _ := range singleN {
		returnKey = k
		break
	}

	return returnKey
}
