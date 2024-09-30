package main

// Given two strings needle and haystack, return the index of the
// first occurrence of needle in haystack, or -1 if needle is not
// part of haystack.

func strStr(haystack string, needle string) int {
OUTER:
	for i := range haystack {
		for j, k := i, 0; k != len(needle); j, k = j+1, k+1 {
			if j == len(haystack) || needle[k] != haystack[j] {
				continue OUTER
			}
		}

		return i
	}

	return -1
}
