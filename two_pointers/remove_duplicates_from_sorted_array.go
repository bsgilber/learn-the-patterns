package main

/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that
each unique element appears only once. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the
result be placed in the first part of the array nums. More formally, if there are k elements after
removing the duplicates, then the first k elements of nums should hold the final result. It does not
matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place
with O(1) extra memory.

EASY
https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/

func removeDuplicates(nums []int) int {
	var counter int = 0
	var pointer1 int = 0

outer:
	for {
		if pointer1 > len(nums)-1 {
			break
		}

		var pointer2 int = 1
		for {
			if pointer1 == len(nums)-1 || pointer1+pointer2 > len(nums)-1 {
				nums[counter] = nums[pointer1]
				break outer
			}

			if nums[pointer1] == nums[pointer1+pointer2] {
				pointer2++
				continue
			}

			nums[counter] = nums[pointer1]
			pointer1 = pointer1 + pointer2
			counter++
			break
		}
	}
	return counter + 1
}

func removeDuplicatesBetter(nums []int) int {
	var pointer1 int = 0
	var pointer2 int = 1

	for {
		if pointer2 > len(nums)-1 {
			break
		}

		if nums[pointer1] != nums[pointer2] {
			pointer1++
			nums[pointer1] = nums[pointer2]
		}
		pointer2++
	}
	return pointer1 + 1
}
