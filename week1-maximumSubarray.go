// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
// Example:
// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Follow up:
// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

package main

func maxSubArray(nums []int) int {
	var (
		sum, max int = 0, nums[0]
		maxSum   int = nums[0]
	)
	for i, v := range nums {
		sum += v
		if sum < 0 {
			sum = 0
		}
		if maxSum < sum {
			maxSum = sum
		}
		if i > 0 && nums[i] > max {
			max = nums[i]
		}
	}
	if max < 0 {
		maxSum = max
	}

	return maxSum
}
