// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
// You are given a target value to search. If found in the array return its index, otherwise return -1.
// You may assume no duplicate exists in the array.
// Your algorithm's runtime complexity must be in the order of O(log n).
// Example 1:
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
// Example 2:
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1

package main

func search(nums []int, target int) int {
	start, end, mid := 0, len(nums)-1, 0

	for start <= end {
		mid = start + (end-start)/2
		if target == nums[mid] {
			return mid
		} else if nums[mid] >= nums[start] {
			if target < nums[mid] && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] < nums[start] {
			if target > nums[mid] && target < nums[start] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
