package main

import (
	"fmt"
	"sort"
)

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that
i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

MEDIUM
https://leetcode.com/problems/3sum/
*/

func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	var results [][]int = [][]int{}
	tracker := make(map[string]int)
	ln := len(nums)

	if ln < 3 {
		return results
	}

	for i, num1 := range nums {
		if i == ln-2 {
			break
		}

		/*
			the "two pointers" are around index i, start (i+1) and end (n-1) are two potential
			indices to try. Because we are sorting to start we know if the sum > 0 then the last
			element is too large, and we move inwards; vice versa for sum < 0

			https://en.wikipedia.org/wiki/3SUM
		*/
		var start int = i + 1
		var end int = ln - 1

		for {
			if start >= end {
				break
			}

			num2 := nums[start]
			num3 := nums[end]
			if num1+num2+num3 == 0 {
				if _, ok := tracker[fmt.Sprintf("%d%d%d", num1, num2, num3)]; !ok {
					tracker[fmt.Sprintf("%d%d%d", num1, num2, num3)] = 1
					results = append(results, []int{num1, num2, num3})
				}
				start++
				end--
			} else if num1+num2+num3 > 0 {
				end--
			} else {
				start++
			}
		}

	}
	return results
}
