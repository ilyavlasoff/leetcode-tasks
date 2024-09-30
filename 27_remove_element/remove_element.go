package main

//Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
//The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.
//
//Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:
//
//- Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
// The remaining elements of nums are not important as well as the size of nums.
//- Return k.

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	p := len(nums) - 1

	for p >= 0 && nums[p] == val {
		p--
	}

	if p < 0 {
		return 0
	}

	for i := 0; i < p; i++ {
		if nums[i] == val {
			nums[i], nums[p] = nums[p], nums[i]
			p--
			for p >= 0 && nums[p] == val {
				p--
			}
		}
	}

	return p + 1
}
