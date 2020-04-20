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
