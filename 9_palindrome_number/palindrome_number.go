package main

import "strconv"

// Given an integer `x`, return `true` if `x is a palindrome
// and false otherwise.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	v := strconv.Itoa(x)

	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		if v[i] != v[j] {
			return false
		}
	}

	return true
}
