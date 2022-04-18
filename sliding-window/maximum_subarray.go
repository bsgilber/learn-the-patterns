package main

/*
Given an integer array nums, find the contiguous subarray (containing at least one number)
	which has the largest sum and return its sum.

A subarray is a contiguous part of an array.

EASY
https://leetcode.com/problems/maximum-subarray/
*/

func maxSubArray(nums []int) int {
	var best int
	var sum int

	// base case where length of array = 1 for fast exit
	if len(nums) == 1 {
		return nums[0]
	}

	// otherwise start with the first element
	best = nums[0]
	sum = nums[0]

	// sliding window
	for i, num := range nums {
		if i == 0 {
			continue
		}

		// this logic allows us to drop values to the left of index i
		if num > sum+num {
			sum = num
		} else {
			sum = sum + num
		}

		if sum > best {
			best = sum
		}
	}

	return best
}
