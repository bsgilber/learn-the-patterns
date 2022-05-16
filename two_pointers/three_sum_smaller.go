package main

import "sort"

/*
Given an array of n integers nums and an integer target, find the number of index triplets
i, j, k with 0 <= i < j < k < n that satisfy the condition nums[i] + nums[j] + nums[k] < target.

MEDIUM
https://leetcode.com/problems/3sum-smaller/
*/

func threeSumSmaller(nums []int, target int) int {
	sort.Sort(sort.IntSlice(nums))

}
