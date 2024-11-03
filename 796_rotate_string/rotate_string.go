package main

/*
Source: https://leetcode.com/problems/rotate-string/

Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s.
A shift on s consists of moving the leftmost character of s to the rightmost position.
For example, if s = "abcde", then it will be "bcdea" after one shift.
*/

func rotateString(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }

    if s == goal {
        return true
    }

    f, e := s[0], s[len(s)-1]

    for i := 0; i != len(goal)-1; i++ {
        if goal[i] == e && goal[i+1] == f {
            border := i+1

            if s == goal[border:] + goal[:border] {
                return true
            }
        }
    }

    return false
}