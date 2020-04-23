// Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.
// Example 1:
// Input:nums = [1,1,1], k = 2
// Output: 2
// Note:
// The length of the array is in range [1, 20,000].
// The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].
//    Hide Hint #1
// Will Brute force work here? Try to optimize it.
//    Hide Hint #2
// Can we optimize it by using some extra space?
//    Hide Hint #3
// What about storing sum frequencies in a hash table? Will it be useful?
//    Hide Hint #4
// sum(i,j)=sum(0,j)-sum(0,i), where sum(i,j) represents the sum of all the elements from index i to j-1. Can we use this property to optimize it.
// https://leetcode.com/problems/subarray-sum-equals-k/solution/

package main

func subarraySum(nums []int, k int) int {
	var (
		sumMap   map[int]int = make(map[int]int)
		cnt, sum int
	)
	sumMap[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := sumMap[sum-k]; ok {
			cnt += v
		}
		if _, ok := sumMap[sum]; ok {
			sumMap[sum]++
		} else {
			sumMap[sum] = 1
		}
	}
	return cnt
}

//Time Limit Exceeded
func _subarraySum(nums []int, k int) int {
	var (
		sumMap map[int]int = make(map[int]int)
		cnt    int
	)
	sumMap[-1] = 0
	for i := 0; i < len(nums); i++ {
		sumMap[i] = sumMap[i-1] + nums[i]
		if sumMap[i] == k {
			cnt++
		}
		for j := 0; j < i; j++ {
			if sumMap[i]-sumMap[j] == k {
				cnt++
			}
		}
	}
	return cnt
}
