package main

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

EASY
https://leetcode.com/problems/two-sum/
*/

func twoSum(nums []int, target int) []int {
	var indexPair []int = []int{-1, -1}

	for i, num := range nums {
		for j, n := range nums {
			if j == i {
				continue
			}

			if num+n == target {
				indexPair = []int{j, i}
				break
			}
		}
	}

	return indexPair
}

func twoSumUsingDict(nums []int, target int) []int {
	var indexPair []int = []int{-1, -1}
	m := make(map[int]int)

	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			indexPair = []int{idx, i}
		}

		m[num] = i
	}

	return indexPair
}
