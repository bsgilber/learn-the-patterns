package main

/*
Given an array of positive integers nums and a positive integer target,
	return the minimal length of a contiguous subarray
	[numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than
	or equal to target. If there is no such subarray, return 0 instead.

MEDIUM
https://leetcode.com/problems/minimum-size-subarray-sum/
*/

func minSubArrayLen(target int, nums []int) int {
	var bestLen int = len(nums) + 1
	var sum int = 0  // sum of the subarray being evaluated
	var left int = 0 // a reference to the left-most index

	for i, num := range nums {
		sum = sum + num
	inner:
		for {
			// while block only needed to trim sums > target
			if sum < target {
				break inner
			} else {
				// check if the subarray is smaller than the current-best subarray length
				// adding one is clear for the case where target = nums[0], left = 0
				// 		(0-0+1) yields a subarray of length 1
				if (i - left + 1) <= bestLen {
					bestLen = i - left + 1
				}
				// snake eats its' tail, drop the left-most value in the subarray
				sum = sum - nums[left]
				left = left + 1
			}
		}
	}

	// if bestLen is unchanged then the target value was never hit
	if bestLen == len(nums)+1 {
		return 0
	}

	return bestLen
}
