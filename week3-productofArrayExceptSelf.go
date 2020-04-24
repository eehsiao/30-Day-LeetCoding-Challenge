// Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
// Example:
// Input:  [1,2,3,4]
// Output: [24,12,8,6]
// Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array (including the whole array) fits in a 32 bit integer.
// Note: Please solve it without division and in O(n).
// Follow up:
// Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)

package main

func productExceptSelf(nums []int) []int {
	var (
		result []int = make([]int, len(nums))
	)

	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = nums[i-1] * result[i-1]

	}
	R := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * R
		R *= nums[i]

	}

	return result
}

// wrong
func _productExceptSelf(nums []int) []int {
	var (
		result     []int = make([]int, len(nums))
		i, product int   = 0, 1
	)

	for i = 0; i < len(nums)/2; i++ {
		if nums[i] != 0 {
			product = product * nums[i]
		}
		if nums[len(nums)-i-1] != 0 {
			product = product * nums[len(nums)-i-1]
		}
	}
	if i*2 != len(nums) {
		product = product * nums[i]
	}

	if product != 0 {
		for i = 0; i < len(nums)/2; i++ {
			if nums[i] != 0 {
				result[i] = product / nums[i]
			}
			if nums[len(nums)-i-1] != 0 {
				result[len(nums)-i-1] = product / nums[len(nums)-i-1]
			}
		}
		if i*2 != len(nums) {
			if nums[i] != 0 {
				result[i] = product / nums[i]
			}
		}
	}

	return result
}
