package main

//There is an integer array nums sorted in ascending order (with distinct values).
//Prior to being passed to your function,nums is possibly rotated at an unknown pivot index k(`1 <= k < nums.length`) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] 0-indexed.
//For example,[0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
//
//Given the array nums after the possible rotation and an integer target, return
//the index of target if it is in nums, or -1 if it is not in nums.
//
//You must write an algorithm with O(log n) runtime complexity.

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for start < end {
		mid := (start + end) / 2

		if nums[mid] >= nums[0] {
			if nums[0] <= target && nums[mid] >= target {
				end = mid
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && nums[len(nums)-1] >= target {
				start = mid + 1
			} else {
				end = mid
			}
		}
	}

	if nums[start] == target {
		return start
	}

	return -1
}
